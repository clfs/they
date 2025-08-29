package uci

import "fmt"

func ExampleParse() {
	b := []byte("uci")

	cmd, err := Parse(b)
	if err != nil {
		// handle
	}

	switch cmd.(type) {
	case *UCI:
		fmt.Println("Received a uci command!")
	case *IsReady:
		fmt.Println("Received an isready command!")
	}
	// Output:
	// Received a uci command!
}
