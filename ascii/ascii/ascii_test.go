package ascii

import "testing"
import "fmt"

// TestGreetingsASCII is a test
func TestGreetingsASCII(t *testing.T) {
	ascii := GreetingASCII()
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 {
			fmt.Println("Not only ascii! ")
			return false
		}
	}
	fmt.Println("only ascii! ")
	return true
}
