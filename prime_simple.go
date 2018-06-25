package main

import "fmt"

func main() {
  count := 0;
  for i := 2; i < 20000000; i++ {
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
//Run Time: 34.5 sec
//Memory: 1,490 K
