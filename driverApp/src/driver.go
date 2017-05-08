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

func (d *driver) start(r float64, n Network) string {
	// Initalise the location
	index := int(r)
	d.location = n.locations[index]
	return "starting"
}

func (d *driver) move(r *rand.Rand) string {
	from := d.location.name
	path := d.pickPath(r.Float64())
	d.location, _ = path.getNeighbour(d.location)
	if d.checkExit(r.Float64()) {
		d.location = path.exit
		d.exitCity = path.city
	}
	d.tryMeetJohn()
	return fmt.Sprintf(
		"Driver %d heading from %v to %v.",
		d.driverID, from, d.location.name,
	)
}

func (d *driver) pickPath(r float64) path {
	index := int(r) * len(d.location.paths)
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
