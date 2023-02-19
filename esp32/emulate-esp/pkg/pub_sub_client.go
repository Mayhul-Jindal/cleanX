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
        //create a ClientOptions struct setting the broker address, clientid, turn
        //off trace output and set the default message handler
        opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
        opts.SetClientID("go-mqtt-client-1")

        //create and start a client using the above ClientOptions
        c := MQTT.NewClient(opts)

        //we are going to try connecting for max 10 times to the server if the connection fails.
        for i := 0; i < 10; i++ {
                if token := c.Connect(); token.Wait() && token.Error() == nil {
                        break
                } else {
                        fmt.Println(token.Error())
                        time.Sleep(1 * time.Second)
                }
        }

        //subscribe to the topic /go-mqtt/sample and request messages to be delivered
        //at a maximum qos of zero, wait for the receipt to confirm the subscription
       //same thing needs to go here as well.
        if token := c.Subscribe("some_topic", 0, nil); token.Wait() && token.Error() != nil {
                fmt.Println(token.Error())
                os.Exit(1)
        }

        // this is the shell where we will take input from the user and publish the message on the topic until user enters `bye`.

        for {
                var message string
                fmt.Print(">> ")
                // create a new bffer reader.
                reader := bufio.NewReader(os.Stdin)
                // read a string.
                message, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                }
                if strings.Compare(message, "\n") > 0 {
                        // if there is a message, publish it.
                        token := c.Publish("test/working", 0, false, message)
                        if strings.Compare(message, "bye\n") == 0 {
                                // if message == "bye" then exit the shell.
                                break
                        }
                        token.Wait()
                }
        }

        //unsubscribe from /go-mqtt/sample
        if token := c.Unsubscribe("some_topic"); token.Wait() && token.Error() != nil {
                fmt.Println(token.Error())
                os.Exit(1)
        }

        c.Disconnect(250)

}