package chalk

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println(Success("Success"))
	fmt.Println(Info("Info"))
	fmt.Println(Warn("Warn"))
	fmt.Println(Error("Error"))
	fmt.Println(Debug("Debug"))
}
