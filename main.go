package main

import (
	"fmt"

	"sandbox3.0/task"
)

func main() {
	limit := 20
	primes := task.GeneratePrimes(limit)
	fmt.Println(primes)

	primesSieve := task.GeneratePrimesSieve(limit)
	fmt.Println(primesSieve)
}
