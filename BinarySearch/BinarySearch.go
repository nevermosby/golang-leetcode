package main

import(
	"fmt"
)
func main()  {
	
	ex001 := []int{-1,0,1,3,5,9,12,15,20,21}
	target := 12

	result := search(ex001, target)
	fmt.Println("result: ", result)
}
// leetcode #704
func search(nums []int, target int) int {

    low := 0
    high := len(nums) -1
    mid := -1
    guess := -1
    for low <= high {
		mid = (low+high) / 2
		// need to handle when mid is not a intger
		// however golang will take care of it, it will use round() by default
		fmt.Println("mid is: ",mid)
        guess = nums[mid]

        if guess == target {
            return mid
        }
        if guess > target {
            high = mid -1
        } else {
            low = mid + 1
        }
    }
    return -1
}

