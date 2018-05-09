package main

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}
	if n <= 26 {
		return string('A' + n - 1)
	}

	if n % 26 != 0 {
		return convertToTitle(n / 26) + convertToTitle(n % 26)
	} else {
		return convertToTitle(n / 26 - 1) + convertToTitle(26)
	}
}

func main() {

}