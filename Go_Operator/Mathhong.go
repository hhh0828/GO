package main

//나누기
func Divide(a float64, b float64) (float64, bool) {

	if b == 0 {
		return 0, false
	}
	return a / b, true

}

//반환 이름 지정
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return 0, success
	}
	result = a / b
	success = true
	return result, success

}
