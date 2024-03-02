package task

import (
	"sort"
	"sync"
)

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GeneratePrimes generates prime numbers up to a given limit.
func GeneratePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func GeneratePrimesSieve(limit int) []int {
	// Validation.
	if limit < 1 {
		return []int{}
	}

	// Create sieve.
	sieve := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		sieve[i] = true
	}
	p := 2
	for p*p <= limit {
		if sieve[p] {
			for i := p * p; i <= limit; i += p {
				sieve[i] = false
			}
		}
		p += 1
	}

	// Check prime number with buffered channel and goroutine.
	var wg sync.WaitGroup
	result := make(chan int, limit)
	for i := 0; i <= limit; i++ {
		wg.Add(1)
		go func() {
            defer wg.Done()
            if sieve[i] {
                result <- i
            }
        }()
	}

	// Close the result channel after all goroutines have finished.
    wg.Wait()
    close(result)
    
    // Collect prime number and sort.
	primes := make([]int, len(result))
    j := 0
	for prime := range result {
        primes[j] = prime
        j++
	}
    sort.Ints(primes)

	return primes
}

