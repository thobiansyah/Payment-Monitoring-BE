package model

import "gorm.io/gorm"

//entity
type User struct {
	gorm.Model
	ID       uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"not null"`
	Role     Role
}

//request

//response
type UserResponse struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Role     RoleResponse `json:"role"`
}

func FormatGetAllUserResponse(users []User) []UserResponse {
	usersFormatter := []UserResponse{}

	for _, user := range users {
		userFormatter := UserResponse{}
		userFormatter.ID = user.ID
		userFormatter.Name = user.Name
		userFormatter.Username = user.Username
		userFormatter.Role.ID = user.Role.ID
		userFormatter.Role.Name = user.Role.Name

		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
