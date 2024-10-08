package abstractfactory

type Adidas struct {}

func (a *Adidas) makeShoe() IShoe {
    return &AdidasShoes {
        Shoe: Shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a *Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            logo: "adidas",
            size: 14,
        },
    }
}
