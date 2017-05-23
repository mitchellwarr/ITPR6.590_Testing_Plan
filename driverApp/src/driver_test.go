package main

import (
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

// func TestVisitMessage(t *testing.T){
// 	// Refer to FUN-AKINA-COUNT
// 	//// Driver that has visited John
// 	aDriver := driver{}
// 	aDriver.visitCount = 0
// 	aDriver.exitCity = "nothing"

// 	expectedMessage := "Driver 1 met with John Jamieosn 0 time(s).\n
// 	That passenger missed out!"+
// 	"-----"
// 	actualMessage := aDriver.visitMessage()
// 	AssertTrue(t, "", message, expectedMessage)

// }

// func TestStart(){

// }

// func TestMove(){

// }

// func TestPickNeightbour(){

// }

// func TestTryMeetJohn(){

// }

// func TestChechExit(){

// }
