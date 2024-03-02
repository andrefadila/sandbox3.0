package task

import (
	// "fmt"
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
	// fmt.Println(sieve)
    
	primes := make([]int, 0, limit)
	for p = 2; p <= limit; p++ {
		if sieve[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func CheckPrimes(nums []int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		if isPrime(num) {
			results <- num
		}
	}
}
