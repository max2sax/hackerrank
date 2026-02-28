package bracketmatch

/*
 * Complete the 'areBracketsProperlyMatched' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts STRING code_snippet as parameter.
 */
type node struct {
	val  rune
	next *node
	prev *node
}

type stack struct {
	len  int
	head *node
	tail *node
}

func (s *stack) push_front(v rune) {
	s.len++
	if s.head == nil { // also means tail is nil
		s.head = &node{
			val: v,
		}
		s.tail = s.head
		return
	}
	oldHead := s.head
	newHead := &node{
		val:  v,
		next: oldHead,
	}
	oldHead.prev = newHead
	s.head = newHead
}

func (s *stack) pop_front() (r rune) {
	if s.len == 0 {
		return
	}
	s.len--
	front := s.head
	if front == nil {
		return
	}
	s.head = s.head.next
	r = front.val
	return
}

func areBracketsProperlyMatched(code_snippet string) bool {
	// Write your code here
	// create stack.
	// if opening bracket, push
	// if closing bracket, pop
	if len(code_snippet) < 2 {
		return false
	}
	brackets := &stack{}
	numOpening := 0
	numClosing := 0
	var lastOpen rune
	for _, r := range code_snippet {
		if r == '{' || r == '(' || r == '[' {
			brackets.push_front(lastOpen)
			lastOpen = r
			numOpening++
			continue
		}
		if r == '}' {
			if lastOpen != '{' {
				return false
			}
			lastOpen = brackets.pop_front()
			numClosing++
		}
		if r == ')' {
			if lastOpen != '(' {
				return false
			}
			lastOpen = brackets.pop_front()
			numClosing++
		}
		if r == ']' {
			if lastOpen != '[' {
				return false
			}
			lastOpen = brackets.pop_front()
			numClosing++
		}
	}
	if lastOpen != rune(0) {
		return false
	}
	return true
}
