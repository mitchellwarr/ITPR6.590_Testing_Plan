package driver

import(
    "bufio"
    "fmt"
    "os"
)

func main() {
	const driverNumbers := 5
	randomGenerator := getRandomGen()
	for i:= 0; i <= driverNumbers; i++ {
		d:= driver{}
		startJourney(d)
	}
}

func getRandomGen(d *driver) {
    fmt.Print("Enter seed: ")
	input := bufio.NewReader(os.Stdin).ReadString('\n')
	seed := rand.NewSource(input)
	return rand.New(seed)
}

func startJourney(d *driver) {
	for driverInCity(d) {
		fmt.println(d.move())
	}
	fmt.println(d.visitMessage())
}

func driverInCity(d *driver) {
	return true
}

type node struct {
	name string
	paths []path
}

type path struct {
	node_01 *node
	node_02 *node
}

func (p *path) getNeighbour(n *node) {
	if n == p.node_01 {
		return p.node_02
	}
	if n == p.node_02 {
		return p.node_01
	}
	return err
}