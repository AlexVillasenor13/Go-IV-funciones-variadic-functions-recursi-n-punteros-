package main

import "fmt"

func fibonacci(n uint) uint {
    if n == 0 {
		return 0
	}
    if n == 1 { 
		return 1
	}
    return fibonacci(n-1) + fibonacci(n-2)
}

func sum(args ...int) int { 
	result := 0
	for _, v := range args {
        if result < v{
            result = v
        }
	}

	return result
}

func main() {

}