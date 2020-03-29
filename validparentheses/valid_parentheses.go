package validparentheses

// StackNode ...
type StackNode struct {
	Bracket rune
	Next    *StackNode
}

// MyStack ...
type MyStack struct {
	Top *StackNode
}

// NewStackNode ...
func NewStackNode(bracket rune) *StackNode {
	return &StackNode{
		Bracket: bracket,
		Next:    nil,
	}
}

// Pop ...
func (stack *MyStack) Pop() rune {

	bracket := stack.Top.Bracket
	stack.Top = stack.Top.Next

	return bracket
}

// Push ...
func (stack *MyStack) Push(bracket rune) {
	stackNode := NewStackNode(bracket)
	stackNode.Next = stack.Top
	stack.Top = stackNode
}

// Peek ...
func (stack *MyStack) Peek() rune {
	return stack.Top.Bracket
}

// IsEmpty ...
func (stack *MyStack) IsEmpty() bool {
	return stack.Top == nil
}

// IsValid ...
func IsValid(s string) bool {

	var openCloseBrackets map[rune]rune = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var answer bool = true

	bracketsStack := &MyStack{Top: nil}

	for _, bracket := range s {

		if _, ok := openCloseBrackets[bracket]; ok {
			bracketsStack.Push(bracket)
		} else {

			if bracketsStack.IsEmpty() || bracket != openCloseBrackets[bracketsStack.Peek()] {
				answer = false
				break
			}
			bracketsStack.Pop()
		}
	}

	if !bracketsStack.IsEmpty() {
		answer = false
	}
	return answer
}
