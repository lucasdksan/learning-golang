package main

func main() {
	arr := []string{"Python", "Go", "Java"}

	arr = append(arr, "C++")
	arr = append(arr, "Ruby")

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
