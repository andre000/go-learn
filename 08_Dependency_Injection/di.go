package dependency

import (
	"fmt"
	"io"
)

// Greet returns "Hello, {name}" to a Writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
