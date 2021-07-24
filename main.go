package main

/*
This application is periodically used by telegraf to query the temperature sensors.

An example output is as follows:

temperature-sensor,chip=28-00000a69b18c temperature=16.88 1603911540198853959
temperature-sensor,chip=28-00000a6999ff temperature=-22.19 1603911540198853959
*/

import (
	"fmt"
	"github.com/yryz/ds18b20"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			fmt.Printf("temperature-sensor,chip=%s temperature=%.2f %d\n", sensor, t, now)
		}
	}
}
