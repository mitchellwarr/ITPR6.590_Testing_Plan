# ITPR6590 - 0A23 TEST PLAN

## Introduction
This master level test plan is designed to set the requirements for the Hastings Driving Application.

### Scope
The scope of this system exists within the functional requirements provided in the requirements documentation. This plan will define how to test that the program has met all of its requirements, as well as lining out the individual unit tests that will define the underlining behaviour. The test plan will not test external modules, hardware or the Go language. The tests created by the testers will not be tested. The test plan will not be used to report on the efficiently or performance of the system in any environment.

### Goals
The goal of the test plan is to cover as many scenarios as possible, with the least amount of tests. We will achieve this goal by using equivalence classes, to cover a large number of inputs.
The top priority of the test plan is to prevent errors affecting the end user. The testing will be put in place to ensure the mitigation of the program suffering from an unhandled run-time error. 

### Constraints and Limitations
The test plan cannot be 100% full coverage. The test plan is limited by the tester’s knowledge of the system, the time it takes for tests to be created, and the creativity of the testers. The tests cannot be fully comprehensive as computers are limited by the computation time of running a test.

### Budget constraints
The development will be done under the constaint of time. The developed, documented and tested application will be delivered before the due date of 8th June 2017. The development will not incurr any expenses, and the budget of the development process is $0.

### Pass and Fail Criteria
The ultimate goal of the test phase is to ensure that no defects are visible to the user. The test phase will be successful when the below criteria is meet:
#### Unit Test Level
* All functional requirements have a test case
* All test cases have a unit test
* 100% of unit test assertions pass
#### Master Test Level
* All unit tests pass
* No additional defects are known

## Test items
Please see the high level testing [Component_Test_Plan](https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/blob/master/Component_Test_Plan.md) 

## Features to be tested
```
features from users perspective
links back to requirements
```
## Features not to be tested
This section specifies the features and functionalities that are out of the scope for testing. It shall contain reasons of
why these features will not be tested

## Approach to testing
Testing will be performed in the testing framework “Testing” for the chosen language “Go”. The test plan will be written in a “Behaviour Driven Development” structured english style. 

The test plan priority is to ensure the application does not crash, due to unhandled runtime errors.
Unit tests will be to automate testing, and document the tests run. Integration tests will be used to test the methods in an application context. The tests will use data from a data structure to mock a database.

The priority of the testing will be to ensure that the a large number of cases are covered, using equivalence classes, to test invalid, boundary and valid.

The tests will produce a result or either, “Pass” or “Fail”. The tests that result as “Pass” will not be developed further, and are considered to meet the Hastings Driver Application initial requirements. The requirements that “Fail” will be marked as not behaving as expected. These failed requirements will be returned to development to ensure the test eventually passes.

## Test Deliverables
* Test plan document.
  * Test Features and items 
  * Sample test data
  * Findings report

## Development Team
### Training Needs
All testers need to have an understanding of the functional and non-functional requirements of the application. All testers need to have knowledge of Go and the testing package built into Go. Staff will need to be familar with the [to-do list software](http://abstractspoon.weebly.com/), and [GitHub](https://github.com/).

### Responsibilities
The roles are shared, and both Sarah and Mitch will be responsible for all components and deliverables. Sarah will oversee the quality of documentation and Mitch will oversee the quality and design on the software and tests 

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




