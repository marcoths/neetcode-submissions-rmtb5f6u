//      i
// [30,38,30,36,35,40,28]
//   0, 1, 2, 3, 4, 5, 6  
// 
// [0,]
// []
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int,0)
	res := make([]int,len(temperatures))


	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			res[topIdx] = i - topIdx
		}
		stack = append(stack, i)
	}
	return res
}
