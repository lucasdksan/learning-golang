package main

func par_slice(slice []int) []int {
	var new_par_slice []int

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			new_par_slice = append(new_par_slice, slice[i])
		}
	}

	return new_par_slice
}

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	result := par_slice(slice)

	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
