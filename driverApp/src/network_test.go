package main

import (
	"testing"
)

func TestGetNeighbour(t *testing.T) {
	node01 := &node{id: 0}
	node02 := &node{id: 1}
	node03 := &node{id: 3}
	paths := path{node01: node01, node02: node02}
	node01.paths = append(make([]path, 0), paths)
	node02.paths = append(make([]path, 0), paths)

	n1, _ := paths.getNeighbour(node02)
	n2, _ := paths.getNeighbour(node01)
	_, n3 := paths.getNeighbour(node03)

	AssertEqual(t, "Node01 should be neighbour to Node02", n1, node01)
	AssertEqual(t, "Node02 should be neighbour to Node01", n2, node02)
	AssertNotEqual(t, "Node03 should be neighbour to none", n3, nil)
}

func TestGetFileByte(t *testing.T) {
	_, err := getFileByte("empty")
	bytes, _ := getFileByte(MapFile)

	AssertNotEqual(t, "Non existant file should return an err", err, nil)
	AssertNotEqual(t, "File should be returned as bytes", bytes, nil)
}

func TestLoadNewNetwork(t *testing.T) {
	testNetwork := LoadNewNetwork(MapFile)
	emptyNetwork := LoadNewNetwork("empty")

	AssertNotEqual(t, "Test Network should not be empty", len(testNetwork.locations), 0)
	AssertEqual(t, "Empty file should return an empty network", len(emptyNetwork.locations), 0)
}
