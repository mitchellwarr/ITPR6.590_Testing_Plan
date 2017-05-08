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
	} else if d.visitCount == 0 {
		message += "That passenger missed out!"
	}

	message += "\n"

	// Refer to FUN-DASHES
	message += "-----"

	return message
}

func (d *driver) start(r *rand.Rand, n Network) string {
	// Initalise the location
	var index int
	for{
		index = index = int(r.Float64()) //int(r.Float64() * float64(len(n.locations)))
		if index != OutsideCityID{
			break
		}
	}
	
	d.location = n.locations[index]
	d.tryMeetJohn()
	startMessage := "starting at " + n.locations[index].name
	return startMessage
}

func (d *driver) move(r *rand.Rand) string {
	from := d.location.name
	d.location = d.pickNeighbour(r)
	d.tryMeetJohn()
	return fmt.Sprintf(
		"Driver %d heading from %v to %v.",
		d.driverID, from, d.location.name,
	)
}

func (d *driver) pickNeighbour(r *rand.Rand) *node {
	index := int(r.Float64()) * len(d.location.paths)
	pathTaken := d.location.paths[index]
	location, _ := pathTaken.getNeighbour(d.location)
	if d.checkExit(r) {
		location = pathTaken.exit
		d.exitCity = pathTaken.city
	}
	return location
}

func (d *driver) tryMeetJohn() {
	if d.location.id == AkinaCityID {
		d.visitCount++
	}
}

func (d *driver) checkExit(r *rand.Rand) bool {
	return r.Float64() < ChanceToExit
}
