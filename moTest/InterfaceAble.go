package moTest

import "fmt"


type Country struct {
	Name string
}
type City struct {
	Name string
}
type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string {
	return "City = " + c.Name
}
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}