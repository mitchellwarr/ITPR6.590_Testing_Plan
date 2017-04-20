package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	const driverNumbers = 5
	fmt.Print("Enter seed: ")
	var input string
	_, err := fmt.Scanf("%v", &input)
	if err != nil {
		fmt.Print(err)
	}
	// randGen, _ := getRandomGen(input)
	network = loadNewNetwork("./map.json")
	// for i := 0; i <= driverNumbers; i++ {
	// 	d := driver{}
	// 	startJourney(d)
	// }
}

func getRandomGen(input string) (*rand.Rand, error) {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		s := rand.NewSource(0)
		return rand.New(s), err
	}
	s := rand.NewSource(i)
	return rand.New(s), nil
}

// func startJourney(d *driver) {
// 	for driverInCity(d) {
// 		fmt.println(d.move())
// 	}
// 	fmt.println(d.visitMessage())
// }

// func driverInCity(d *driver) {
// 	return true
// }
