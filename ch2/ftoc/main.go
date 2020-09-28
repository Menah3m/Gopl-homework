package main

import "fmt"

const freezingF, boilingF = 32.0, 212.0

func main() {
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF)) // 32 F = 0 C
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF)) // 212 F = 100C

}

func fToC(f float64) float64{
	return (f - 32) * 5 / 9
}