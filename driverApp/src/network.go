package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type node struct {
	name  string
	paths []path
}

type path struct {
	node01 *node
	node02 *node
	exit   string
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

type network struct {
	locations []*node
}

type locationJSON struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type pathJSON struct {
	P1   int    `json:"p1"`
	P2   int    `json:"p2"`
	Exit string `json:"exit"`
}

type mapJSON struct {
	Locations []locationJSON `json:"locations"`
	Paths     []pathJSON     `json:"paths"`
}

func loadMapJSON(jsonInput []byte) mapJSON {
	var mapjson mapJSON
	json.Unmarshal(jsonInput, &mapjson)
	return mapjson
}

func LoadNewNetwork(filePath string) network {
	file := getFileByte(filePath)
	mapjson := loadMapJSON(file)

	newNetwork := network{locations: make([]*node, 0)}
	for i, el := range mapjson.Locations {
		newNode := &node{name: el.Name, paths: make([]path, 0)}
		newNetwork.locations = append(newNetwork.locations, newNode)
	}
	for i, el := range mapjson.Paths {
		newPath := path{
			node01: newNetwork.locations[el.P1],
			node02: newNetwork.locations[el.P2],
			exit:   el.Exit,
		}
		newNetwork.locations[el.P1].paths = append(newNetwork.locations[el.P1].paths, newPath)
		newNetwork.locations[el.P2].paths = append(newNetwork.locations[el.P2].paths, newPath)
	}
	return newNetwork
}

func getFileByte(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return file
}
