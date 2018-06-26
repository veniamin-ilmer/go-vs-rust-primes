package main

func main() {
	const N = 20000000

	var primes [N]int
	var head int

	for num := 2; num < N; num++ {
		isPrime := true
		for _, prime := range primes[:head] {
			if prime*prime > num {
				break
			}
			if num%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes[head] = num
			head++
		}
	}

	println(head)
}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.0 sec
//Memory before: 7,000 K
//Memory After: 16,900 K
