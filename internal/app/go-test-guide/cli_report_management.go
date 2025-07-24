package gotestguideapp

import (
	"fmt"

	gotestguide "github.com/roemer/go-test-guide"
)

func UploadReport(client *gotestguide.Client, projectID int, converter, report string) error {
	task, _, err := client.ReportManagement.UploadReport(projectID, converter, report)
	if err != nil {
		return fmt.Errorf("failed to upload report: %w", err)
	}
	fmt.Println("Report uploaded successfully. Task ID:", task.TaskID)
	return nil
}

func AddArtifact(client *gotestguide.Client, tceId int64, filePath string, comment string, category string) error {
	_, err := client.ReportManagement.AddArtifact(tceId, filePath, comment, category)
	if err != nil {
		return fmt.Errorf("failed to add artifact: %w", err)
	}
	fmt.Println("Artifact added successfully")
	return nil
}
