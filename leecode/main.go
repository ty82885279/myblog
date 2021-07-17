package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var maxSum, sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	if len(nums) == 1 {
		return float64(sum)
	}
	maxSum = sum
	for i := k; i < len(nums); i++ {

		sum -= nums[i-k]
		sum += nums[i]
		if maxSum < sum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func change(arr [3]int, sli []int) {
	fmt.Println("----2----")
	fmt.Printf("%p\n", sli)
	arr[0] += 1
	sli[0] += 1
	sli = append(sli, 10)
	fmt.Println("----3----")
	fmt.Printf("%p\n", sli)
	//fmt.Println("-------")
	//fmt.Println(sli)
	//fmt.Println("-------")
}
func main() {

	arr := [3]int{2, 4, 6}
	sli := []int{3, 5, 7}
	fmt.Println("----1----")
	fmt.Printf("%p\n", sli)
	change(arr, sli)
	fmt.Println(arr)
	fmt.Println(sli)

}
