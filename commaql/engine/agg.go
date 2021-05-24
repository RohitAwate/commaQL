package engine

func SumInt(items []int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return
}

func SumDouble(items []float64) (sum float64) {
	for _, item := range items {
		sum += item
	}
	return
}

func AvgInt(items []int) float64 {
	return float64(SumInt(items)) / float64(len(items))
}

func AvgDouble(items []float64) float64 {
	return SumDouble(items) / float64(len(items))
}
