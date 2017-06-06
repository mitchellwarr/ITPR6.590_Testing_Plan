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

```  TestStart
```
## Driver
#### ITPR6.590A2-DRV001 
Drivers can move to another node given an input  

#### ITPR6.590A2-DRV002 
Drivers will end their driving session only if they’re on an exit  

#### ITPR6.590A2-DRV003 
Test the DriverInCity method. When a driver has exited the city, their location will be "outside city"
  Preconditions: Create 5 drivers, and assign all to a different location

  ###### For each location: Example written for `Outside City`

  Execute:  call `driverInCity` method, with a location of `Outside City`  
  Message:  OutsideCity driver should not be inside city  
  Assert:   Method returns false  

#### ITPR6.590A2-DRV004
Drivers can start at a set location

### LOC - Locations
__ITPR6.590A2-LOC001__  
Node map should be loadable

__ITPR6.590A2-LOC002__  
Node map should be set up with four city locations    

__ITPR6.590A2-LOC003__  
Node map should be set up with four exit locations  
