package abstractfactory

import "fmt"

type Store struct {
	shoes []Shoe
	shirt []Shirt
}

func (s *Store) PurchaseShoe(shoe Shoe) {
	if s.shoes == nil {
		s.shoes = make([]Shoe, 0)
	}
	s.shoes = append(s.shoes, shoe)
}

func (s *Store) PurchaseShirt(shirt Shirt) {
	if s.shirt == nil {
		s.shirt = make([]Shirt, 0)
	}
	s.shirt = append(s.shirt, shirt)
}

func (s *Store) ListAll() {
	for i, v := range s.shoes {
		fmt.Printf("shoe %d: logo->%s, size->%d\n", i, v.GetLogo(), v.GetSize())
	}

	for i, v := range s.shirt {
		fmt.Printf("shirt %d: logo->%s, size->%d\n", i, v.GetLogo(), v.GetSize())
	}
}
