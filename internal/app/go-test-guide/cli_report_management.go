package gotestguideapp

import (
	"fmt"

	gotestguide "github.com/roemer/go-test-guide"
)

func UploadReport(client *gotestguide.Client, projectId int, converter, report string) error {
	task, _, err := client.ReportManagement.UploadReport(projectId, converter, report)
	if err != nil {
		return fmt.Errorf("failed to upload report: %w", err)
	}
	fmt.Println("Report uploaded successfully. Task ID:", task.TaskID)
	return nil
}

func DeleteReport(client *gotestguide.Client, reportId int64) error {
	task, _, err := client.ReportManagement.DeleteReport(reportId)
	if err != nil {
		return fmt.Errorf("failed to delete report: %w", err)
	}
	fmt.Println("Report deleted successfully. Task ID:", task.TaskID)
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
