# Interpreted_Requirements

## Map
Figure one shows an interpretation of the design of the map. 
![Map of city](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/20170406_151857%5B1%5D.jpg)
Figure One

The driver will decide where to begin in the map using our random seed.
The driver will decide what his next place to visit will be using a probablity calculation.
The driver will have a specified chance of taking an exit over another location. This means that most of the time, the driver will continue to drive to the next loction in the city. The specified chance has not been set and will require tweaking. Initially it will start at 20%.

## Seed
In order to fulfil the psuedo randomness of the requirements, the random genertor used needs to be initialized with a user inputted seed. This seed will set up the random generator. This means that using the same seed will make the program always produce the same output.  

## Reporting
This application is console based. There is no specifications on how the driver will report where it in on the map. Therefore, the driver will send string messages to the system which will print on screen. These returned string messages can be taken from the driver and tested.  

## Location network
In order to replicate the city network of paths and locations, a graph routing system will be used. this means that several nodes will be used to stand for each location. In the code paths will be used to describe how all the nodes interconnect. Each node will have an id that correlates to a location on the map. In this way we can work out which node is what location, and test that a driver is in a certain location.  
The driver will use the nodes paths to work out which neighbour to move to.  

## Exit locations
Upon exiting through Karamu road then the program shall display that the driver has gone to Napier, if the driver exits by Omahu Road, then the program shall display that the driver has gone to Flaxmere. If the driver exits through any other road then we dont display any message.
