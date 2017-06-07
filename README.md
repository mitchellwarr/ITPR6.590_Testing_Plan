# ITPR6.590_Testing_Plan

Work produced by Mitchell Warr and Sarah Anderson for the EIT's Software Testing.   
Feb 2017 - June 2017. 

##### Assignment 2
###### Executive Summary 
- [x] [Executive Summary](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/Executive_Summary.md)  

###### Activity log 
- [x] [ToDo list](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/SoftwareTesting.tdl)  
- [x] [Commit log](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/commits/master)  

###### Test plan 
- [x] [Test Plan](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/Test_Plan.md)  

##### Assignment 3
###### Findings Report:
- [x] [Findings report](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/testing_report.md)  
- [x] [Evidence of the tests](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/test_results.PNG)  

###### Test documentation:
- [x] [Unit](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/unit_tests.md)
- [x] [Component](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/Component_Test_Plan.md)

###### Actual tests:
- [x] [main_test.go](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/driverApp/src/main_test.go)  
- [x] [driver_test.go](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/driverApp/src/driver_test.go)  
- [x] [network_test.go](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/driverApp/src/network_test.go)  



## CitySim2017 Requirements Specification

_FUN-CITY-LOCS_. The program shall loosly simulate the Hastings city with four locations: **Mayfair**, **Akina**, **Stortford lodge**, and **Mahora**.

_FUN-OUTSIDE-CITY_. A fifth location, Outside City, shall stand for a driver outside the city limits.

_FUN-EXITS_. The city has four main exits, Karamu Road, Havelock Road, Railway road and Omahu Road. (refer to the map below)

_FUN-STREETS_. The are a number of streets that connect up _FUN-CITY-LOCS_ to the _FUN-EXITS_. Select a path using one or more streets that link up each _FUN-EXITS_ with a _FUN-CITY-LOCS_ location in such a way that a driver can traverse from one _FUN-EXITS_ through a _FUN-CITY-LOCS_ via the chosen streets to another _FUN-CITY-LOCS_ and exit on another _FUN-EXITS_. One-way streets are also permitted.

_FUN-FIVE-DRIVERS_. Five drivers, numbered 1 through 5, shall traverse the city, one after the other.

_FUN-AKINA-COUNT_. At the end of each drive, the simulation shall print out how many times that driver visited **Akina**.  The format shall be "Driver `n` met with John Jamieosn `x` time(s).", where `n` is the driver number (1 through 5) and `x` is the number of times that the driver visited the **Akina** location.  If a driver starts on the **Akina** location, this counts as a time visiting the **Akina**.  

_FUN-AKINA-EDGES_: If a driver visited **Akina** three or more times, then an additional line shall be printed stating "This driver needed lots of help!" If a driver never visits **Akina**, the additional line shall be printed stating "That passenger missed out!"  If a driver visited **Akina** one or two times, then no additional line shall be printed.  This shall be the last line printed out for each iteration.

_FUN-START-LOC_. A driver may start in any of the four locations listed in _FUN-CITY-LOCS_.  Drivers may not start outside of the city.

_FUN-ITERATION_. At each iteration, a Driver will drive from the current Location to one of the possible Locations that can be reached from the original Location.  The decision shall be made pseudorandomly based on a seed passed in from the command line, as covered by _FUN-ARGS_.

_FUN-END_. The simulation shall end if a Driver encounters the Outside City Location.

_FUN-ARGS_. The system shall accept an integer seed from the command line for the pseudorandom number generator.  No other arguments shall be accepted.  If no arguments are provided,  more than one argument is provided, or the single argument is not an integer, the system shall inform the user and exit.

_FUN-OTHER-CITIES_. If a driver exits the city via Karamu road, then the program shall display that the driver has gone to Napier.  If a driver exits the city via Omahu Road, then the program shall display that the driver has gone to Flaxmere.

_FUN-DASHES_. After each iteration's display on the screen, a line of five dashes (i.e., `-----`) shall be printed.  This line of dashes shall occur after all information from that iteration has been printed out.

It may be easier to understand the map of the city visually.

```
Hastings City Map	
https://www.google.co.nz/maps/@-39.6407753,176.8411868,15z

Hastings suburbs
https://upload.wikimedia.org/wikipedia/commons/5/5a/Hastings%2C_New_Zealand_numbered_suburbs_map.png
