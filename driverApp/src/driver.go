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
	pathTaken := d.pickNeighbourPath(r)
	d.location, _ = pathTaken.getNeighbour(d.location)
	d.tryExit(r)
	d.tryMeetJohn()
	return fmt.Sprintf(
		"Driver %d heading from %v to %v via [%v].",
		d.driverID, from, d.location.name, pathTaken.exit,
	)
}

func (d *driver) pickNeighbourPath(r *rand.Rand) path {
	index := int(r.Float64() * 2)
	return d.location.paths[index]
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
