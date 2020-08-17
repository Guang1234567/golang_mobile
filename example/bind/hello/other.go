package hello

import (
	"fmt"
)

func Greetings2222(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type Counter222 struct {
	Value int
}

func (c *Counter222) Inc() { c.Value++ }

func NewCounter222() *Counter222 { return &Counter222{5} }
