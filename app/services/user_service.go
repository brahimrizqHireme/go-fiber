package services

import (
	"database/sql"
	// "errors"
	// inputs "github.com/brahimrizqHireme/go-fiber/app/inputes"
	// "github.com/brahimrizqHireme/go-fiber/app/models"
	// "github.com/brahimrizqHireme/go-fiber/app/utils"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// func (s *UserService) CreateUser(input inputs.CreateUserInput) (*models.User, error) {
// 	if err := input.Validate(); err != nil {
// 		return nil, err
// 	}

// 	user := &models.User{
// 		ID:       utils.GenerateUUID(),
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: input.Password,
// 		IsActive: true,
// 	}

// 	insertQuery := `
//         INSERT INTO users (id, name, email, password, isActive)
//         VALUES (?, ?, ?, ?, ?)
//     `
// 	_, err := s.db.Exec(insertQuery, user.ID.String(), user.Name, user.Email, user.Password, user.IsActive)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (s *UserService) GetUserByID(id string) (*models.User, error) {
// 	query := "SELECT id, name, email, password, isActive FROM users WHERE id = ?"
// 	row := s.db.QueryRow(query, id)

// 	var user models.User
// 	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive)
// 	if err == sql.ErrNoRows {
// 		return nil, errors.New("user not found")
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (s *UserService) UpdateUser(user *models.User) error {
// 	updateQuery := `
//         UPDATE users
//         SET name = ?, email = ?, password = ?, isActive = ?
//         WHERE id = ?
//     `
// 	_, err := s.db.Exec(updateQuery, user.Name, user.Email, user.Password, user.IsActive, user.ID.String())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *UserService) DeleteUser(id string) error {
// 	deleteQuery := "DELETE FROM users WHERE id = ?"
// 	_, err := s.db.Exec(deleteQuery, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
