package gotestguide

import "encoding/json"

////////////////////////////////////////////////////////////
// This file contains types that are used to upload a report from a structured object.
////////////////////////////////////////////////////////////

type UploadReport struct {
	// Name of the report
	Name string `json:"name"`
	// Timestamp of the report with milliseconds
	Timestamp                int64                     `json:"timestamp"`
	TestCases                []IAbstractUploadTestCase `json:"testcases"`
	OptionalReportIdentifier string                    `json:"optionalReportIdentifier,omitempty"`
}

type IAbstractUploadTestCase interface {
	GetType() string
}

type UploadTestCaseFolder struct {
	Name      string                    `json:"name"`
	TestCases []IAbstractUploadTestCase `json:"testcases"`
}

func (f *UploadTestCaseFolder) GetType() string {
	return "testcasefolder"
}

func (f *UploadTestCaseFolder) MarshalJSON() ([]byte, error) {
	type Alias UploadTestCaseFolder
	return json.Marshal(&struct {
		Type string `json:"@type"`
		*Alias
	}{
		Type:  f.GetType(),
		Alias: (*Alias)(f),
	})
}

type UploadTestCase struct {
	Name          string  `json:"name"`
	Description   string  `json:"description,omitempty"`
	Verdict       Verdict `json:"verdict"`
	Timestamp     int64   `json:"timestamp"`
	ExecutionTime int     `json:"executionTime,omitempty"`

	Constants  []*Constant  `json:"constants,omitempty"`
	Attributes []*Attribute `json:"attributes,omitempty"`

	SetupTestSteps     []IAbstractTestStep `json:"setupTestSteps,omitempty"`
	ExecutionTestSteps []IAbstractTestStep `json:"executionTestSteps,omitempty"`
	TeardownTestSteps  []IAbstractTestStep `json:"teardownTestSteps,omitempty"`
	Parameters         []*Argument         `json:"parameters,omitempty"`
	Artifacts          []string            `json:"artifacts,omitempty"`
	ArtifactRefs       []*ArtifactRef      `json:"artifactRefs,omitempty"`
	Review             *Review             `json:"review,omitempty"`
	ParamSet           string              `json:"paramSet,omitempty"`
	Environments       []*TestEnvironment  `json:"environments,omitempty"`
	Recordings         []*Recording        `json:"recordings,omitempty"`
}

func (f *UploadTestCase) GetType() string {
	return "testcase"
}

func (f *UploadTestCase) MarshalJSON() ([]byte, error) {
	type Alias UploadTestCase
	return json.Marshal(&struct {
		Type string `json:"@type"`
		*Alias
	}{
		Type:  f.GetType(),
		Alias: (*Alias)(f),
	})
}
