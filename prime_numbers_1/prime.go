// Basic golang training, prime number locator (method 1)
// Andrew Alston

package main

import (
	"flag"
	"fmt"
	"time"
)

func FindPrime(Min, Max int) []int {
	// IsPrime is a boolean we will use to track if a number is or isn't prime on our inner loop
	var IsPrime bool
	// res is our slice of integers to store the prime numbers in for return,
	var res []int

	// Our outer loop
	for i := Min; i < Max; i++ {
		// Assume that i is a prime number at the start of each iteration
		IsPrime = true
		// Since anything divided by a number greater than half of itself is always going to produce
		// a non-zero remainder - we can run the loop to half of our input number
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				// The moment we find a clean division, we can break out of the inner loop,
				// after setting the variable to tell us that this isn't a prime number
				IsPrime = false
				break
			}
		}
		// If this isn't a prime number, append it to our result list
		if IsPrime {
			res = append(res, i)
		}
	}
	// Finally return the result list
	return res
}

func main() {
	var StartTime time.Time

	// These are our CLI flags - they can be specified in any order on the command line.
	// The first variable specified in the cli flag itself, the second is the default for the flag, the third
	// is the help for the flag itself and is what is shown when you print the defaults
	Minimum := flag.Int("min", 2, "Minimum number in range to search for primes")
	Maximum := flag.Int("max", 4000, "Maximum number in range to search for primes")
	DumpPrimes := flag.Bool("dump-prime", false, "Dump the list of prime numbers located")
	TimeExecution := flag.Bool("timing", false, "Dump the execution and prime the results")
	flag.Parse()

	// Note: flags are always pointers, so we have to de-reference them, hence the asterix
	if *Minimum > *Maximum || *Minimum == *Maximum {
		fmt.Printf("Minimum in range must be smaller than range maximum")
		flag.PrintDefaults()
		return
	}
	if *Minimum < 2 {
		fmt.Printf("Please specify a minimum number greater than or equal to two\n")
		flag.PrintDefaults()
		return
	}

	// If our timing flag is set - grab the start time before we start looking for the prime numbers
	if *TimeExecution {
		StartTime = time.Now()
	}

	// Find the prime numbers calling our FindPrime function with our minimum and maximum arguments
	results := FindPrime(*Minimum, *Maximum)

	// If our timing flag is set - prime the time its taken to find all our prime numbers
	if *TimeExecution {
		fmt.Printf("Took us %s to find all primes in a range of %d numbers\n", time.Since(StartTime), *Maximum-*Minimum)
	}
	fmt.Printf("Found %d prime numbers between %d and %d\n", len(results), *Minimum, *Maximum)

	// If our dump flag is set - dump the list of located prime numbers
	if *DumpPrimes {
		fmt.Printf("Located the following prime numbers in the range %d -> %d\n", *Minimum, *Maximum)
		for i, p := range results {
			// \t is an escape code for a tab
			// \n is a new line character
			// %d tells us that the variable or constant being referenced is an integer
			// If we wanted to print a float, that would be a %f - more about number formatting later
			fmt.Printf("\t[%d]: %d\n", i, p)
		}
	}
}
