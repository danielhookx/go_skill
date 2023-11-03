package abstractfactory

type SportsFactory interface {
	CreateShirt() Shirt
	CreateShoe() Shoe
}
