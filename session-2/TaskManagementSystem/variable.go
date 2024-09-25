package main

import "fmt"

const (
	AZE = "Azerbaijan"
	TUR = "Turkey"
	Pi  = 3.1415926
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Printf("Salam, %s. Bu il %dcu ildir.\n", "Orxan", 2024)
	fmt.Printf("1 kB is %f bytes\n", KB)
}
