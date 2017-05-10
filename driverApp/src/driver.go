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

	// Refer to FUN-OTHER-CITIES
	// Interpreted requirements, we do not display where the driver when if they did not exit via Karamu road or via Omahu Road
	// This is not good coding convention, but, this excercise is to work with testing, not coding
	if d.exitCity == "Napier" || d.exitCity == "Flaxmere" {
		message += fmt.Sprintf("Driver has gone to %v", d.exitCity)
		message += "\n"
	}

	// Refer to FUN-DASHES
	message += "-----"

	return message
}

// This method requires a rand, and not a float, because we need to continue to randomly choose a start position until they are inside the city :)
func (d *driver) start(r *rand.Rand, n Network) string {
	// Initalise the location and make sure driver starts in the city
	var index int = OutsideCityID
	for index == OutsideCityID {
		index = int(r.Float64() * float64(len(n.locations)))
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