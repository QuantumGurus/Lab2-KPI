package lab2

import (
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	inputData, err := io.ReadAll(h.Input)
	if err != nil {
		return err
	}

	result, err := PrefixToPostfix(string(inputData))
	if err != nil {
		return err
	}

	_, err = h.Output.Write([]byte(result + "\n"))
	return err
}
