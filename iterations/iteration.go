package iterations


func Repeat(character string, repeatCount int)  (output string) {
	for i := 0; i < repeatCount; i++ {
		output = output + character
	}
	return 
}