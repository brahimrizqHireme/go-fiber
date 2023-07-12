package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Name      string     `json:"name" gorm:"not null" validate:"required"`
	Email     string     `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password  string     `json:"password,omitempty" gorm:"not null" validate:"required,min=6"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	IsActive  bool       `json:"is_active" gorm:"not null;default:false"`
	ProjectID *uuid.UUID `json:"project_id,omitempty" gorm:"type:uuid;index;default:null"`
	Project   *Project   `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
}

type Project struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
