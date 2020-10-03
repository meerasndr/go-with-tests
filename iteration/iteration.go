package iteration

func Repeat(x string, times int) string{
	//repeated := ""
	var repeated string
	for i := 0; i < times ; i++ {
		repeated += x
	}
	return repeated
}