package gotestguide

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type (
	ReportManagementServiceInterface interface {
		// Retrieve information about all available X2ATX converters.
		GetConverters() ([]*Converter, *http.Response, error)
		// Upload a new report.
		UploadReport(projectId int, converterId string, reportPath string) (*TaskRef, *http.Response, error)
		// Uploads a new report from the given objects.
		//UploadReportTyped(projectId int, report *TestGuideReport) (*TaskRef, *http.Response, error)
		// Retrieve all test case executions for the supplied report ID (ATX ID).
		GetTestCaseExecutions(reportId int64) ([]*TestCaseExecutionLink, *http.Response, error)
		// Retrieve details about a specific test case execution.
		GetTestCaseExecution(tceId int64) (*TestCaseExecution, *http.Response, error)
		// Retrieve current state of report upload.
		GetUploadStatus(taskId string) (*UploadStatus, *http.Response, error)
		// Provides metadata for uploaded reports.
		GetHistory(projectId int, startDate time.Time, endTime time.Time, offset int, limit int) ([]*ReportHistoryItem, *http.Response, error)
		// Adds an artifact to an existing test case execution.
		AddArtifact(tceId int64, filePath string, comment string, category string) (*http.Response, error)
	}
	ReportManagementService struct {
		client *Client
	}
)

var _ ReportManagementServiceInterface = (*ReportManagementService)(nil)

func (s *ReportManagementService) GetConverters() ([]*Converter, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "api/report/converter", nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*Converter{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ReportManagementService) UploadReport(projectId int, converterId string, reportPath string) (*TaskRef, *http.Response, error) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)
	// Add the report file to the zip archive.
	fileContent, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file %s: %w", reportPath, err)
	}
	f, err := w.Create(filepath.Base(reportPath))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create zip entry for file %s: %w", reportPath, err)
	}
	_, err = f.Write(fileContent)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to write file %s to zip: %w", reportPath, err)
	}
	// Close the zip
	if err := w.Close(); err != nil {
		return nil, nil, fmt.Errorf("failed to close zip writer: %w", err)
	}

	// Send the request
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("api/report/reports?projectId=%d&converterId=%s", projectId, converterId), bytes.NewReader(buf.Bytes()))
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

/*func (s *ReportManagementService) UploadReportTyped(projectId int, report *TestGuideReport) (*TaskRef, *http.Response, error) {
	reportBytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal report: %w", err)
	}
	file, err := os.CreateTemp("", "report-*.json")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(reportBytes); err != nil {
		return nil, nil, fmt.Errorf("failed to write report to temp file: %w", err)
	}
	if err := file.Close(); err != nil {
		return nil, nil, fmt.Errorf("failed to close temp file: %w", err)
	}
	return s.UploadReport(projectId, "json2atx", file.Name())
}*/

func (s *ReportManagementService) GetHistory(projectId int, startDate time.Time, endDate time.Time, offset int, limit int) ([]*ReportHistoryItem, *http.Response, error) {
	reqUrl := fmt.Sprintf("api/report/reports/history?projectId=%d&startDate=%s&endDate=%s&offset=%d&limit=%d",
		projectId, url.QueryEscape(startDate.Format(time.RFC3339)), url.QueryEscape(endDate.Format(time.RFC3339)), offset, limit)
	req, err := s.client.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*ReportHistoryItem{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ReportManagementService) GetTestCaseExecutions(reportId int64) ([]*TestCaseExecutionLink, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/report/reports/%d", reportId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = []*TestCaseExecutionLink{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ReportManagementService) GetTestCaseExecution(tceId int64) (*TestCaseExecution, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/report/testCaseExecution/%d", tceId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &TestCaseExecution{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ReportManagementService) GetUploadStatus(taskId string) (*UploadStatus, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("api/report/reports/uploadstatus/%s", taskId), nil)
	if err != nil {
		return nil, nil, err
	}
	var responseObject = &UploadStatus{}
	resp, err := s.client.Do(req, &responseObject)
	if err != nil {
		return nil, resp, err
	}
	return responseObject, resp, nil
}

func (s *ReportManagementService) AddArtifact(tceId int64, filePath string, comment string, category string) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Create the multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the file part
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}
	// Add the file content
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}
	// Add additional fields
	if len(comment) > 0 {
		if err = writer.WriteField("comment", comment); err != nil {
			return nil, err
		}
	}
	if len(category) > 0 {
		if err = writer.WriteField("category", category); err != nil {
			return nil, err
		}
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}

	// Create the request
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("api/report/testCaseExecution/%d/artifacts", tceId), body)
	if err != nil {
		return nil, err
	}
	// Set the content type to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Run the request
	resp, err := s.client.Do(req, nil)
	return resp, err
}
