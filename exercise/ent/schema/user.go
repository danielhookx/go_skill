package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nick_name").Default(""),
		field.String("f1").Default(""),
		field.String("f2").Default(""),
		field.String("f3").Default(""),
		field.String("f4").Default(""),
		field.String("f5").Default(""),
		field.String("f6").Default(""),
		field.String("f7").Default(""),
		field.String("f8").Default(""),
		field.String("f9").Default(""),
		field.String("f10").Default(""),
		field.String("f11").Default(""),
		field.String("f12").Default(""),
		field.String("f13").Default(""),
		field.String("f14").Default(""),
		field.String("f15").Default(""),
		field.String("f16").Default(""),
		field.String("f17").Default(""),
		field.String("f18").Default(""),
		field.String("f19").Default(""),
		field.String("f20").Default(""),
		field.String("f21").Default(""),
		field.String("f22").Default(""),
	}
}
