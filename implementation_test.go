package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	tests := []struct {
		inputPrefix string
		expected    string
	}{
		{"+ 3 3", "3 3 +"},
		{"+ * 3 4 - 7 2", "3 4 * 7 2 - +"},
		{"+ 5 * 6 7", "5 6 7 * +"},
		{"* + 8 2 3", "8 2 + 3 *"},
		{"- * 5 2 + 6 3", "5 2 * 6 3 + -"},
		{"- + - + 5 / * 8 - 6 2 5 11 5 * 9 + 4 5", "5 8 6 2 - * 5 / + 11 - 5 + 9 4 5 + * -"},
		{"- - - + - * 5 5 7 8 * - 10 3 5 / 7 5 1", "5 5 * 7 - 8 + 10 3 - 5 * - 7 5 / - 1 -"},
		{"* - A / B C - / A K L", "A B C / - A K / L - *"},
	}

	for _, test := range tests {
		t.Run(test.inputPrefix, func(t *testing.T) {
			result, _ := PrefixToPostfix(test.inputPrefix)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPrefixToPostfix_EmptyLine(t *testing.T) {
	invalidPrefix := ""
	expectedErrorMessage := "empty line"

	_, err := PrefixToPostfix(invalidPrefix)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}

}

func TestPrefixToPostfix_InvalidToken(t *testing.T) {
	invalidPrefix := "= 3 5"
	expectedErrorMessage := "invalid token"

	_, err := PrefixToPostfix(invalidPrefix)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
