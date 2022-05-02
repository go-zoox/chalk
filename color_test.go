package chalk

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	col := Color{31}
	if col.Value() != 31 {
		t.Errorf("Expected 1, got %d", col.Value())
	}

	expected := "\u001b[31m"
	if col.String() != expected {
		t.Errorf("Expected %s, got %s", expected, col.String())
	}

	expected = fmt.Sprintf("%shello%s", expected, ResetColor)
	if col.Color("hello") != expected {
		t.Errorf("Expected %s, got %s", expected, col.Color("hello"))
	}
}
