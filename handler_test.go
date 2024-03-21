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

// Include more tests for error handling and other cases.
