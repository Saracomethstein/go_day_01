package handling

import (
	"fmt"
	"os"
)

func Error(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
}
