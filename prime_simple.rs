fn main() {
  let count = (2..10000000).filter(|&p| is_prime(p)).count();
  println!("{}", count);
}

fn is_prime(n: u64) -> bool {
  (2..).take_while(|x| x * x <= n).all(|i| n % i != 0)
}

//Compiled: cargo build --release
//Run Time: 11.9 sec
//Memory: 548 K
