package trap

/*
	方法：单调栈。栈内元素保持单调递减。当前元素比栈顶元素大则出栈并计算高度和宽度
*/

func trap(height []int) int {
	n := len(height)
	ans := 0
	if n < 3 {
		return ans
	}

	stack := NewStack(n)

	for i, curHeight := range height {
		// 当前值大于栈顶元素，则循环出栈
		for stack.Len() > 0 && curHeight > height[stack.Top()] {
			popedIndex := stack.Pop() // 取出元素索引
			if stack.Len() == 0 {
				break
			}

			// 宽度
			w := i - stack.Top() - 1 // 宽度
			// 高度, 左右边界的最小值
			h := min(height[stack.Top()], curHeight) - height[popedIndex]
			// 容量
			ans += w * h
		}

		stack.Push(i)
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type Stack struct {
	data []int
}

func NewStack(n int) *Stack {
	return &Stack{
		data: make([]int, 0, n),
	}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Top() int {
	return s.data[s.Len()-1]
}

func (s *Stack) Pop() int {
	n := s.Len()
	d := s.data[n-1]
	s.data = s.data[:n-1]
	return d
}

func (s *Stack) Len() int {
	return len(s.data)
}
