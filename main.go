package main

func main() {
	getRegards()
	getOtherthing(2)
	getSomething(3)
}

func getRegards() string {
	return "Hello little man"
}

func getSomething(n int) string {
	if n > 2 {
		return "n > 2"
	}

	return "n <= 2"
}

func getOtherthing(n int) string {
	if n == 3 {
		return "amazing"
	}

	if n == 4 {
		return "kkk"
	}

	return "oh :("
}
