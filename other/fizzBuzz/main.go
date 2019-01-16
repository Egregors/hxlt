package main

import "log"

func isFizz(a int) bool {
	return a%3 == 0
}

func isBuzz(a int) bool {
	return a%5 == 0
}

func fizzBuzz(b, e int) {
	if b > e {
		panic("b > e")
	}

	for i := b; i <= e; i++ {

		if isFizz(i) && isBuzz(i) {
			log.Println("FizzBuzz")
			continue
		}

		if isFizz(i) {
			log.Println("Fizz")
			continue
		}

		if isBuzz(i) {
			log.Println("Buzz")
			continue
		}

		log.Println(i)
	}

}

func main() {
	fizzBuzz(11, 20)
}
