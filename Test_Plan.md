# ITPR6590 - 0A23 TEST PLAN

## Introduction
This master level test plan is designed to set the requirements for the Hastings Driving Application.

### Scope
The scope of this system exists within the functional requirements provided in the requirements documentation. This plan will define how to test that the program has met all of its requirements, as well as lining out the individual unit tests that will define the underlining behaviour. The test plan will not test external modules, hardware or the Go language. The tests created by the testers will not be tested. The test plan will not be used to report on the efficiently or performance of the system in any environment.

### Goals
The program is intended to simulate a driver in hastings, tryign to pick up John Jamieson in Akina. It is a very simple and abstracted simulation. The goal of this test plan is to outline how these simulation requirements for the program, translate into testable scenarios, and testable code. It will be impossible to cover every last scenario, however this plan will outline key scenarios and equivalance classes that cover the most important aspects of the user experience. It is also intended to balance high test coverage, with low time and code investments spent into creating these tests.  
This goal will be achieved by selecting important equivalence classes to cover a large number of inputs; as well as selecting the most important, and likely, edge cases that are made possible by the code structure.  
The top priority is to prevent errors affecting the end user. The testing will be put in place to ensure the mitigation of the program suffering from an unhandled run-time error. 

### Constraints and Limitations
It is not intended that everything possibility will be covered by the described tests and methods. Time investment poses as the biggest limitation on both the construction of this plan, and on the breadth of tests that can be put in place for the development of the program.  
This plan is also focused on a skill level befitting students. It is intended to not cover areas that fall outside the range that a student could be expected to consider: Both in the depth of documentation, and in the tests written.  

### Budget constraints
The size of this plan, and the program created off the resulting size of this plans description, will be developed within a very tight time constraint. In terms of time, the team has only a small amount to spare on this project. The developed document and tested application are to be delivered before the due date of 8th June 2017. The team has been planned to work for only a few hours a week.  
In terms of cost, with each team member working under a conservative $30 an hour, this project is likely to take $4,000 to produce. However, under the budget, money will be saved by the fact that the team is made up of students, who are working for free and under their own time.  
Overall the development will not incurr any expenses, and the budget of the development process is $0. However, the project does create a significant time investment that will keep the program small.

### Pass and Fail Criteria
The ultimate goal of the test phase is to ensure that no defects are visible to the user. Once the program has meet all tests outlined in this plan, the program will be said to meet the requirements. The tests will therefore server as the baseline, against which; confirmation can be drawn that the program has finished and sastifies the business requiremnts. The test phase will be successful when the below criteria is meet:
#### Unit Test Level
* All functional requirements have a test case
* All test cases have a unit test
* 100% of unit test assertions pass
#### Master Test Level
* All unit tests pass
* No additional defects of the program can be found through typical user use cases

## Test items
Please see the high level testing [Component_Test_Plan](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/Component_Test_Plan.md) 

# ToDo
```
## Features to be tested
features from users perspective  
links back to requirements  
  
5 drivers
messages
changing location randomly
exiting system
starting randomly
```
## Features not to be tested
In order to fulfil the feature outlining the psuedo-random nature of the program, a random generation library needs to be used. The effectiveness of this library is not to be tested. To which, it will not be tested that each time you call the random generator you get a different and random value in return. This is a third party library and therefore, is out of scope. What will be tested is on a high level, that given a seed, the program will always produce the same result.

## Approach to testing
The testing method selected is a form of TDD. There will be a focus on unit tests that define the structure of the code, however these tests will only cover the set requirements. This means that only enough tests will be planned so that it can be proved when the program meets all requirements. Further exploritory tests will be done to find any code breaking bugs. Once found, unit tests will be written to ensure these bugs don't reamerge.  

Unit tests will be written before the cods is written. However, the code structure will be explored and decided on first. Only after the structure of classes and their relationships are set, are unit tests written for the requirements, and then methods written to sastify those tests.  

The testing framework used in the programs development will be the native "Testing" framework available in the Go language.  

To ensure that the code tested not only fulfils the requirements, but will continue to do so under a variety of inputs; edge cases, boundry cases, and possible equivalence class combinations will be tested as well, where possible. This helps to ensure that the program, even though satisfying the requirements, will not then become faulty under unforseen usage scenarios.    

The tests will produce a result or either, “Pass” or “Fail”. The tests that result as “Pass” will not be developed further, and are considered to meet the Hastings Driver Application initial requirements. The requirements that “Fail” will be marked as not behaving as expected. These failed requirements will be returned to development to ensure the test eventually passes. To comunicate between the testing process and coding process of development, all issues will be tracked as they fail. A new issue will be logged describing what aspect was being tested, and once the development process has gotten the code to satisfy the test; the issue will be closed and marked as resolved. If a previously resolved issue reapears, it will be marked as a new issue.   

## Test Deliverables
* Test plan document.
  * Test Features and items 
  * Sample test data
  * Findings report

## Development Team
### Training Needs
All testers need to have an understanding of the functional and non-functional requirements of the application. All testers need to have knowledge of Go and the testing package built into Go. Staff will need to be familar with the [to-do list software](http://abstractspoon.weebly.com/), and [GitHub](https://github.com/).

### Responsibilities
The testing team is composed of two members being Sarah Anderson, and Mitchell Warr. All roles will be shared amung the entire team. However, each of the members will take a more focused lead in certain roles. Mitchell will focus on the development of the program outlined in the test plan, and outlining the component test plan. Sarah will focus on creating the drafts for test plan documentation, and documenting and outlining the individual test cases. Both team members will share in writing and creating unit tests, and finishing all written documentation. Both Sarah and Mitch will be responsible for all components and deliverables, with each having the responibility to overlook and polish all aspects of the deliverables. It is intended that the test design is shared equally, while the implementation is split between members.  

## Schedule
The schedule will be kept up to date using [to do list](http://abstractspoon.weebly.com/) and uploaded to the [git repo](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/SoftwareTesting.tdl). 
Sarah and Mitch will meet once a week to disucss and implement the assignment applcaition.

## Testing Tasks
### Test scenarios
Test scenarios are documented on in our [Test scenarios](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/test_case_scenrios)

### Test Scrips
Test scripts will be inside our [driver application](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/driverApp) 

### Issue tracking
Issue tracking will be kept on [github issue tracker](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/issues)

## Envrionmental needs 
The application is developed in Go. Set up on the envrioment is documented on [GoLang's website](https://golang.org/doc/install) 

## Risks and contingencies
* The application does not compile and thus no tests are able to be run.
  * Develop in a TDD style where the application remains at a managable state
  * Keep the repositiory up-to-date with good mile stones, so a roll back can be done.
* The development becomes behind schedule
  * Maintain communication between team members (through telegram and github)




