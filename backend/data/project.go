package data

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type Project struct {
	Id          uuid.UUID
	Name        string
	Teacher     *User
	Student     *User
	Description string
	Tags        []string
}

var ErrEntryNotFound = errors.New("entry was not found")

func GetAllProjects() []Project {
	return InMemProjectDb
}

func IsUnassigned(p Project) bool {
	return p.Student == nil
}

func GetUnassignedProject() []Project {
	var unassignedProjects []Project

	for _, project := range InMemProjectDb {
		if IsUnassigned(project) {
			unassignedProjects = append(unassignedProjects, project)
		}
	}

	return unassignedProjects
}

// GetProjectById Return a project and an error
// is the project is found err is nil
// otherwise an error will be thrown
func GetProjectById(is uuid.UUID) (Project, error) {
	for _, project := range InMemProjectDb {
		if project.Id == is {
			return project, nil
		}
	}
	return Project{}, ErrEntryNotFound
}
