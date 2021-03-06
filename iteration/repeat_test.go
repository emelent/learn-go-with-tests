package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeating(t *testing.T) {
	t.Run("repeat a, 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat c, 8 times", func(t *testing.T) {
		repeated := Repeat("c", 8)
		expected := "cccccccc"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 3)
	fmt.Println(result)
	// Output: bbb
}

func TestStringsRepeat(t *testing.T) {
	actual := strings.Repeat("b", 3)
	expected := "bbb"

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
