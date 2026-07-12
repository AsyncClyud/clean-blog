package storage

import (
	"blog/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CheckIfUserExist(User models.User) (user_id int, hashed_password string, success bool) {
	var user models.User
	var users []models.User
	rows, err := ur.db.Query("SELECT Id, Password FROM users WHERE Username = $1", User.Username)
	if err != nil {
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Password)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return 0, "", false
	} else {
		return user.Id, user.Password, true
	}

}

func (ur *UserRepository) CreateUser(user models.User) (success bool) {
	_, _, UserExist := ur.CheckIfUserExist(user)
	if !UserExist {
		_, err := ur.db.Exec("INSERT INTO Users(Username, Password, Created_at) VALUES($1, $2, $3)", user.Username, user.Password, time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		return true
	} else {
		return false
	}
}

func (ur *UserRepository) GetUserInfo(user_id int) (user_info string, err error) {
	rows, err := ur.db.Query("SELECT Id, Username, Bio, Created_at FROM Users WHERE Id = $1", user_id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Bio, &user.Created_at)
		if err != nil {
			log.Fatalln(err)
			rows.Err()
		}
	}
	result, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("Error: %v", err)
	}

	return string(result), nil
}
