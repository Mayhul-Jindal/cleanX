package pkg

type EspInternal struct{
	Temperature float64
}

type EspGps struct{
	X float64
	Y float64
}

type EspMagnetometer struct{
	X float64
	Y float64
	Z float64
}

type EspDht11 struct{
	Temperature float64
	Humidity float64
}

type EspSonar struct{
	Distance int
}

type EspWaterQuality struct{
	Tds int
	Do int
	Ph int
	Turbidity int
}