package main

var (
	alphabetNum = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	result []string
)

func letterCombinations(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(alphabetNum[k]); i++ {
		res += alphabetNum[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}
}
