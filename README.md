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

## The Task
Now that I've dabbled with Rust a little bit, I decided to run a test, comparing the speed and memory usage of Go vs Rust, compiled with maximum optimization.

Here's the task:
Count up all the numbers to 20,000,000 that are prime.

As an example, if you look at numbers up to 10, you find that these are prime:

2, 3, 5, 7

So that means there are 4 prime numbers in the range up to 10.

How many prime numbers are there up to 20,000,000?

There are several ways to do this calculation. I've built these ways in Rust and Go.

To make things fair, I made sure to use exactly the same algorithm in both languages.

## Attempt 1: Simply calculate if each number is prime.
The way to tell if a number is prime, is by trying to divide it by different numbers. If it's not divisible by any numbers, it's prime.

An efficient way to do this, is to iterate from 2 up to the square root of the number.

You start from 2, because we already know all numbers are divisible by 1.

You end at the square of the number, because if there were no divisors up to that point, there can't be further divisors. (The only way that 6 is divisible into 12 is because a lower number - 2 is divisible by 12.)

Another way to say "loop until you reach the square root of the number" is "loop until your iterator ^ 2 equals the number". I use this due to computers preferring working with integers over decimals.

Finally, for efficiency, break out early if you find any divisors.

This algorithm was written in `prime_simple.go` and `prime_simple.rs` for Go and Rust.

### prime_simple.go
go build -ldflags="-s -w" prime.go

Compiled into 1,664 KiB.

It ran in 34.5 Seconds.

The total memory usage was 1,490 KiB.

### prime_simple.rs
cargo build --release

Compiled into 152 KiB.

It ran in 31.9 Seconds.

The total memory usage was 550 KiB.

### Analysis
Although I stripped as much debugging information as Go allows, there is still some residual data, explaining the huge difference in size between the two compiled versions.

For the amount of work done by each execution, the amount of time run is relatively similar. Rust was 2.6 seconds faster than Go.

The Go code took up 2.7x more space in memory than the Rust code. Probably because of the extra garbage collection code, and other runtime checks that Go requires. (Rust would have all the garbage collection and other handling done at compile time.)

## Attempt 2: Remembering previous primes
A further optimization: Instead of iterating through all integers to up the square root of a number, only iterate through all prime numbers up to the square root.

For example, to check if 13 is prime:

* You first check if it divides by 2
* Then you check if it divides by 3
* Then you don't actually need to check if it divides by 4, because if it did, it would have divided by 2 earlier. Skip to the next prime: 5.

If you can remember which numbers you found were prime earlier, you can use those numbers to check for future primes.

To be memory efficient, since we don't know exactly how many numbers we need to remember, we can just create a growing array.

Both Rust and Go have a way to build a growing array.

In Go, any time you want to add a value to an array, you can use the function "append".

In Rust, any time you want to add a value to an array, (a growable array - a Vector), you can use the function "push".

This algorithm was written in `prime_growing_memory.go` and `prime_growing_memory.rs` for Go and Rust.

### prime_growing_memory.go
go build -ldflags="-s -w" prime.go

Compiled into 1,664 KiB.

It ran in 6.2 seconds.

The starting memory usage was 1,532 KiB.
After rerunning the binary multiple times, I found that the ending memory usage kept ranging between 29,000 KiB and 45,000 KiB.

### prime_growing_memory.rs
cargo build --release

Compiled into 153 KiB.

It ran in 6.0 Seconds.

The starting memory usage was 564 KiB.
The ending memory usage was stable at 10,512 KiB.

### Analysis
Even though the algorithm was significantly different, the compiled size stayed relatively similar for both executables.

The run time was signficantly better, and Rust only had a 0.2 second lead.

Memory usage is where Go and Rust were significantly different.

Note that a total of 1,270,607 primes were recorded. With a 64 bit integer, that's 9,927 KiB worth of data.

Adding 9,927 to the initial 564 = 10,490 KiB. We see that Rust's increase in memory is extremely close to how much memory was actually required.

On the other hand, Go used 3 to 4.5 times as much data.

Understandibly, this was probably because Go's garbage collector made the determination not to clean up memory with old arrays.

## Attempt 3: Slices
Go pushes the concept: Don't use Arrays. Everything should be a slice.

So, it might be more fair to build things "the Go way", and see if Go then gets an advantage over Rust.

Here's what that means: Instead of building an array, then appending values on top of it, rebuilding the array with a new size over and over, we will have one large array prebuilt, and just keep updating a reference pointer to the array.

A Slice is a reference pointer to different points in an array, that looks and feels like an array. You can keep on "expanding it" even though in memory, the array keeps being in the same place without changing.

Although Go uses this concept extensively, Rust does not. Although Rust contains the syntax required to build a Slice, you cannot easily increase the size of the Slice, without running into mutibility borrowing problems. It's a bug in Rust that's being tracked here - https://github.com/rust-lang/rust/issues/43234

Since I was not able to use Slices in Rust, I had to manually keep track of reference pointers to the array.

Basically, I had a variable tracking the "slice length", and kept iterating through the array up to the "slice length", and kept increasing the "slice length" every time there was a new prime number.

This reflects how Go would handle the slice "under the hood".

In both cases, I created an integer array space of 20,000,000. This is because, without running the application, I don't know exactly how many primes there would be.

20,000,000 is simply the maximum amount of numbers that could ever be prime in my run, so that's how much space I made in advance.

This algorithm was written in `prime_slice_memory.go` and `prime_slice_memory.rs` for Go and Rust.

### prime_slice_memory.go
go build -ldflags="-s -w" prime.go

Compiled into 1,664 KiB.

It ran in 6.0 seconds.

The starting memory usage was about 7,000 KiB.
The ending memory usage was stable at about 16,900 KiB.

### prime_slice_memory.rs
go build -ldflags="-s -w" prime.go

Compiled into 153 KiB.

It ran in 5.1 seconds.

The starting memory usage was 576 KiB.
The ending memory usage was stable at 10,516 KiB.

### Analysis
Compiled size remains the same as before.

Although keeping one array around all the time instead of rebuilding it made Go 0.2 seconds faster, Rust significantly benefited from the optimization becoming 0.9 seconds faster.

The memory use was most surprising. A 64 bit 20,000,000 integer array should take up **156,250 KiB** of memory.

Even though we declared the array to have this much space early on, both Rust and Go did not allocate that much memory. Both of them seemed to wait and only use memory when it was necessary.

Incredibly, for Rust, the memory usage was remarkably similar to what it was back when we simply pushed values into the array.

Somehow, Rust has been very good at predicting almost exactly how much memory will be used each time.

Although Go seemed to also have some kind of memory optimization, it wasn't nearly as close as Rust.

## Conclusion
Although Go and Rust are relatively similar in speed (with Rust usually being a little better), Rust seems to be significantly better at memory management.

The fact that it can predict memory use at compile time, and not require background garbage collection, is an extremely beneficial feature.

That, along with the increased legibility that comes with using Rust Higher Order Functions, makes this a very useful and powerful language.
