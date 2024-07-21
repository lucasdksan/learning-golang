package main

func sun_arr(arr []int) int {
	var result int = 0

	for i := 0; i < len(arr); i++ {
		result = result + arr[i]
	}

	return result
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}

	arr[2] = -2

	result := sun_arr(arr)

	println(result)
}
