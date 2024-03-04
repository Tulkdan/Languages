package iteration

import "testing"

func TestRepeat(t *testing.T) {
    t.Run("should repeat the number of times passed to function", func(t *testing.T) {
        repeated := Repeat("a", 3)
        expected := "aaa"

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
