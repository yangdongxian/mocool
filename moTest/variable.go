package moTest

import "fmt"

var name string
var age int
var address string

const Pi = 3.1415

func TraceString() {
	var s string = "eastsheen.东东"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v,%c ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v,%c ", r, r)
	}

	fmt.Println()
	var arr1 []string = []string{"aa", "bb", "cc"}
	var arr2 []string = []string{"dd", "ee", "ff", "cc"}

	slices := append(arr1, arr2...)

	fmt.Println("slices:", slices)
}
