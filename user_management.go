package gotestguide

import (
	"fmt"
	"net/http"
)

type (
	UserManagementServiceInterface interface {
		// Retrieve information about the current user.
		Whoami() (*User, *http.Response, error)
		// Retrieve information on all users.
		GetUsers() ([]*User, *http.Response, error)
		// Retrieve all project roles for a project.
		GetRoles(projectId int) ([]*ProjectRole, *http.Response, error)
	}
	UserManagementService struct {
		client *Client
	}
)

var _ UserManagementServiceInterface = (*UserManagementService)(nil)

func (s *UserManagementService) Whoami() (*User, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "api/userManagement/whoami", nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &User{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *UserManagementService) GetUsers() ([]*User, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "api/userManagement/users", nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*User{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *UserManagementService) GetRoles(projectId int) ([]*ProjectRole, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/userManagement/roles?projectId=%d", projectId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*ProjectRole{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}
