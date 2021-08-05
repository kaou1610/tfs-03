package main

import "fmt"

func Swap(a, b int) (int, int) {
	return b, a
}

func BubleSort(arr []int) ([]int){
	for i:=0 ; i< len(arr); i++ {
		for j:= len(arr)-1 ; j>0 ; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = Swap(arr[j], arr[j-1])
			}
		}
	}
	return arr
}

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2
	return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func Merge(left, right []int) []int {

	n, i, j := len(left)+len(right), 0, 0
	arr := make([]int, n, n)

	for k := 0; k < n; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}


func main()  {
	a := []int{2, 1, 3, 7, 5, 6, 4}
	b := MergeSort(a)
	fmt.Print(b)
	
}