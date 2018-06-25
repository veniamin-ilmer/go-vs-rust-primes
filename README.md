# go-vs-rust-primes
Counting up how many primes exist within the first 20,000,000 numbers.
Comparing Go and Rust executions.

## Introduction
About a year ago, I learned Go.
I liked how you can write simple "python" like code, but compile into fast code.

As dove more into Go however, I started to doubt some of the efficiency.
I tried building concurrent processes with a shared variable, and realized Go failed to stop race conditions of updating the variable at the same time.

Up till now, I avoided Rust.. Rust seemed more difficult to program than Go.
It required more keywords, like mutability, and had Higher Order Functions, which I didn't understand the point of.
Why have all these extra functions (map, iter, filter, etc), when all of this can be done with a for loop?

But with Rust being rated the [most loved language for the 3rd time in a row](https://insights.stackoverflow.com/survey/2018/#most-loved-dreaded-and-wanted), I decided to try it out.

Although there was a learning curve, Rust is becoming my favorite language.
All of the things I noted earlier as reasons I didn't want to use Rust (mutability, Higher Order Functions), are now reasons why I like it. I've learned to see these qualities as benefits of the language.

I now understand, Go's simplified language, makes Go less legible when writing advanced code.
Here's an applicable analogy: English is a relatively complicated language and is difficult to learn for foreigners.
But using English, you talk about very complicated abstract philosophy.
On the other hand, a language like [Toki Pona](https://en.wikipedia.org/wiki/Toki_Pona) is very easy to learn. It only has about 120 words.
But using such a simple language limits your ability to effectively communicate advanced, abstract concepts.

Similarly in a simple language like Go, once you start writing advanced code, it starts looking "messy" and difficult to read.
Rust's use of Higher Order Functions, makes advanced code look less messy.

## The Test

Now that I've dabbled with Rust a little bit, I decided to run a test, comparing the speed and memory usage of Go vs Rust, compiled with maximum optimization.

Here's the task:
Count up the numbers between 0 to 20,000,000 of prime numbers 
