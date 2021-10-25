package main

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	incomingNumber := x
	num := 0
	for incomingNumber > 0 {
		num *= 10
		num += incomingNumber % 10
		incomingNumber /= 10
	}
	if x ==  num{
		return true
	}
	return false
}
