package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_Success(t *testing.T) {
	input := "+ 3 3"
	expected := "3 3 +\n"

	var output bytes.Buffer
	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: &output,
	}

	err := handler.Compute()
	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}

func TestComputeHandler_InvalidInput(t *testing.T) {
	// Input containing invalid characters for the PrefixToPostfix conversion
	input := "+ = 3"
	var output bytes.Buffer
	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: &output,
	}

	err := handler.Compute()

	// Assuming PrefixToPostfix returns an error for invalid characters
	assert.Error(t, err, "expected an error for invalid input characters")
}

// Include more tests for error handling and other cases.
