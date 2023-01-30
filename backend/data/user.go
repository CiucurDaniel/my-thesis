package data

import uuid "github.com/satori/go.uuid"

type User struct {
	id uuid.UUID
	email    string
	password string
	role     string // student, teacher, admin
}

// GetUserById Return a project and an error
// is the project is found err is nil
// otherwise an error will be thrown
func GetUserById(is uuid.UUID) (User, error) {
	for _, user := range InMemUserDb {
		if user.id == is {
			return user, nil
		}
	}
	return User{}, ErrEntryNotFound
}