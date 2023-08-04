package main

import (
	"fmt"

	"github.com/Sona-28/gopackagedemo/calc"
)

func main() {
	sum := calc.Add(10, 20)
	fmt.Println(sum)
}
