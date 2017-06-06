# Unit tests

## System
#### ITPR6.590A2-SYS001
The random generator can be set by an input  
	
  ##### Valid input
  Execute: call `getRandomGen` method, with an input of `100`  
  Message: Random gen with seed 100 should be 0.8165026937796166  
  Assert:  getRandomGen returns `0.8165026937796166`  
  
  ##### Input as a string
  Execute:  call `getRandomGen` method, with an input of `one hundred`  
  Message:  Random gen with error gen should be 0.9451961492941164  
  Assert:  getRandomGen returns `0.8165026937796166` because 0 is actually uses.    
  
  Message:  Random gen should return an error  
  Assert:  Error is not null.  
  
#### ITPR6.590A2-SYS002
The system should iterate through 5 drivers  
  
#### ITPR6.590A2-SYS003
The system should print on screen appropriate messages  

  ##### Refer to FUN-AKINA-EDGES
  Message: Driver should show no vists  
  Assert: Message says that "That passenger missed out!"  
  
  Message: Driver with negative vists should show nothing  
  Assert: Message is blank  
  
  Message: Driver with few visists should show nothing  
  Assert: Message is blank  
  
  Message: Driver with many visits should need help  
  Assert: Message says that "This driver needed lots of help!"  

  ###### Refer to FUN-AKINA-COUNT
  Message: Driver shows that it did not meet John  
  Assert: Visit message is formatted correctly and correct data 0 times  
  
  Message: Driver shows that it meet John  
  Assert: Visit message is formatted correctly and correct data 3 times     
  
  Message:  Driver shows that it John less than zero  
  Assert: Visit message is formatted correctly and correct data -1 times     
  
  Message: Driver should display exiting Napier  
  Assert: Driver has gone to Napier  
  
  Message: Driver should display no exit  
  Assert: No message is displayed when driver is exiting not from napier  
  
  ###### Refer to FUN-OTHER-CITIES
  Execute:  call `driverInCity` method, with a location of `Outside City`    
  Message:  OutsideCity driver should not be inside city  
  Assert:   Method returns false  

## Driver
#### ITPR6.590A2-DRV001 
Drivers can move to another node given an input  

  Execute: call `move` method,
  Message:  
  Assert:   
  
#### ITPR6.590A2-DRV002 
Drivers will end their driving session only if they’re on an exit  
  Execute: call `move` method,
  Message:  
  Assert:   

#### ITPR6.590A2-DRV003 
Test the DriverInCity method. When a driver has exited the city, their location will be "outside city"  
  ###### For each location: Example written for `Outside City`
  Preconditions: Create 5 drivers, and assign all to a different location  

  Execute:  call `driverInCity` method, with a location of `Outside City`  
  Message:  OutsideCity driver should not be inside city  
  Assert:   Method returns false  

#### ITPR6.590A2-DRV004
Drivers can start at a set location

  Execute:  call `Start` method, with a a seed of -1  
  Message:  Driver is forced to location 1, when invalid input: -1  
  Assert:   Location id is 1  
   
  Execute:  call `Start` method, with a a seed of 1  
  Message:  Driver is forced to location 1, when valid input: 1  
  Assert:   Location id is 1   


### LOC - Locations
__ITPR6.590A2-LOC001__  
Node map should be loadable

Execute:  call `getFileByte` method, with an invalid file  
Message: Non existant file should return an err   
Assert:  The `error` returned from the `getFileByte()` method is not `null` 

Execute:  call `getFileByte` method, with an empty file  
Message: File should be returned as bytes  
Assert: The `bytes` returned from the `getFileByte()` method is not `null` 

__ITPR6.590A2-LOC002__  
Node map should be set up with four city locations    

Execute:  call `getNeighbour` method, with an a range of nodes
Message: Node01 should be neighbour to Node02  
Assert: The node returned from the `getNeighbour()` method for the Node02 is Node01  

Message: Node02 should be neighbour to Node01  
Assert: The node returned from the `getNeighbour()` method for the Node02 is Node02 

Message: Node03 should be neighbour to none01  
Assert: The `error` returned from the `getNeighbour()` method is not `null`  

__ITPR6.590A2-LOC003__  
Node map should be set up with four exit locations  

Execute:  call `LoadNewNetwork` method, with an invalid file  
Message: Test Network should not be empty     
Assert: Number of locations of network is no 0

Execute:  call `LoadNewNetwork` method, with an empty file  
Message: Empty file should return an empty network  
Assert: Number of locations of network is 0
