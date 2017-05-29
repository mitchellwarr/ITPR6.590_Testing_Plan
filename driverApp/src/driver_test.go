package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDriverInCity(t *testing.T) {
	outsideCity := &node{id: 0}
	mayfair := &node{id: 1}
	akina := &node{id: 2}
	stortfordLodge := &node{id: 3}
	mahora := &node{id: 4}

	driverOutsideCity := driver{}
	driverOutsideCity.location = outsideCity

	driverMayfair := driver{}
	driverMayfair.location = mayfair

	driverAkina := driver{}
	driverAkina.location = akina

	driverStortfordLodge := driver{}
	driverStortfordLodge.location = stortfordLodge

	driverMahora := driver{}
	driverMahora.location = mahora

	boolDriverOutsideCity := driverOutsideCity.driverInCity()
	boolDriverMayfair := driverMayfair.driverInCity()
	boolDriverAkina := driverAkina.driverInCity()
	boolDriverStortfordLodge := driverStortfordLodge.driverInCity()
	boolDriverMahora := driverMahora.driverInCity()

	AssertFalse(t, "OutsideCity driver should not be inside city", boolDriverOutsideCity)
	AssertTrue(t, "Mayfair driver should be inside city", boolDriverMayfair)
	AssertTrue(t, "Akina driver should be inside city", boolDriverAkina)
	AssertTrue(t, "StortfordLodge driver should be inside city", boolDriverStortfordLodge)
	AssertTrue(t, "Mahora driver should be inside city", boolDriverMahora)
}

func TestVisitMessage(t *testing.T) {
	aDriver := driver{driverID: 1}
	aDriver.visitCount = 0
	aDriver.exitCity = ""

	expectedMessage := fmt.Sprintf(MessageDriverMetWithJJ, 1, 0) + "\n" + MessagePassengerMissedOut
	actualMessage := aDriver.visitMessage()
	AssertEqual(t, "Driver should join the correct messages", actualMessage, expectedMessage)
}

func TestVisitMessageJohn(t *testing.T) {
	driver0 := driver{driverID: 1}
	driver0.visitCount = 0
	driver1 := driver{driverID: 1}
	driver1.visitCount = 3
	driver2 := driver{driverID: 1}
	driver2.visitCount = -1

	AssertEqual(t, "Driver shows that it did not meet John", driver0.visitMessageJohn(), fmt.Sprintf(MessageDriverMetWithJJ, 1, 0))
	AssertEqual(t, "Driver shows that it meet John", driver1.visitMessageJohn(), fmt.Sprintf(MessageDriverMetWithJJ, 1, 3))
	AssertEqual(t, "Driver shows that it John less than zero", driver2.visitMessageJohn(), fmt.Sprintf(MessageDriverMetWithJJ, 1, -1))
}

func TestVisitMessageCount(t *testing.T) {
	driver0 := driver{}
	driver0.visitCount = 0
	driver1 := driver{}
	driver1.visitCount = -1
	driver2 := driver{}
	driver2.visitCount = 2
	driver3 := driver{}
	driver3.visitCount = 3

	AssertEqual(t, "Driver should show no vists", driver0.visitMessageCount(), MessagePassengerMissedOut)
	AssertEqual(t, "Driver with negative vists should show nothing", driver1.visitMessageCount(), "")
	AssertEqual(t, "Driver with few visists should show nothing", driver2.visitMessageCount(), "")
	AssertEqual(t, "Driver with many visits should need help", driver3.visitMessageCount(), MessageDriverNeedsHelp)
}

func TestVisitMessageExitCity(t *testing.T) {
	dNapier := driver{exitCity: "Napier"}
	dNothing := driver{exitCity: ""}

	AssertEqual(t, "Driver should display exiting Napier", dNapier.visitMessageExitCity(), fmt.Sprintf(MessageDriverLeftToCity, "Napier"))
	AssertEqual(t, "Driver should display no exit", dNothing.visitMessageExitCity(), "")
}

