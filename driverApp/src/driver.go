package main

import "math/rand"

type driver struct {
	location   *node
	visitCount int
}

func (d *driver) Move(r *rand.Rand) *node {
	index := int(r.Float64() * 2)
	d.location, _ = d.location.paths[index].getNeighbour(d.location)
	if d.location.name == "a place" {
		d.visitCount++
	}
	return d.location
}

func (d *driver) VisitMessage() string {
	return "visit message"
}
