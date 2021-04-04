package main
import "fmt"

func main() {
	go animation()
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func animation() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
