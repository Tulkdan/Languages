package abstractfactory

import (
	"errors"
)

type ISportFactory interface {
    makeShoe()  IShoe
    makeShirt() IShirt
}

const ErrUnmappedBrand = "Wrong brand type passed"

func GetSportsFactory(brand string) (ISportFactory, error) {
    if brand == "adidas" {
        return &Adidas{}, nil
    }

    if brand == "nike" {
        return &Nike{}, nil
    }

    return nil, errors.New(ErrUnmappedBrand)
}
