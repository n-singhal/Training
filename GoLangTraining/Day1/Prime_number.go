package main

import "fmt"

func main() {
	CheckPrime(19) //Prime number
	CheckPrime(35) //Not a prime number
}

func CheckPrime(num int) {
	isPrime := true
	if num == 0 || num == 1 {
		fmt.Printf("%d is not a prime number\n", num)
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				fmt.Printf("%d is not a prime number\n", num)
				isPrime = false
				break
			}
		}

		if isPrime == true {
			fmt.Printf("%d is a prime number\n", num)
		}
	}
}
