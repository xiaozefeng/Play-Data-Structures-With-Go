package stack

func isValid(s string) bool {
	st := &stack{
		data: make([]string, 10),
		size: 0,
	}
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == "(" || c == "[" || c == "{" {
			st.push(c)
		} else {
			if st.isEmpty() {
				return false
			}
			top := st.pop()
			if c == ")" && top != "(" {
				return false
			}
			if c == "]" && top != "[" {
				return false
			}
			if c == "}" && top != "{" {
				return false
			}
		}
	}
	return st.isEmpty()
}

type stack struct {
	data []string
	size int
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) push(e string) {
	s.data = append(s.data, e)
	s.size++
}

func (s *stack) pop() string {
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.size--
	return ret
}
