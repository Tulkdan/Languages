package abstractfactory

import (
	"testing"
)

func TestAbstract_Factory(t *testing.T) {
    t.Run("should create a factory for adidas", func(t *testing.T) {
        want := "adidas"
        adidasFactory, _ := GetSportsFactory(want)

        adidasShoe := adidasFactory.makeShoe()
        adidasShirt := adidasFactory.makeShirt()

        AssertString(t, adidasShoe.getLogo(), want)
        AssertString(t, adidasShirt.getLogo(), want)
    })

    t.Run("should create a factory for nike", func(t *testing.T) {
        want := "nike"
        nikeFactory, _ := GetSportsFactory(want)

        nikeShoe := nikeFactory.makeShoe()
        nikeShirt := nikeFactory.makeShirt()

        AssertString(t, nikeShoe.getLogo(), want)
        AssertString(t, nikeShirt.getLogo(), want)
    })

    t.Run("shoult throw an error when unmapped brand is passed", func(t *testing.T) {
        _, err := GetSportsFactory("puma")

        if err == nil {
            t.Fatal("expected an error but did not received it")
        }

        AssertString(t, err.Error(), ErrUnmappedBrand)
    })
}

func AssertString(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("expected logo to be %s, got %s", want, got)
    }
}
