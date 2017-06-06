package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	DriverCount               = 5
	OutsideCityID             = 0
	AkinaCityID               = 2
	MapFile                   = "./map.json"
	ChanceToExit              = 0.2 // 1 divided by 5
	MessageDriverMetWithJJ    = "Driver %d met with John Jamieosn %v time(s)."
	MessageDriverNeedsHelp    = "This driver needed lots of help!"
	MessagePassengerMissedOut = "That passenger missed out!"
	MessageDriverLeftToCity   = "Driver has gone to %v"
	MessageDriverStarting     = "Driver %d starting at %v."
	MessageDriverHeading      = "Driver %d heading from %v to %v."
)

func main() {
	fmt.Print("Enter seed: ")
	input := getInput()
	randGen, err := getRandomGen(input)
	if err != nil {
		return
	}
	cityNetwork := LoadNewNetwork(MapFile)
	for i := 1; i <= DriverCount; i++ {
		d := newDriver(i)

		fmt.Println(d.start(randGen.Float64(), cityNetwork))
		// Repeat until not in the city
		for d.driverInCity() {
			fmt.Println(d.move(randGen))
		}

		fmt.Println(d.visitMessage())
		// Refer to FUN-DASHES
		fmt.Println("-----")
	}
}

func getRandomGen(input string) (*rand.Rand, error) {
	i, err := strconv.ParseInt(input, 10, 64)
	var seedInt int64
	if err == nil {
		seedInt = int64(i)
	} else {
		seedInt = 0
	}
	r := rand.New(rand.NewSource(seedInt))
	return r, err
}

func getInput() string {
	var input string
	for {
		_, err := fmt.Scanf("%v", &input)
		if err == nil {
			break
		} else {
			fmt.Print(err)
		}
	}
	return input
}
