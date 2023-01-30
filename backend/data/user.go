package data

type User struct {
	email    string
	password string
	role     string // student, teacher, admin
}

var InMemDb = []User{
	{email: "daniel.ciucur@gmail.com", password: "1234", role: "student"},
	{email: "mike.ciucur@gmail.com", password: "1234", role: "teacher"},
	{email: "diego.ciucur@gmail.com", password: "1234", role: "admin"},
}
