package main

import (
	"fmt"
	"log"

	"github.com/d2r2/go-dht"
)

type (
	// or centigrade
	Celsius    float32
	Kelvin     float32
	Fahrenheit float32
)

func Celsius2Fahrenheit(c float32) float32 {
	return float32(c*9/5 + 32)
}

func main() {
	// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
	// You may enable "boost GPIO performance" parameter, if your device is old
	// as Raspberry PI 1 (this will require root privileges). You can switch off
	// "boost GPIO performance" parameter for old devices, but it may increase
	// retry attempts. Play with this parameter.
	// Note: "boost GPIO performance" parameter is not work anymore from some
	// specific Go release. Never put true value here.
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		Celsius2Fahrenheit(temperature), humidity, retried)
}
