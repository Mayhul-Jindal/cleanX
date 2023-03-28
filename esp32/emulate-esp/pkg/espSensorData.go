package pkg

type EspInternal struct{
	Temperature int
}

type EspGps struct{
	X int
	Y int
}

type EspMagnetometer struct{
	X int
	Y int
	Z int
}

type EspDht11 struct{
	Temperature int
	Humidity int
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