func TestStart(t *testing.T) {
	// Driver should not start outside the city
	// ID 1 0 - 0.2499,
	// ID 2 0.2500 - 0.4999
	// ID 3 0.5000 - 0.7499
	// ID 4 0.7500 - 0.9999
	testNetwork := LoadNewNetwork(MapFile)

	// Invalid inputs:
	// Negative Number
	d1 := driver{}
	d1.start(-0.5, testNetwork)

	// Number above 1
	d2 := driver{}
	d2.start(1, testNetwork)
	AssertEqual(t, "Driver is forced to location 1, when invalid input: -1", d1.location.id, 1)
	AssertEqual(t, "Driver is forced to location 1, when invalid input: 1", d2.location.id, 1)

	// Refer to FUN-AKINA-COUNT: The start method checks if you have met John even when starting at a Akina
	d3 := driver{}
	d3.start(0.3, testNetwork) // 0.3 forces starting at Akina

	AssertEqual(t, "Driver starting at Akina, will meet john", d3.visitCount, 1)

	// Starting Visit message
	d4 := driver{}
	d4Message := d4.start(0, testNetwork)
	expectedMayfair := fmt.Sprintf(MessageDriverStarting, 0, "Mayfair")
	AssertEqual(t, "Driver should visit Mayfair", d4Message, expectedMayfair)

	d4b := driver{}
	d4bMessage := d4b.start(0.24, testNetwork)
	AssertEqual(t, "Driver should visit Mayfair", d4bMessage, expectedMayfair)

	d5 := driver{}
	d5Message := d5.start(0.25, testNetwork)
	exepctedAkina := fmt.Sprintf(MessageDriverStarting, 0, "Akina")
	AssertEqual(t, "Driver should visit Akina", d5Message, exepctedAkina)

	d6 := driver{}
	d6Message := d6.start(0.5, testNetwork)
	expectedStortfordLodge := fmt.Sprintf(MessageDriverStarting, 0, "Stortford Lodge")
	AssertEqual(t, "Driver should visit Stortford Lodge", d6Message, expectedStortfordLodge)

	d7 := driver{}
	d7Message := d7.start(0.75, testNetwork)
	expectedMahora := fmt.Sprintf(MessageDriverStarting, 0, "Mahora")
	AssertEqual(t, "Driver should visit Mahora", d7Message, expectedMahora)
}

// func TestMove(){

// }

// func TestPickPath(){

// }

func TestNewDriver(t *testing.T) {
	d0 := newDriver(0)
	d4 := newDriver(4)

	AssertEqual(t, "Driver constructor should update driver ID - 0", d0.driverID, 0)
	AssertEqual(t, "Driver constructor should update driver ID - 4", d4.driverID, 4)
}

func TestTryMeetJohn(t *testing.T) {
	dAkina := driver{location: &node{id: AkinaCityID}, visitCount: 0}
	dNotAkina := driver{location: &node{id: AkinaCityID + 1}, visitCount: 0}

	dAkina.tryMeetJohn()
	dNotAkina.tryMeetJohn()

	AssertEqual(t, "Driver in akina should have met John", dAkina.visitCount, 1)
	AssertEqual(t, "Driver outside akina should not have met John", dNotAkina.visitCount, 0)
}

func TestCheckExit(t *testing.T) {
	d := driver{}

	AssertFalse(t, "Driver will not exit when number is equal to our constant 'ChanceToExit'", d.checkExit(ChanceToExit))
	AssertFalse(t, "Driver will not exit when number is boarderline more than our constant 'ChanceToExit'", d.checkExit(ChanceToExit+0.001))
	AssertFalse(t, "Driver will not exit when number is 1", d.checkExit(1))
	AssertFalse(t, "Driver will not exit when number is postivie max float", d.checkExit(math.MaxFloat64))
	AssertTrue(t, "Driver will exit when number is boarderline under the chance ", d.checkExit(ChanceToExit-0.001))
	AssertTrue(t, "Driver will exit when number is 0", d.checkExit(0))
	AssertTrue(t, "Driver will exit when number is negative max float", d.checkExit(-math.MaxFloat64))
}
