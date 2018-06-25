package main

import "fmt"

func main() {

  var array_primes [20000000] int;
  var slice_primes []int = array_primes[:0];

  for i := 2; i < 20000000; i++ {
    is_prime := true;
    for j := 0; j < len(slice_primes) && slice_primes[j] * slice_primes[j] <= i; j++ {
      if i % slice_primes[j] == 0 {
        is_prime = false;
        break;
      }
    }
    if is_prime {
      slice_length := len(slice_primes);
      slice_primes = slice_primes[:slice_length + 1];  //Extend slice by 1.
      slice_primes[slice_length] = i; //Set value
    }
  }
  fmt.Println(len(slice_primes));

}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.0 sec
//Memory before: 7,000 K
//Memory After: Ranges between 16,900 K
