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
	cityNetwork := loadNewNetwork(MapFile)
	for i := 0; i <= DriverCount; i++ {
		d := newDriver(i)
		fmt.Println(d.start(randGen, cityNetwork))
		for d.driverInCity() {
			fmt.Println(d.move(randGen))
		}
		fmt.Println(d.visitMessage())
	}
}

func getRandomGen(input string) (*rand.Rand, error) {
	i, err := strconv.ParseInt(input, 10, 64)
	s := rand.NewSource(0)
	if err == nil {
		s = rand.NewSource(i)
	}
	return rand.New(s), err
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
