package main

func main() {
	const N = 20000000

	count := 0
	for i := 2; i < N; i++ {
		if isPrime(i) {
			count += 1
		}
	}

	println(count)
}

func isPrime(value int) bool {
	for i := 2; i*i <= value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

//Compiled: go build -ldflags="-s -w"
//Run Time: 34.5 sec
//Memory: 1,490 K
