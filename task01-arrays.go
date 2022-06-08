package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0
	var counter float32 = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
		counter++
	}
	return sum / counter
}
