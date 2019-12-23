package ocheredi

var x []string

func Push(str string) {
	x = append(x, str)
}
func Pop() string {
    if len(x) == 0 {
        return ""
    }
    elem := x[0]
    x = x[1:]
    return elem
}
