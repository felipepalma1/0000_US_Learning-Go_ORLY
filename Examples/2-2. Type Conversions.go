package Examples

import "fmt"

type Result struct {
	X int
	Y float64
	Z float64
	D int
}

func Example_02_02_Type_Conversions() Result {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)

	result := Result{
		X: x, Y: y, Z: z, D: d}

	fmt.Println(result)
	return result
}
