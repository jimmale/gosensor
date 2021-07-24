package main

/*
This application is periodically used by telegraf to query the temperature sensors.

An example output is as follows:

temperature-sensor,chip=28-00000a69b18c temperature=16.88 1603911540198853959
temperature-sensor,chip=28-00000a6999ff temperature=-22.19 1603911540198853959
*/

import (
	_ "embed"
	"fmt"
	"github.com/yryz/ds18b20"
	"os"
	"time"
)

//go:embed LICENSE
var LICENSE string

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--license" || os.Args[1] == "--licence" {
			fmt.Println("This application contains libraries from http://github.com/yryz/ds18b20 and is")
			fmt.Println("licensed to you under the terms of the Apache License:\n")
			fmt.Println("================================================================================")
			fmt.Println(LICENSE)
			fmt.Println("================================================================================")
			os.Exit(0)
		}
	}

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
