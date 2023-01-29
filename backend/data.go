package main

type user struct {
	email    string
	password string
	role     string // student, teacher, admin
}

var InMemDb = []user{
	{email: "daniel.ciucur@gmail.com", password: "1234", role: "student"},
	{email: "mike.ciucur@gmail.com", password: "1234", role: "teacher"},
	{email: "diego.ciucur@gmail.com", password: "1234", role: "admin"},
}
