package pkg

import (
    "bufio"
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
    "strings"
    "time"
)

func PubSubClientTest(){
        
        opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
        opts.SetClientID("go-mqtt-client-1")

        c := MQTT.NewClient(opts)

        for i := 0; i < 10; i++ {
                if token := c.Connect(); token.Wait() && token.Error() == nil {
                        break
                } else {
                        fmt.Println(token.Error())
                        time.Sleep(1 * time.Second)
                }
        }

        for {
                var message string
                fmt.Print(">> ")
                reader := bufio.NewReader(os.Stdin)
                message, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                }

                if strings.Compare(message, "\n") > 0 {

                        token := c.Publish("test/working", 0, false, message)
                        if strings.Compare(message, "bye\n") == 0 {
                                break
                        }
                        token.Wait()
                }
        }

        c.Disconnect(250)

}