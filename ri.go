package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano() | int64(os.Getpid()))

	countRelativelyPrime := 0

	for i := 0; i < count; i++ {
		n := rand.Int63()
		m := rand.Int63()
		g := gcd(n, m)
		if g == 1 {
			countRelativelyPrime++
		}
	}

	fmt.Printf("%.6f\n", float64(countRelativelyPrime)/float64(count))
}

// Euclid's algorithm for finding Greatest Common Divisor
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
