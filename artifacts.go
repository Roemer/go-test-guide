package gotestguide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type (
	ArtifactsServiceInterface interface {
		// Create a depository.
		CreateDepository(projectId int, depositoryId string, depositoryName string) (*DepositoryIdResponse, *http.Response, error)
		// Retrieve all depositories.
		GetDepositories(projectId int) ([]*Depository, *http.Response, error)
		// Get all information of a depository.
		GetDepository(depositoryId string) (*Depository, *http.Response, error)
		// Delete a depository.
		DeleteDepository(depositoryId string) (*http.Response, error)
		// Upload an artifact.
		UploadArtifact(depositoryId string, artifactPath string, attributes ...*Attribute) (*ArtifactCreatedResponse, *http.Response, error)
		// Get all information of an artifact.
		GetArtifact(artifactId string) (*Artifact, *http.Response, error)
		// Get all storages of a given depository.
		GetStorages(depositoryId string) ([]IStorage, *http.Response, error)
		// Get all information of a storage.
		GetStorage(depositoryId string, storageNumber int) (IStorage, *http.Response, error)
		// Create Storage.
		CreateStorage(depositoryId string, storage IStorage) (*StorageNumberResponse, *http.Response, error)
		// Delete the given storage. Files in the storage are not automatically removed.
		DeleteStorage(depositoryId string, storageNumber int, removeAllFilesFromStorage *bool) (*TaskRef, *http.Response, error)
		// Activate this storage.
		ActivateStorage(depositoryId string, storageNumber int) (*http.Response, error)
		// Deactivate the currently active storage in this depository.
		DeactivateStorage(depositoryId string) (*http.Response, error)
	}
	ArtifactsService struct {
		client *Client
	}
)

var _ ArtifactsServiceInterface = (*ArtifactsService)(nil)

func (s *ArtifactsService) CreateDepository(projectId int, depositoryId string, depositoryName string) (*DepositoryIdResponse, *http.Response, error) {
	newDepository := &Depository{
		ID:   depositoryId,
		Name: depositoryName,
	}
	// Prepare the body
	bodyBytes, err := json.Marshal(newDepository)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal object: %w", err)
	}

	// Prepare the request
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("api/artifact/depositories?projectId=%d", projectId), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	var responseObject = &DepositoryIdResponse{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) GetDepositories(projectId int) ([]*Depository, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/artifact/depositories?projectId=%d", projectId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*Depository{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) GetDepository(depositoryId string) (*Depository, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/artifact/depositories/%s", depositoryId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &Depository{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) DeleteDepository(depositoryId string) (*http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("api/artifact/depositories/%s", depositoryId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *ArtifactsService) UploadArtifact(depositoryId string, artifactPath string, attributes ...*Attribute) (*ArtifactCreatedResponse, *http.Response, error) {
	// Prepare the url
	reqUrl := fmt.Sprintf("api/artifact/artifacts?depositoryId=%s", depositoryId)
	for _, attr := range attributes {
		reqUrl += fmt.Sprintf("&attributes=%s%%3D%s", url.QueryEscape(attr.Key), url.QueryEscape(attr.Value))
	}

	// Read the artifact file
	file, err := os.Open(artifactPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file %s: %w", artifactPath, err)
	}
	defer file.Close()

	// Prepare the body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the file part
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, nil, err
	}
	// Add the file content
	if _, err = io.Copy(part, file); err != nil {
		return nil, nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}

	// Prepare the request
	req, err := s.client.NewRequest(http.MethodPost, reqUrl, body)
	if err != nil {
		return nil, nil, err
	}

	// Send the request
	var responseObject = &ArtifactCreatedResponse{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) GetArtifact(artifactId string) (*Artifact, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/artifact/artifacts/%s", artifactId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &Artifact{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) GetStorages(depositoryId string) ([]IStorage, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/artifact/depositories/%s/storages", depositoryId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = storages{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject.Storages, resp, nil
}

func (s *ArtifactsService) GetStorage(depositoryId string, storageNumber int) (IStorage, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/artifact/depositories/%s/storages/%d", depositoryId, storageNumber), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject json.RawMessage
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	storage, err := unmarshalStorage(responseObject)
	if err != nil {
		return nil, nil, err
	}
	return storage, resp, nil
}

func (s *ArtifactsService) CreateStorage(depositoryId string, storage IStorage) (*StorageNumberResponse, *http.Response, error) {
	// Prepare the body
	bodyBytes, err := json.Marshal(storage)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal object: %w", err)
	}

	// Prepare the request
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("api/artifact/depositories/%s/storages", depositoryId), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	var responseObject = &StorageNumberResponse{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) DeleteStorage(depositoryId string, storageNumber int, removeAllFilesFromStorage *bool) (*TaskRef, *http.Response, error) {
	urlObject := url.URL{
		Path: fmt.Sprintf("api/artifact/depositories/%s/storages/%d", depositoryId, storageNumber),
	}
	if removeAllFilesFromStorage != nil {
		query := url.Values{}
		query.Set("removeAllFilesFromStorage", strconv.FormatBool(*removeAllFilesFromStorage))
		urlObject.RawQuery = query.Encode()
	}
	req, err := s.client.NewRequest(http.MethodDelete, urlObject.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &TaskRef{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ArtifactsService) ActivateStorage(depositoryId string, storageNumber int) (*http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("api/artifact/depositories/%s/storages/%d/activate", depositoryId, storageNumber), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

func (s *ArtifactsService) DeactivateStorage(depositoryId string) (*http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("api/artifact/depositories/%s/storages/deactivate", depositoryId), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
