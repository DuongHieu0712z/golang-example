package entity

import "example/common/base"

type User struct {
	base.Entity `bson:",inline"`

	Username     string `bson:"username"`
	PasswordHash string `bson:"passwordHash"`
}
