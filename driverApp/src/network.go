package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type node struct {
	name  string
	id    int
	paths []path
}

type path struct {
	node01 *node
	node02 *node
	exit   *node
	name   string
	city   string
}

func (p *path) getNeighbour(n *node) (*node, error) {
	if n == p.node01 {
		return p.node02, nil
	}
	if n == p.node02 {
		return p.node01, nil
	}
	return nil, errors.New("Path does not contain node")
}

type Network struct {
	locations []*node
}

type mapJSON struct {
	Locations []locationJSON `json:"locations"`
	Paths     []pathJSON     `json:"paths"`
}

type locationJSON struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type pathJSON struct {
	P1   int    `json:"p1"`
	P2   int    `json:"p2"`
	Name string `json:"name"`
	City string `json:"city"`
}

func LoadNewNetwork(filePath string) Network {
	file, _ := getFileByte(filePath)
	mapjson := loadMapJSON(file)

	newNetwork := Network{
		locations: make([]*node, 0),
	}
	for _, el := range mapjson.Locations {
		newNode := &node{
			id:    el.ID,
			name:  el.Name,
			paths: make([]path, 0),
		}
		newNetwork.locations = append(newNetwork.locations, newNode)
	}
	for _, el := range mapjson.Paths {
		newPath := path{
			node01: newNetwork.locations[el.P1],
			node02: newNetwork.locations[el.P2],
			city:   el.City,
			name:   el.Name,
			exit:   newNetwork.locations[0],
		}
		newNetwork.locations[el.P1].paths = append(newNetwork.locations[el.P1].paths, newPath)
		newNetwork.locations[el.P2].paths = append(newNetwork.locations[el.P2].paths, newPath)
	}
	return newNetwork
}

func loadMapJSON(jsonInput []byte) mapJSON {
	var mapjson mapJSON
	json.Unmarshal(jsonInput, &mapjson)
	return mapjson
}

func getFileByte(filePath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
