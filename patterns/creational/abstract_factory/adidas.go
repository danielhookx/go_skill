package abstractfactory

type AdidasFactory struct {
}

func (f *AdidasFactory) CreateShirt() Shirt {
	return &AdidasShirt{
		logo: "adidas",
		size: 20,
	}
}

func (f *AdidasFactory) CreateShoe() Shoe {
	return &AdidasShoe{
		logo: "adidas",
		size: 21,
	}
}

type AdidasShirt struct {
	logo string
	size int
}

func (s *AdidasShirt) GetLogo() string {
	return s.logo
}

func (s *AdidasShirt) GetSize() int {
	return s.size
}

type AdidasShoe struct {
	logo string
	size int
}

func (s *AdidasShoe) GetLogo() string {
	return s.logo
}

func (s *AdidasShoe) GetSize() int {
	return s.size
}
