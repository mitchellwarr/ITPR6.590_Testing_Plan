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
	return "visit message"
}

func (d *driver) start(r *rand.Rand, n Network) string {
	// Initalise the location
	index := int(r.Float64())
	d.location = n.locations[index]
	return "starting"
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
