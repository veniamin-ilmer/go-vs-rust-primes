fn main() {

  let mut vec_primes: Vec<u64> = Vec::with_capacity(20000000);
  for i in 2..20000000 {
    if vec_primes.iter().take_while(|&p| p * p <= i).all(|&p| i % p != 0) {
      vec_primes.push(i);
    }
  }
  println!("{}", vec_primes.len());

}

//Compiled: cargo build --release
//6.0 sec
//Memory Before: 564 K
//Memory After: 10,512 K
