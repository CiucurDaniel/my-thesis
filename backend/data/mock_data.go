package data

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

var InMemUserDb = []User{
	{Id: uuid.NewV4(), Email: "daniel.ciucur@gmail.com", Password: "1234", Role: "Student"},
	{Id: uuid.NewV4(), Email: "mihai.andrei@gmail.com", Password: "1234", Role: "Student"},

	{Id: uuid.NewV4(), Email: "ionut.ionescu@gmail.com", Password: "1234", Role: "Teacher"},
	{Id: uuid.NewV4(), Email: "mike.jimmy@gmail.com", Password: "1234", Role: "Teacher"},

	{Id: uuid.NewV4(), Email: "diego.john@gmail.com", Password: "1234", Role: "Teacher"},
}

var InMemProjectDb = []Project{
	{
		Id:          uuid.NewV4(),
		Name:        "Java Spring web application",
		Student:     &InMemUserDb[0],
		Teacher:     &InMemUserDb[2],
		Description: "This is the project Description",
		Tags:        []string{"programming", "spring framework", "java"},
	},
	{
		Id:          uuid.NewV4(),
		Name:        "Kubernetes hands-on",
		Student:     &InMemUserDb[1],
		Teacher:     &InMemUserDb[3],
		Description: "This project aims to study the use of Kubernetes.",
		Tags:        []string{"kubernetes", "docker", "cloud"},
	},
	{
		Id:          uuid.NewV4(),
		Name:        "Algorithms benchmark 1",
		Student:     nil,
		Teacher:     &InMemUserDb[4],
		Description: "This is the project Description",
		Tags:        []string{"programming", "algorithms", "python"},
	},
	{
		Id:          uuid.NewV4(),
		Name:        "Python OOP vs Java OOP",
		Student:     nil,
		Teacher:     &InMemUserDb[4],
		Description: "Compare OOP aspect between Python and Java",
		Tags:        []string{"OOP", "java", "python"},
	},
	{
		Id:          uuid.NewV4(),
		Name:        "JavaScript React application",
		Student:     nil,
		Teacher:     &InMemUserDb[4],
		Description: "Create a JavaScript delivery application using React",
		Tags:        []string{"google maps", "react", "js"},
	},
}

func init() {
	fmt.Printf("Project uuid %s\n", InMemProjectDb[0].Id)
}