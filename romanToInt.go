package main

func romanToInt(s string) int {
	matchingMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	prevNum := 0
	result := 0
	for _, char := range s {
		if val, ok := matchingMap[string(char)]; ok {
			if result == 0 {
				result = val
			} else if prevNum >= val {
				result += val
			} else {
				result += (val - prevNum) - prevNum
			}
			prevNum = val
		}
	}
	return result
}

func test(s string) int {
	matchingMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	prevNum := 0
	result := 0
	for i := 0; i < len(s); i += 2 {
		next := matchingMap[string(s[i+1])]
		if val, ok := matchingMap[string(s[i])]; ok {
			if result == 0 {
				result = val
			} else if prevNum >= val && val > next{
				result += val
			} else if prevNum < val && val > next{
				result += val - prevNum
			} else if prevNum > val && val < next{
				result += next - val
			}
			prevNum = val
		}
	}

	return result
}
