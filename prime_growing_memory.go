package main

func main() {
	const N = 20000000

	primes := make([]int, 0, N)

	for num := 2; num < N; num++ {
		isPrime := true
		for _, prime := range primes {
			if prime*prime > num {
				break
			}
			if num%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, num)
		}
	}

	println(len(primes))
}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.2 sec
//Memory before: 1,532 K
//Memory After: Ranges between 29,000 K and 45,000 k
