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
func (d *driver) visitMessageJohn() string {
	// Refer to FUN-AKINA-COUNT
	return fmt.Sprintf(
		"Driver %d met with John Jamieosn %v time(s).",
		d.driverID, d.visitCount,
	)
}

func (d *driver) visitMessageCount() string {
	// Refer to FUN-AKINA-EDGES
	if d.visitCount > 2 {
		return "This driver needed lots of help!"
	} else if d.visitCount == 0 {
		return "That passenger missed out!"
	}
	return ""
}

func (d *driver) visitMessageExitCity() string {
	// Refer to FUN-OTHER-CITIES
	// Interpreted requirements, we do not display where the driver when if they did not exit via Karamu road or via Omahu Road
	// This is not good coding convention, but, this excercise is to work with testing, not coding
	if d.exitCity != "" {
		return fmt.Sprintf("Driver has gone to %v", d.exitCity)
	}
}

func (d *driver) visitMessage() string {
	newline := "\n"
	john := visitMessageJohn()
	if john != ""{
		john += newline
	}
	count := visitMessageCount()
	exit := visitMessageExitCity()



	// Refer to FUN-DASHES
	dashes := "-----"

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