package main

import "math/rand"
import "fmt"

type driver struct {
	location   *node
	visitCount int
	driverID   int
}

func newDriver(id int) driver {
	nd := driver{}
	nd.driverID = id
	return nd
}

func (driver) driverInCity() bool {
	return true
}

func (d *driver) visitMessage() string {
	return "visit message"
}

func (d *driver) start(r *rand.Rand, n Network) string {
	return "starting"
}

func (d *driver) move(r *rand.Rand) string {
	from := d.location.name
	d.location = d.pickNeighbour(r)
	d.tryExit(r)
	d.tryMeetJohn()
	return fmt.Sprintf(
		"Driver %d heading from %v to %v.",
		d.driverID, from, d.location.name,
	)
}

func (d *driver) pickNeighbour(r *rand.Rand) *node {
	index := int(r.Float64() * 2)
	location, _ := d.location.paths[index].getNeighbour(d.location)
	return location
}

func (d *driver) tryMeetJohn() {
	if d.location.id == AkinaCityID {
		d.visitCount++
	}
}

func (d *driver) tryExit(r *rand.Rand) {
	if r.Float64() < ChanceToExit {
		d.location, _ = d.location.paths[0].getNeighbour(d.location)
	}
}
