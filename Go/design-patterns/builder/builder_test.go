package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {

    tests := []struct{
        builder  string
        expected House
    }{
        {
            builder: "normal",
            expected: House{
                windowType: "Wooden Window",
                doorType:   "Wooden Door",
                floor:      2,
            },
        },
        {
            builder: "igloo",
            expected: House{
                windowType: "Snow Window",
                doorType:   "Snow Door",
                floor:      1,
            },
        },
    }

    for _, test := range tests {
        t.Run(fmt.Sprintf("should create a %s house with a builder using a director", test.builder), func(t *testing.T) {
            builder := getBuilder(test.builder)
            director := newDirector(builder)
            house := director.buildHouse()

            assertHouse(t, house, test.expected)
        })
    }
}

func assertHouse(t testing.TB, got, want House) {
    t.Helper()

    if got != want {
        t.Errorf("expected house to be %+v, got %+v", want, got)
    }
}
