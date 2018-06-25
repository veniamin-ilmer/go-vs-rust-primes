package main

import "fmt"

func main() {

  var array_primes []int;

  for i := 2; i < 20000000; i++ {
    is_prime := true;
    for j := 0; j < len(array_primes) && array_primes[j] * array_primes[j] <= i; j++ {
      if i % array_primes[j] == 0 {
        is_prime = false;
        break;
      }
    }
    if is_prime {
      array_primes = append(array_primes, i);
    }
  }
  fmt.Println(len(array_primes));
  bufio.NewReader(os.Stdin).ReadBytes('\n')

}

//Compiled: go build -ldflags="-s -w"
//Run Time: 6.2 sec
//Memory before: 1,532 K
//Memory After: Ranges between 29,000 K and 45,000 k
