package moTest


func MapStrToStr (arr []string,fn func(s string) string) []string{
	var newArray = []string{}
	for _, it := range arr{
		newArray = append(newArray,fn(it))
	}
	return newArray
}