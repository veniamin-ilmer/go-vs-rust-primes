package main

import "fmt"

func main() {

  var slice_primes = make([]int, 20000000);
  var slice_length = 0;

  for i := 2; i < 20000000; i++ {
    is_prime := true;
    for j := 0; j < slice_length && slice_primes[j] * slice_primes[j] <= i; j++ {
      if i % slice_primes[j] == 0 {
        is_prime = false;
        break;
      }
    }
    if is_prime {
      slice_primes[slice_length] = i;
      slice_length += 1;
    }
  }
  fmt.Println(slice_length);

}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.2 sec
//Memory before: 1,532 K
//Memory After: 16,900 K
