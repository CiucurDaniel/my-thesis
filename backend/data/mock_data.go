package data

import uuid "github.com/satori/go.uuid"

var InMemUserDb = []User{
	{id: uuid.NewV4(), email: "daniel.ciucur@gmail.com", password: "1234", role: "student"},
	{id: uuid.NewV4(), email: "mihai.andrei@gmail.com", password: "1234", role: "student"},

	{id: uuid.NewV4(), email: "ionut.ionescu@gmail.com", password: "1234", role: "teacher"},
	{id: uuid.NewV4(), email: "mike.jimmy@gmail.com", password: "1234", role: "teacher"},

	{id: uuid.NewV4(), email: "diego.john@gmail.com", password: "1234", role: "teacher"},
}

var InMemProjectDb = []Project{
	{
		id: uuid.NewV4(),
		name:        "Java Spring web application",
		student:     &InMemUserDb[0],
		teacher:     &InMemUserDb[2],
		description: "This is the project description",
		tags:        []string{"programming", "spring framework", "java"},
	},
	{
		id: uuid.NewV4(),
		name:        "Kubernetes hands-on",
		student:     &InMemUserDb[1],
		teacher:     &InMemUserDb[3],
		description: "This project aims to study the use of Kubernetes.",
		tags:        []string{"kubernetes", "docker", "cloud"},
	},
	{
		id: uuid.NewV4(),
		name:        "Algorithms benchmark 1",
		student:     nil,
		teacher:     &InMemUserDb[4],
		description: "This is the project description",
		tags:        []string{"programming", "algorithms", "python"},
	},
}
