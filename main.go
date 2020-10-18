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

func generadorImpares() func() uint {
	i := uint(1) 
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func intercambia(a, b *int){
    temp := *b
    *b = *a
    *a = temp
}

func main() {
    var a int
    var b int
   	fmt.Scan(&a)
   	fmt.Scan(&b)
    fmt.Println(a, b)
    intercambia(&a, &b)
    fmt.Println(a, b)
}