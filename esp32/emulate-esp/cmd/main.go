package main

import (
	"fmt"
	"time"

	"github.com/Mayhul-Jindal/cleanX/esp32/emulate-esp/pkg"
)

func main() {
	mqttClient, clientID := pkg.StartConnection()

	fmt.Println("--------------------------------------------------------")
	// represents one esp32
	func (){
		for{
			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/intenTemp", 0, false, pkg.EspInternal{
				Temperature: 10,
			})
	
			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/gps", 0, false, pkg.EspGps{
				X: 10,
				Y: 100,
			})

			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/magnet", 0, false, pkg.EspMagnetometer{
				X: 10,
				Y: 100,
				Z: 1000,
			})
	
			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/tempnhum", 0, false, pkg.EspDht11{
				Temperature: 234,
				Humidity: 2342,
			})

			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/dist", 0, false, pkg.EspSonar{
				Distance: 100,
			})

			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/waterQuality", 0, false, pkg.EspWaterQuality{
				Do: 124,
				Ph: 123,
				Turbidity: 243,
				Tds: 12,
			})

			fmt.Println("--------------------------------------------------------")
			time.Sleep(2 * time.Second)
		}
	}()
}
