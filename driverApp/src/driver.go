package main

import "math/rand"
import "fmt"

type driver struct {
	location   *node
	visitCount int
	driverID   int
	exitCity   string
}

func newDriver(id int) driver {
	nd := driver{}
	nd.driverID = id
	return nd
}

func (d *driver) driverInCity() bool {
	return d.location.id != OutsideCityID
}

func (d *driver) visitMessage() string {
	// Driver 4 has gone to Napier
	// Driver 4 met with John Jamieson 1 time(s).
	var message string = ""

	// Refer to FUN-AKINA-COUNT
	message += fmt.Sprintf(
		"Driver %d met with John Jamieosn %v time(s).",
		d.driverID, d.visitCount,
	)

	message += "\n"

	// Refer to FUN-AKINA-EDGES
	if d.visitCount > 2 {
		message += "This driver needed lots of help!"
		message += "\n"
	} else if d.visitCount == 0 {
		message += "That passenger missed out!"
		message += "\n"
	}	

	// Refer to FUN-DASHES
	message += "-----"

	return message
}

func (d *driver) start(r float64, n Network) string {
	// Initalise the location and make sure driver starts in the city
	var index int = OutsideCityID
	for index == OutsideCityID {
		index = int(r * float64(len(n.locations)))
	}

	d.location = n.locations[index]
	d.tryMeetJohn()
	startMessage := fmt.Sprintf(
		"Driver %d starting at %v.",
		d.driverID, n.locations[index].name,
	)
	return startMessage
}

func (d *driver) move(r *rand.Rand) string {
	from := d.location.name 
	path := d.pickPath(r.Float64())
	d.location, _ = path.getNeighbour(d.location)
	d.tryMeetJohn()
	
	if d.checkExit(r.Float64()) {
		d.location = path.exit
		d.exitCity = path.city
	}
	
	return fmt.Sprintf(
		"Driver %d heading from %v to %v.",
		d.driverID, from, d.location.name,
	)
}

func (d *driver) pickPath(r float64) path {
	index := int(r * float64(len(d.location.paths)))
	pathTaken := d.location.paths[index]
	return pathTaken
}

func (d *driver) tryMeetJohn() {
	if d.location.id == AkinaCityID {
		d.visitCount++
	}
}

func (d *driver) checkExit(r float64) bool {
	return r < ChanceToExit
}


// func (d *driver) pickPath(r float64) path {
// 	index := int(r) * len(d.location.paths)
// func (d *driver) pickNeighbour(r *rand.Rand) *node {
// 	index := int(r.Float64()) * len(d.location.paths)
// 	pathTaken := d.location.paths[index]
// 	return pathTaken
// 	location, _ := pathTaken.getNeighbour(d.location)
// 	if d.checkExit(r) {
// 		location = pathTaken.exit
// 		d.exitCity = pathTaken.city
// 	}
// 	return location
// }

// func (d *driver) pickPath(r float64) path {
// 	index := int(r) * len(d.location.paths)
// }
// func (d *driver) pickPath(r float64) path {
// 	index := int(r) * len(d.location.paths)
// 	pathTaken := d.location.paths[index]
// 	return pathTaken
// 	location, _ := pathTaken.getNeighbour(d.location)
// 	if d.checkExit(r) {
// 		location = pathTaken.exit
// 		d.exitCity = pathTaken.city
// 	}
// 	return location
// }
