package model

//request
type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//response
type LoginUserResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func FormatLoginUserResponse(user User, token string) LoginUserResponse {
	userFormatter := LoginUserResponse{}
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.Role = user.Role.Name
	userFormatter.Token = token

	return userFormatter
}
