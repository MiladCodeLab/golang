package main

func print10times(s string) {
	for val := range 10 {
		print(val, " ", s, "\n")
	}
}

func main() {

	go print10times("Goroutines")

}
