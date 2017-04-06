## High Level Component Tests
### LOC - Locations
__ITPR6.590A2-LOC001__  
Node map should be loadable

	Execution:  
	Successfully retrieve JSON map file  
	Load JSON map file

	Postconditions:  
	Test that all nodes have been created  

__ITPR6.590A2-LOC002__  
Node map should be set up with four city locations  

	Preconditions:  
	Load map JSON file  

	Execution:  
	Build Locations from file  
	Retrieve node network  

	Postconditions:  
	Test that there are four city nodes	  

__ITPR6.590A2-LOC003__  
Node map should be set up with four exit locations  

	Preconditions:  
	Load map JSON file  

	Execution:  
	Build Locations from file  
	Retrieve node network  

	Postconditions:  
	Test that there are four exit nodes  

### DRV - Driver  

__ITPR6.590A2-DRV001__  
Drivers can move to another node given an input  

	Precondition:  
	Build driver and network  

	Execution:  
	Set drivers starting path  
	Move driver with float  

	Postconditions:  
	Check that driver has moved to the right node  

__ITPR6.590A2-DRV002__  
Drivers will end their driving session only if they’re on an exit  

	Precondition:  
	Build driver and network  

	Execution:  
	Set drivers starting path  
	Move driver to a city location  
	Move driver to an exit  

	Postconditions:  
	Check that the driver is still driving on city locations  
	Check that the driver stops driving at the exit  

__ITPR6.590A2-DRV003__  
Drivers can start at a set location  

	Precondition:  
	Build driver and network  
	Pick a seeded number  

	Execution:  
	Set drivers starting location with the seeded number  

	Postconditions:  
	Check that the driver is at the expected starting location  

### SYS - System  

__ITPR6.590A2-SYS001__  
The random generator can be set by an input  

	Precondition:  
	Build system  

	Execution:  
	Set seed input value  
	Get next random number  

	Postconditions:  
	Check random number against expected  

__ITPR6.590A2-SYS002__  
The system should iterate through 5 drivers  

	Precondition:  
	Build system  

	Execution:  
	Set seed  
	Run system  

	Postconditions:  
	Check that there have only been 5 drivers  

__ITPR6.590A2-SYS003__  
The system should print on screen appropriate messages  

	Precondition:  
	Build system  

	Execution:  
	Set seed  
	Run system  
	Save messages returned by drivers  

	Postconditions:  
	Check that the messages sent by the drivers are expected  
