package main

func main() {
	arr := []string{"Python", "Go", "Java"}

	arr = append(arr[:1], arr[2:]...)

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
