fn main() {

  let mut array_primes: Vec<u64> = vec![0;20000000];
  let mut slice_len = 0;

  for i in 2..20000000 {
    if array_primes.iter().take_while(|&p| p * p <= i && *p != 0).all(|&p| i % p != 0) {
      array_primes[slice_len] = i; //Set value
      slice_len += 1;
    }
  }
  println!("{}", slice_len);

}

//Compiled: cargo build --release
//5.1 sec
//Memory Before: 576 K
//Memory After: 10,516 K
