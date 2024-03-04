## Benchmark

Benchmark GeneratePrimes function 10 times.

```go test -benchmem -run=^$ -bench ^BenchmarkGeneratePrimes$ sandbox3.0/test -count=10 > old.txt```

Benchmark GeneratePrimes with Sieve of Erastothenes function 10 times.

```go test -benchmem -run=^$ -bench ^BenchmarkGeneratePrimes$ sandbox3.0/test -count=10 > new.txt```

Compare the benchmark results.

```benchstat old.txt new txt```

The result show GeneratePrimes comparation

                        │   old.txt     │                new.txt                        │
                        │   sec/op      │       sec/op             vs base              │
    GeneratePrimes-8    |  44.27µ ± 4%  |   281.64µ ± 7%     +536.25% (p=0.000 n=10)    |


                        │   old.txt     │                new.txt                        │
                        │   B/op        │       B/op             vs base                │
    GeneratePrimes-8    |  3.992Ki ± 0% |   72.953Ki ± 0%     +1727.40% (p=0.000 n=10)  |

                        │   old.txt     │                new.txt                        │
                        │   allocs/op   │       allocs/op             vs base           │
    GeneratePrimes-8    |  9.000 ± 0%   |   1005.000 ± 0%     +11066.67% (p=0.000 n=10) |

