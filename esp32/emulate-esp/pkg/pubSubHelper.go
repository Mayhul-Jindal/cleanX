package pkg

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"github.com/eclipse/paho.mqtt.golang"
)

func generateRandomClientID() string {
	rand.Seed(time.Now().UnixNano())
	var id string = "mqtt-client-"
	for i := 0; i < 10; i++ {
	   randomNumber := rand.Intn(10) 
	   id += fmt.Sprintf("%v", randomNumber)
	}
	return id
 }

func StartConnection() (*mqtt.Client, string){
	brokerURI := "tcp://broker.hivemq.com:1883"
	clientID := generateRandomClientID()
	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker(brokerURI).SetClientID(clientID))

	for i := 0; i < 10; i++ {
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			fmt.Fprintf(os.Stderr, "Connection failed: %v\n", token.Error())
			time.Sleep(1 * time.Second)
		}else{
			fmt.Fprintf(os.Stdout, "Connection success\n")
			break
		}
	}
	
	return &mqttClient, clientID
}

func PublishData(mqttClient *mqtt.Client, topic string, qos byte, retain bool, data any){
	payload := fmt.Sprintf("%+v", data) 
	
	if token := (*mqttClient).Publish(topic, qos, retain, payload); token.Wait() && token.Error() != nil {
        fmt.Fprintf(os.Stderr, "Publishing of Topic: %v failed: %v\n", topic, token.Error())
	}else{
		fmt.Fprintf(os.Stdout, "Publishing of Topic: %v Payload: %v success\n", topic, payload)
	}
}