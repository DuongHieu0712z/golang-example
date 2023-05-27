package request

import (
	"example/data/entity"
	"example/utils"
)

type RegisterRequest struct {
	Username        string `json:"username"        form:"username"        binding:"required,alphanum"`
	Password        string `json:"password"        form:"password"        binding:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"eqfield=Password"`
}

func (req RegisterRequest) Map(data *entity.User) {
	if data == nil {
		panic("register request is null")
	}

	data.Username = req.Username
	data.PasswordHash = utils.HashPassword(req.Password)
}
