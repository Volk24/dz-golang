package main

import "fmt"

func main() {

	const UsdEur float64 = 0.94691
	const UsdRub float64 = 100.0
	const EurRub float64 = UsdEur * UsdRub

	fmt.Print(EurRub)
}
