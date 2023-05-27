package entity

import "example/common/base"

type Todo struct {
	base.Entity `bson:",inline"`

	Name string `bson:"name"`
}
