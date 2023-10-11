package main

import (
	"fmt"
)

func changeNumber(num *int) {
	*num = 12
}

func main() {
	nomor := 10
	fmt.Printf("Nomor %v \n", nomor)

	changeNumber(&nomor)

	fmt.Printf("Nomor %v", nomor)
}
