package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	DriverCount   = 5
	OutsideCityID = 0
	AkinaCityID   = 2
	MapFile       = "./map.json"
	ChanceToExit  = 1 / 5
)

func main() {
	fmt.Print("Enter seed: ")
	input := getInput()
	randGen, _ := getRandomGen(input)
	cityNetwork := LoadNewNetwork(MapFile)
	for i := 1; i <= DriverCount; i++ {
	for i := 0; i <= DriverCount; i++ {
	// lower one wrong
	}
		d := newDriver(i)

<<<<<<< HEAD
		fmt.Println(d.start(randGen, cityNetwork))
=======
		fmt.Println(d.start(randGen.Float64(), cityNetwork))
>>>>>>> e69fecec3cc3a61a9593f1bb2cd3ade6e4e2d63f

		// Repeat until not in the city
		for d.driverInCity(){
		if d.driverInCity() {
		// if is wrong
		}
			fmt.Println(d.move(randGen))
		}

		fmt.Println(d.visitMessage())
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
		}
		fmt.Print(err)
	}
	return input
}
