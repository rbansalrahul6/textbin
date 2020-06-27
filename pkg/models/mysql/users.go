package mysql

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/rbansalrahul6/textbin/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

/*DB methods for user model*/

// insert new user in DB
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (name, email, hashed_password, created)
	VALUES (?, ?, ?, UTC_TIMESTAMP())`
	_, err = m.DB.Exec(query, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

// verify if user exists in DB -> returns user id if exists
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// fetch user details
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
