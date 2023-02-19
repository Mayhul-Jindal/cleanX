package pkg

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type ESP_Data struct{
	Temperature float32
}

type GPS_Data struct{
	X float64
	Y float64
}

type Sensor_Data struct{
	Temperature float32
	Tds float32
	Sonar int
	Gps GPS_Data
}

func Emulate(){
	rand.Seed(time.Now().UnixNano())
	
	espDataTest := ESP_Data{
		Temperature: rand.Float32(),
	}
	espDataTestJSON, _ := json.Marshal(espDataTest);

	sensorDataTest := Sensor_Data{
		Temperature: float32(rand.Intn(100)),
		Tds: float32(rand.Intn(200)),
		Sonar: rand.Intn(50),
		Gps: GPS_Data{
			X: 45,
			Y: 78,
		},
	}
	sensorDataTestJSON, _ := json.Marshal(sensorDataTest);

	fmt.Println(string(espDataTestJSON))
	fmt.Printf("Esp Data: %s\nSensor Data: %s", string(espDataTestJSON), string(sensorDataTestJSON))
}