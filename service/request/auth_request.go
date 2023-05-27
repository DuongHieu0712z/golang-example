package request

type RegisterRequest struct {
	Username        string `json:"username"        binding:"required"`
	Password        string `json:"password"        binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
