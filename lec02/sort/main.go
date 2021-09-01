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

	i, j :=  0, 0
	arr := []int{}
	for i < len(left) && j <len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		arr = append(arr, left[i])
	}
	for ; j < len(right); j++ {
		arr = append(arr, right[j])
	}
	return arr
}


func main()  {
	a := []int{2, 1, 3, 7, 5, 6, 4}
	b := MergeSort(a)
	fmt.Print(b)
	
}