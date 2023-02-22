#include <WiFi.h>
#include <PubSubClient.h>
#include "DHT.h"
#include <Wire.h>
#include <QMC5883L.h>


// ---------------------------------------- wifi ----------------------------------------
const char* ssid = "test";
const char* password = "12345678";

void WiFiStationConnected(WiFiEvent_t event, WiFiEventInfo_t info){
  Serial.println("Connected to AP successfully!");
}

void WiFiGotIP(WiFiEvent_t event, WiFiEventInfo_t info){
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void WiFiStationDisconnected(WiFiEvent_t event, WiFiEventInfo_t info){
  Serial.println("Disconnected from WiFi access point");
  Serial.print("WiFi lost connection. Reason: ");
  Serial.println(info.wifi_sta_disconnected.reason);
  Serial.println("Trying to Reconnect");
  WiFi.begin(ssid, password);
}

// ----------------------------------------- MQTT ----------------------------------------
WiFiClient esp32Client;
PubSubClient client(esp32Client);

class MQTT{
  private:
    String client_id;
    const char *mqtt_broker;
    int mqtt_port;
    
  public:
    MQTT(String client_id, const char *mqtt_broker, int mqtt_port){
      Serial.println("MQTT brocker credentials are set.");
      this->client_id = client_id;
      this->mqtt_broker = mqtt_broker;
      this->mqtt_port = mqtt_port;
      
      client.setServer(this->mqtt_broker, this->mqtt_port);
      client.setCallback(MQTT::callback);
      while (!client.connected()) {
          this->client_id += String(WiFi.macAddress());
          Serial.printf("The client %s is connecting to the public mqtt broker\n", this->client_id.c_str());
          if (client.connect(this->client_id.c_str())) {
              Serial.println("Public hiveMQ mqtt broker connected!");
          } else {
              Serial.print("Failed with state ");
              Serial.println(client.state());
              delay(2000);
          }
      }
      // publish and subscribe
      client.publish("boat/sensors/test", "Hello world");
      client.subscribe("boat/sensors/test");
    }

    static void callback(char *topic, byte *payload, unsigned int len){
        Serial.print("Topic: ");
        Serial.println(topic);
        Serial.print("Message: ");
        for (int i = 0; i < len; i++) {
            Serial.print((char) payload[i]);
        }
        Serial.println();
        Serial.println("-----------------------");
    }

    void publish_data(const char* topic, float payload){
      client.publish(topic, String(payload).c_str());
    }

    void subscribe_to_topic(const char* topic){
      client.subscribe(topic);
    }
};

// ----------------------------------------- TDS ----------------------------------------
#define TdsSensorPin 34
#define VREF 3.3              // analog reference voltage(Volt) of the ADC
#define SCOUNT  30            // sum of sample point

int analogBuffer[SCOUNT];     // store the analog value in the array, read from ADC
int analogBufferTemp[SCOUNT];
int analogBufferIndex = 0;
int copyIndex = 0;

float averageVoltage = 0;
float tdsValue = 0;
float temperature = 25;       // current temperature for compensation

// median filtering algorithm
int getMedianNum(int bArray[], int iFilterLen){
  int bTab[iFilterLen];
  for (byte i = 0; i<iFilterLen; i++)
  bTab[i] = bArray[i];
  int i, j, bTemp;
  for (j = 0; j < iFilterLen - 1; j++) {
    for (i = 0; i < iFilterLen - j - 1; i++) {
      if (bTab[i] > bTab[i + 1]) {
        bTemp = bTab[i];
        bTab[i] = bTab[i + 1];
        bTab[i + 1] = bTemp;
      }
    }
  }
  if ((iFilterLen & 1) > 0){
    bTemp = bTab[(iFilterLen - 1) / 2];
  }
  else {
    bTemp = (bTab[iFilterLen / 2] + bTab[iFilterLen / 2 - 1]) / 2;
  }
  return bTemp;
}

float tds(){
  static unsigned long analogSampleTimepoint = millis();
  if(millis()-analogSampleTimepoint > 40U){     //every 40 milliseconds,read the analog value from the ADC
    analogSampleTimepoint = millis();
    analogBuffer[analogBufferIndex] = analogRead(TdsSensorPin);    //read the analog value and store into the buffer
    analogBufferIndex++;
    if(analogBufferIndex == SCOUNT){ 
      analogBufferIndex = 0;
    }
  }   
  
  static unsigned long printTimepoint = millis();
  if(millis()-printTimepoint > 800U){
    printTimepoint = millis();
    for(copyIndex=0; copyIndex<SCOUNT; copyIndex++){
      analogBufferTemp[copyIndex] = analogBuffer[copyIndex];
      
      // read the analog value more stable by the median filtering algorithm, and convert to voltage value
      averageVoltage = getMedianNum(analogBufferTemp,SCOUNT) * (float)VREF / 4096.0;
      
      //temperature compensation formula: fFinalResult(25^C) = fFinalResult(current)/(1.0+0.02*(fTP-25.0)); 
      float compensationCoefficient = 1.0+0.02*(temperature-25.0);
      //temperature compensation
      float compensationVoltage=averageVoltage/compensationCoefficient;
      
      //convert voltage value to tds value
      tdsValue=(133.42*compensationVoltage*compensationVoltage*compensationVoltage - 255.86*compensationVoltage*compensationVoltage + 857.39*compensationVoltage)*0.5;

      return tdsValue;
    }
  }
}


// ----------------------------------------- DHT-11 ----------------------------------------
#define DHTTYPE DHT11
int dht_pin = 15;

// ----------------------------------------- SONAR ----------------------------------------
#define SOUND_VELOCITY 0.034
#define CM_TO_INCH 0.393701

class Sonar{
private:
  int trigPin;
  int echoPin;
  long duration;
  float distance;
public:  
  Sonar(int trigPin, int echoPin){
    this->trigPin = trigPin;
    this->echoPin = echoPin;

    pinMode(this->trigPin, OUTPUT);
    pinMode(this->echoPin, INPUT);
  }

  float calculateDistance(){
    digitalWrite(this->trigPin, LOW);
    delayMicroseconds(2);
    digitalWrite(this->trigPin, HIGH);
    delayMicroseconds(10);
    digitalWrite(this->trigPin, LOW);
  
    this->duration = pulseIn(this->echoPin, HIGH);
    this->distance = duration * SOUND_VELOCITY/2;
    
    return distance;  
  }
  
  void condition(float distance){
    Serial.print("Status: ");
    if(distance < 10){
      Serial.println("Too close! Moving away");
      return; 
    }
    Serial.println("Ok");
    return;
  }
};

// 
MQTT *mqtt;
QMC5883L compass;

void qmc(){
  int x,y,z;
  compass.read(&x,&y,&z);
  float heading = atan2(y, x);
  float declinationAngle = 0.0404;
  heading += declinationAngle;

  if(heading < 0)
    heading += 2*PI;

  if(heading > 2*PI)
    heading -= 2*PI;
    
  float headingDegrees = heading * 180/M_PI; 


  Serial.print("x: ");
  Serial.print(x);
  mqtt->publish_data("boat/sensor/qmc/x", x);
  
  Serial.print("    y: ");
  Serial.print(y);
  mqtt->publish_data("boat/sensor/qmc/y", y);
  
  Serial.print("    z: ");
  Serial.print(z);
  mqtt->publish_data("boat/sensor/qmc/z", z);
  
  Serial.print("    heading: ");
  Serial.print(heading);
  mqtt->publish_data("boat/sensor/qmc/heading", heading);
  
  Serial.print("    Radius: ");
  Serial.print(headingDegrees);
  mqtt->publish_data("boat/sensor/qmc/headingDegrees", headingDegrees);
  
  Serial.println();
  delay(100);  
}


DHT dht(dht_pin, DHTTYPE);
Sonar *sonar1;

void setup(){
  // wifi
  Serial.begin(115200);
  WiFi.disconnect(true);
  delay(1000);
  WiFi.onEvent(WiFiStationConnected, WiFiEvent_t::ARDUINO_EVENT_WIFI_STA_CONNECTED);
  WiFi.onEvent(WiFiGotIP, WiFiEvent_t::ARDUINO_EVENT_WIFI_STA_GOT_IP);
  WiFi.onEvent(WiFiStationDisconnected, WiFiEvent_t::ARDUINO_EVENT_WIFI_STA_DISCONNECTED);
  WiFi.begin(ssid, password);
  Serial.println();
  Serial.println("Wait for WiFi... ");

  // mqtt
  mqtt = new MQTT("esp32-client-", "broker.hivemq.com", 1883);

  // tds
  pinMode(TdsSensorPin,INPUT);

  // dht-11
  pinMode(dht_pin, INPUT);
  dht.begin();

  // sonar
  sonar1 = new Sonar(4,2);

  // QMC5883L
  compass.init();
}

void loop(){
  Serial.println(sonar1->calculateDistance());
  mqtt->publish_data("boat/sensor/sonar", sonar1->calculateDistance());
  
  Serial.println(dht.readTemperature());
  mqtt->publish_data("boat/sensor/temp", dht.readTemperature());
  
  Serial.println(dht.readHumidity());
  mqtt->publish_data("boat/sensor/humid", dht.readHumidity());

  Serial.print(tds());
  float tdsVal = tds();
  mqtt->publish_data("boat/sensor/tds", tdsVal);
  
  client.loop();
  delay(1000);
}