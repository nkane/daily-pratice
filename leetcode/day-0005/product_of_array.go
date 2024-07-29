package leetcode

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	for i := 0; i < len(products); i++ {
		products[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		products[i] = products[i-1] * nums[i-1]
	}
	right := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		products[i] *= right
		right *= nums[i]
	}
	return products
}

func productExpectSelfSlower(nums []int) []int {
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))
	products := make([]int, len(nums))
	for i := 0; i < len(products); i++ {
		rightProducts[i] = 1
		leftProducts[i] = 1
		products[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		leftProducts[i] = nums[i-1] * leftProducts[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightProducts[i] = nums[i+1] * rightProducts[i+1]
	}
	for i := 0; i < len(products); i++ {
		products[i] = leftProducts[i] * rightProducts[i]
	}
	return products
}
