package ocheredi

var x []string

func Push(str string) {
	x = append(x, str)
}
func UnPush() string {
	var sub []string

	elnum := len(x)

	if elnum == 0 {
		return ""
	}
	first := x[0]
	if elnum == 1 {

		x[0] = ""
		return first
	}
	for i := 1; i < elnum; i++ {
		sub = append(sub, x[i])
	}
	x = sub
	return first
}
