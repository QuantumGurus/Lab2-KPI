package lab2

import (
	"errors"
	"strings"
	"unicode"
)

// Stack represents a stack data structure.
type Stack struct {
	items []string
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// IsOperator checks if a given token is an operator.
func IsOperator(token string) (bool, error) {
	if strings.ContainsAny(token, "+-*/^") {
		if len(token) != 1 {
			return false, errors.New("please separate all characters with spaces")
		}
	}

	return strings.ContainsAny(token, "+-*/^"), nil
}

func IsOperand(token string) bool {
	for _, char := range token {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// PrefixToPostfix converts a prefix expression to postfix.
func PrefixToPostfix(prefix string) (string, error) {
	tokens := strings.Fields(prefix)
	stack := Stack{}
	if len(tokens) == 0 {
		return "", errors.New("empty line")
	}
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		isOperator, err := IsOperator(token)
		if err != nil {
			return "", err
		}
		if !isOperator {
			if IsOperand(token) {
				stack.Push(token)
			} else {
				return "", errors.New("invalid token")
			}
		} else {
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			postfix := operand1 + " " + operand2 + " " + token
			stack.Push(postfix)
		}
	}

	return stack.Pop(), nil
}
