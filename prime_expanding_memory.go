package main

import "fmt"

func main() {

  var slice_primes []int;

  for i := 2; i < 20000000; i++ {
    is_prime := true;
    for j := 0; j < len(slice_primes) && slice_primes[j] * slice_primes[j] <= i; j++ {
      if i % slice_primes[j] == 0 {
        is_prime = false;
        break;
      }
    }
    if is_prime {
      slice_primes = append(slice_primes, i);
    }
  }
  fmt.Println(len(slice_primes));
  bufio.NewReader(os.Stdin).ReadBytes('\n')

}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.2 sec
//Memory before: 1,532 K
//Memory After: Ranges between 29,000 K and 45,000 k
