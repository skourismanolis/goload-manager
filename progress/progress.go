package progress

func GetBar(progress float64, size int) string {
	if progress > 1 {
		progress = 1
	}
	ret := "["
	current := int(float64(size) * progress)

	for i := 0; i < current; i++ {
		ret += "="
	}
	if progress < 1 {
		ret += ">"
	}

	for i := 0; i < size-(current+1); i++ {
		ret += "-"
	}
	ret += "]"
	return ret
}
