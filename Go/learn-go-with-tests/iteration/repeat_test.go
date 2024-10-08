package iteration

import (
    "testing"
    "fmt"
)

func TestRepeat(t *testing.T) {
    t.Run("should repeat the number of times passed to function", func(t *testing.T) {
        repeated := Repeat("a", 3)
        expected := "aaa"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("should only repeat the first character of a string", func(t *testing.T) {
        repeated := Repeat("first", 5)
        expected := "fffff"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("should trims string", func(t *testing.T) {
        repeated := Repeat("  spaaaace  ", 2)
        expected := "ss"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("b", 2)
    fmt.Println(repeated)
    // Output: bb
}
