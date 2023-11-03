package abstractfactory

type NikeFactory struct {
}

func (f *NikeFactory) CreateShirt() Shirt {
	return &NikeShirt{
		logo: "nike",
		size: 14,
	}
}

func (f *NikeFactory) CreateShoe() Shoe {
	return &NikeShoe{
		logo: "nike",
		size: 15,
	}
}

type NikeShirt struct {
	logo string
	size int
}

func (s *NikeShirt) GetLogo() string {
	return s.logo
}

func (s *NikeShirt) GetSize() int {
	return s.size
}

type NikeShoe struct {
	logo string
	size int
}

func (s *NikeShoe) GetLogo() string {
	return s.logo
}

func (s *NikeShoe) GetSize() int {
	return s.size
}
