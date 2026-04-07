// ["1", "2", "+", "3", "*", "4","-"]
//                            i
// s = [9,4]
// n = 

func evalRPN(tokens []string) int {
	stack := make([]int,0)

	for _, t := range tokens {
		if n, err := strconv.Atoi(t); err == nil {
			stack = append(stack, n)
		}   else {
			switch t {
				case "+":
					a := stack[len(stack)-1]
					b := stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					stack = append(stack, a+b)
				case "*":
					a := stack[len(stack)-1] 
					b := stack[len(stack)-2] 
					stack = stack[:len(stack)-2] 
					stack = append(stack, a*b)
				case "/":
					a := stack[len(stack)-1] 
					b := stack[len(stack)-2] 
					stack = stack[:len(stack)-2] 
					stack = append(stack, b/a)
				case "-":
					a := stack[len(stack)-1] 
					b := stack[len(stack)-2] 
					stack = stack[:len(stack)-2] 
					stack = append(stack, b-a)

			}
		}
	}
	return stack[0]
}
