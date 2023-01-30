package data

import uuid "github.com/satori/go.uuid"

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
	Role     string // Student, Teacher, admin
}

// GetUserById Return a project and an error
// is the project is found err is nil
// otherwise an error will be thrown
func GetUserById(is uuid.UUID) (User, error) {
	for _, user := range InMemUserDb {
		if user.Id == is {
			return user, nil
		}
	}
	return User{}, ErrEntryNotFound
}
