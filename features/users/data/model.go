package data

import (
	"capstone/group3/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	AvatarUrl string
	Handphone string `gorm:"unique"`
	Role      string `gorm:"default:user"`
}

// DTO

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromCore(core users.Core) User {
	return User{
		Name:      core.Name,
		Email:     core.Email,
		Password:  core.Password,
		AvatarUrl: core.AvatarUrl,
		Handphone: core.Handphone,
	}
}
