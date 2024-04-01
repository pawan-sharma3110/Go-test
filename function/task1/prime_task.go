package main

import "fmt"

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var n int
    fmt.Println("Give a number greater than 2:")
    fmt.Scanln(&n)

    for i := 2; i <= n; i++ {
        if isPrime(i) {
            fmt.Println(i, "is prime.")
        }
    }
}

