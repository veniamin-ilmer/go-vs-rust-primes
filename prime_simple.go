package main

import "fmt"

func main() {
  count := 0;
  for i := 2; i < 10000000; i++ {
    if is_prime(i) {
      count += 1;
    }
  }
  fmt.Println(count);
}

func is_prime(value int) bool {
    for i := 2; i * i <= value; i++ {
        if value % i == 0 {
            return false
        }
    }
    return true
}

//Compiled: go build -ldflags="-s -w"
//Runs in 12.9 seconds
//Memory: 1488 K
