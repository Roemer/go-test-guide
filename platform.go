package gotestguide

import (
	"fmt"
	"net/http"
)

type (
	PlatformServiceInterface interface {
		// Retrieve project data.
		GetProject(projectId int) (*Project, *http.Response, error)
	}
	PlatformService struct {
		client *Client
	}
)

var _ PlatformServiceInterface = (*PlatformService)(nil)

func (s *PlatformService) GetProject(projectId int) (*Project, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/platform/projects/%d", projectId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &Project{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}
