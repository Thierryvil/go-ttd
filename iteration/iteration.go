package iteration

import "fmt"

// Repete um caractere N vezes
func Repeat(char string, n int) string {
	var repeat string
	for i := 0; i < n; i++ {
		repeat += char
	}
	return repeat
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
