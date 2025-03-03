package cmd

import (
	"fmt"
	"os"

)


func debugMessage(m string) {
	if len(os.Getenv("DEBUG")) > 0 {
		fmt.Printf("%s", m)
	}
}
