
//     
// [1,2,3,4]
// [1,0,0,0]

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	post := make([]int, n)
	res := make([]int, n)

	pre[0] = 1
	post[n-1] = 1

	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	for i := n-2; i >= 0; i-- {
		post[i] = post[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = pre[i] * post[i]
	}

	return res
}
