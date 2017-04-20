package main

import (
	"errors"
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

func (p *path) getNeighbour(n *node) (*node, error) {
	if n == p.node01 {
		return p.node02, nil
	}
	if n == p.node02 {
		return p.node01, nil
	}
	return nil, errors.New("Path does not contain node")
}

func loadNewNetwork(filePath string) network {

}
