package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"math"
)

func Readfile(s string) []int  {
	f, err := os.Open(s)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	var nums []int
    for scanner.Scan() {
		str, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, str)
    }
	return nums
    
}


func Max(nums []int) int {
	max := nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
	}
	return max
}

func Min(nums []int) int {
	min := nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}
	}
	return min
}

func checkExist(num int, nums[]int) bool {
	for _, value := range nums{
		if num == value {
			return true
		}
	}
	return false
}

func checkPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i:= 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main()  {
	arrNumbers := Readfile("lec02/file/numbers.txt")
	fmt.Println(arrNumbers)
	maxNumber := Max(arrNumbers)
	minNumber := Min(arrNumbers)
	fmt.Println("So co gia tri lon nhat trong file la: ", maxNumber)
	fmt.Println("So co gia tri nho nhat trong file la: ", minNumber)
	s := checkExist(3, arrNumbers)
	if s {
		fmt.Println("Exist" )
	} else {
		fmt.Println("Not Exist" )
	}
	v := checkPrime(arrNumbers[4])
	if v {
		fmt.Println("This is a prime number")
	} else {
		fmt.Println("This is not a prime number" )
	}
}