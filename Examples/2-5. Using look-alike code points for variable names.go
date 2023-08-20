package Examples

import "fmt"

func Example_02_05_Using_look_alike_code_points_for_variable_names() {
	ａ := "hello"   // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061)
	fmt.Println(ａ)
	fmt.Println(a)
}
