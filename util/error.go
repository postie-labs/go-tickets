package util

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
