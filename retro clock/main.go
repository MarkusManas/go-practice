package main

func ComputeTime(n int) (int, int, int) {
	hours := n / 3600
	r := n % 3600
	minutes := r / 60
	seconds := r % 60

	return hours, minutes, seconds
}
func main() {
}
