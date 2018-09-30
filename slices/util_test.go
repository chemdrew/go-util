package slices

import (
	"testing"
)

func TestConcat(t *testing.T) {
	letters0 := Slice{"a", "b", "c"}
	letters1 := Slice{"d", "e", "f"}
	letters2 := Slice{"g", "h", "i"}
	// expected := Slice{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	t.Run("concats the slices", func(t *testing.T) {
		result := letters0.Concat(letters1, letters2)
		if len(result) != 9 {
			t.Errorf("Slices not concatenated correctly")
		}
	})
}

func TestIncludes(t *testing.T) {
	letters := Slice{"a", "b", "c", "d"}
	t.Run("does include", func(t *testing.T) {
		if !letters.Includes("a") {
			t.Errorf("Expected to find match")
		}
	})

	t.Run("does not include", func(t *testing.T) {
		if letters.Includes("e") {
			t.Errorf("Did not expect to find a match")
		}
	})
}
