package main

import (
	"fmt"

	logger "github.com/reecerussell/advanced-logging"
)

func main() {
	f := "Hello %s!"
	v := "World"
	s := fmt.Sprintf(f, v)

	logger.Print(s)
}
