# go-test-guide
A library / cmd to interact with Test.Guide.

## CLI
This project also provides a compiled CLI binary (based on the Go Module) that allows to directly interact with Test.Guide.

### Usage
```
go-test-guide [global options] [command [command options]]
```

### Server and Authentication
To define the base url of the server and the token, you can either pass `--base-url` and `--token` to the requests or set the environment variables `TEST_GUIDE_BASE_URL` and `TEST_GUIDE_TOKEN`.

### Commands
* `report-management` (`rm`): Manage reports
  * `upload-report`: Upload a new report
  * `add-artifact`: Add a new artifact

### Examples
Upload a report:
```
go-test-guide rm upload-report --project 111 --converter JUnitMatlab --report test-report.xml --token "<token>" --base-url "https://test-guide.mydomain.com"
```

## Go Module
This repository provides a Go module that can be used in your Go applications to interact with Test.Guide.

### Installation
```
go get github.com/roemer/go-test-guide
```

### Modules
* `Artifacts`
  * `CreateDepository`
  * `GetDepositories`
  * `GetDepository`
  * `DeleteDepository`
  * `UploadArtifact`
  * `GetArtifact`
  * `GetStorages`
  * `GetStorage`
  * `CreateStorage`
  * `DeleteStorage`
  * `ActivateStorage`
  * `DeactivateStorage`
* `Platform`
  * `GetProject`
* `ReportManagement`
  * `GetConverters`
  * `UploadReport`
  * `DeleteReport`
  * `GetTestCaseExecutions`
  * `GetTestCaseExecution`
  * `GetUploadStatus`
  * `GetDeleteStatus`
  * `GetHistory`
  * `AddArtifact`
* `UserManagement`
  * `Whoami`
  * `GetUsers`
  * `GetRoles`

### Usage
Create a client:
```go
client, err := gotestguide.NewClient("server-url", "token")
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}
```

Get a project:
```go
project, _, err := client.Platform.GetProject(projectId)
if err != nil {
    return err
}
```
