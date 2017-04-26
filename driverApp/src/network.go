package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type node struct {
	name  string
	paths []path
}

type path struct {
	node01 *node
	node02 *node
}

type network struct {
}

type locationJSON struct {
	name string
	id   int
}

type pathJSON struct {
	p1   int
	p2   int
	exit string
}

type mapJSON struct {
	locations []locationJSON
	paths     []pathJSON
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

func LoadNewNetwork(filePath string) network {
	file, err := ioutil.ReadFile("./map.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype jsonobject
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
	defer file.close()
}
