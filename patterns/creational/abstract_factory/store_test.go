package abstractfactory

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestListAll(t *testing.T) {
	s := Store{}
	nike := NikeFactory{}
	adidas := AdidasFactory{}
	s.PurchaseShirt(nike.CreateShirt())
	s.PurchaseShoe(nike.CreateShoe())
	s.PurchaseShirt(adidas.CreateShirt())
	s.PurchaseShoe(adidas.CreateShoe())
	s.ListAll()
}
