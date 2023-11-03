package reflect

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopy(t *testing.T) {
	type Student struct {
		Name string
		Age  uint8
	}

	s1 := Student{
		Name: "a",
		Age:  1,
	}

	cs1, err := DeepCopy(s1)
	assert.Nil(t, err)
	assert.NotEqual(t, uintptr(unsafe.Pointer(&s1)), uintptr(unsafe.Pointer(&cs1)))
	assert.True(t, reflect.DeepEqual(s1, cs1))
}

type Student struct {
	Name string
	Age  uint8
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) SetAge(age uint8) {
	s.Age = age
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() uint8 {
	return s.Age
}

func (s *Student) GetSelf() *Student {
	return s
}

type Computer struct {
	Name  string
	Price uint32
}

func (c Computer) SetName(name string) {
	c.Name = name
}

func (c Computer) SetPrice(price uint32) {
	c.Price = price
}

func (c Computer) GetName() string {
	return c.Name
}

func (c Computer) GetPrice() uint32 {
	return c.Price
}

func (c Computer) GetSelf() *Computer {
	return &c
}

func TestCallMethod(t *testing.T) {
	s1 := Student{}
	res, err := CallMethod(s1, "SetName", []interface{}{"daniel"})
	assert.EqualError(t, err, "method not found")
	res, err = CallMethod(&s1, "SetName", []interface{}{"daniel"})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res))
	res, err = CallMethod(&s1, "SetAge", []interface{}{uint8(19)})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res))
	res, err = CallMethod(&s1, "GetName", nil)
	assert.Nil(t, err)
	assert.Equal(t, "daniel", res[0].(string))
	res, err = CallMethod(&s1, "GetAge", nil)
	assert.Nil(t, err)
	assert.Equal(t, uint8(19), res[0].(uint8))
	res, err = CallMethod(&s1, "GetSelf", nil)
	assert.Nil(t, err)
	assert.Equal(t, uintptr(unsafe.Pointer(&s1)), uintptr(unsafe.Pointer(res[0].(*Student))))

	c1 := Computer{
		Name:  "unnamed",
		Price: 10,
	}
	res, err = CallMethod(c1, "SetName", []interface{}{"intel"})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res))
	res, err = CallMethod(c1, "SetPrice", []interface{}{uint32(1000)})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res))
	res, err = CallMethod(c1, "GetName", nil)
	assert.Nil(t, err)
	assert.Equal(t, "unnamed", res[0].(string))
	res, err = CallMethod(c1, "GetPrice", nil)
	assert.Nil(t, err)
	assert.Equal(t, uint32(10), res[0].(uint32))
	res, err = CallMethod(c1, "GetSelf", nil)
	assert.Nil(t, err)
	assert.NotEqual(t, uintptr(unsafe.Pointer(&c1)), uintptr(unsafe.Pointer(res[0].(*Computer))))
}
