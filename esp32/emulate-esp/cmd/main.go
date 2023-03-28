package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Mayhul-Jindal/cleanX/esp32/emulate-esp/pkg"
)

func main() {
	mqttClient, clientID := pkg.StartConnection("tcp://broker.hivemq.com:1883"); 
	if mqttClient == nil{
		os.Exit(1)
	}

	fmt.Println("--------------------------------------------------------")
	// represents one esp32
	func (){
		for{
			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/internTemp", 0, false, pkg.EspInternal{
				Temperature: rTemp(),
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
				Temperature: rTemp1(),
				Humidity: rHumid(),
			})

			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/dist", 0, false, pkg.EspSonar{
				Distance: 100,
			})

			pkg.PublishData(mqttClient, "ssv/boat/" + clientID + "/waterQuality", 0, false, pkg.EspWaterQuality{
				Do: rDo(),
				Ph: rPh(),
				Turbidity: rTurbidity(),
				Tds: rTds(),
			})

			fmt.Println("--------------------------------------------------------")
			time.Sleep(2 * time.Second)
		}
	}()
}

func rTemp() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(5) + 60
}

func rTemp1() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) + 27
}

func rHumid() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(3) + 60
}

func rDo() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) + 6
}

func rPh() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(14) + 0
}

func rTurbidity() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(5) + 1
}

func rTds() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(100) + 50
}