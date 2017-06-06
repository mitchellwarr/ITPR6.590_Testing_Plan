# Test Report

### Test findings – Discoveries
One of our interpreted requirements was to have the the drivers have a change to exit of about 20%. This meant that the program would have a chance to showcase some of the features such as "Visiting john" and displaying varing messages. These features initally would not have been very visible as the drivers would have exited about 50% of the time. 

### Difficulties
Throughout the testing phase of the development, the code implementation was redesigned.  

In the code [issues] (https://github.com/mitchellwarr/ITPR6.590_Testing_Plan/issues?utf8=%E2%9C%93&q=) section of this repositiory, the difficulties were documented.  

The random generator posed an issue with test setup. It wasn't difficult to set up a new generator with a set seed, but knowing that a generator with a seed of 5 will return after running it through the method, meant running the method once with it and taking its result. This meant it was assumed the method was working in the first place in order to use the output as the base line. The only thing this would protect from is if either the workings of the method changed, or if the random generator code changed.  

Abstracting

### What could have been done better

### Observations
A lot of code has needed to be abstracted or operations exposed,
so that more control can be given to testing methods for setting up and verification.  

The visit message was split up into three different functions so that each aspect could be tested individually.
Also, the test messages were made to be global strings so that the testing methods could use them to build the
correct expected string.

### Imposed test limits
With the locations object, little was done to expand its testing. Creating an array of different JSON map files
just to check that the network loader was expansive seemed unnecessary. It was only tested to work with the
created Map JSON file and that if the file couldn't be found, that it won’t break the system.  

For the testing of the random arguments, extra care was not taken to the boundaries of equivalence classes
like 0 - 0.5 - 1, where the boundaries are 0-0.49999 0.5-0.999999. This is because it does not matter if the
random generator comes up with 0.24999999999999999999999999999999999999999999999 and it returns a 1 instead of
a 2; because that’s not the point of the class. Care was taken to test <0 and >=1 however.


```
Failing tests should be included here, along with why you think that they are failing.
```
We currently do not have any failing tests, this may be due to our background in coding. The tests were written by the programmers, and not by an analysist. Theses tests are purely whitebox, which is only one style of testing. We have not organised for blackbox testing to be completed. 
