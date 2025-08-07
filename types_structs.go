package gotestguide

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////
// Argument
////////////////////////////////////////////////////////////

type Argument struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Direction Direction `json:"direction"`
}

func (a *Argument) String() string {
	return fmt.Sprintf("Argument(Name: %s, Value: %s, Direction: %s)", a.Name, a.Value, a.Direction)
}

////////////////////////////////////////////////////////////
// Artifact
////////////////////////////////////////////////////////////

type Artifact struct {
	ID             string                 `json:"id"`
	FileName       string                 `json:"fileName"`
	Extension      string                 `json:"extension"`
	FileSize       int64                  `json:"fileSize"`
	Hash           string                 `json:"hash"`
	UploadDate     time.Time              `json:"uploadDate"`
	LastAccessDate time.Time              `json:"lastAccessDate"`
	Uploader       string                 `json:"uploader"`
	AttributeList  []*ArtifactAttribute   `json:"attributeList"`
	Shares         []*ArtifactShare       `json:"shares"`
	LockedBy       []*LockedArtifactGroup `json:"lockedBy"`
}

func (a *Artifact) String() string {
	return fmt.Sprintf("Artifact(ID: %s, FileName: %s, Extension: %s, FileSize: %d, Hash: %s, UploadDate: %s, Uploader: %s)",
		a.ID, a.FileName, a.Extension, a.FileSize, a.Hash, a.UploadDate.Format(time.RFC3339), a.Uploader)
}

////////////////////////////////////////////////////////////
// ArtifactAttribute
////////////////////////////////////////////////////////////

type ArtifactAttribute struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func (a *ArtifactAttribute) String() string {
	return fmt.Sprintf("ArtifactAttribute(Key: %s, Values: %v)", a.Key, a.Values)
}

////////////////////////////////////////////////////////////
// ArtifactCreatedResponse
////////////////////////////////////////////////////////////

type ArtifactCreatedResponse struct {
	ID string `json:"artifactId"`
}

func (a *ArtifactCreatedResponse) String() string {
	return fmt.Sprintf("ArtifactCreatedResponse(ID: %s)", a.ID)
}

////////////////////////////////////////////////////////////
// ArtifactRef
////////////////////////////////////////////////////////////

type ArtifactRef struct {
	Ref      string `json:"ref,omitempty"`
	Md5      string `json:"md5,omitempty"`
	FileSize int64  `json:"fileSize,omitempty"`
}

func (a *ArtifactRef) String() string {
	return fmt.Sprintf("ArtifactRef(Ref: %s, Md5: %s, FileSize: %d)", a.Ref, a.Md5, a.FileSize)
}

////////////////////////////////////////////////////////////
// ArtifactShare
////////////////////////////////////////////////////////////

type ArtifactShare struct {
	ID          string    `json:"id"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Creator     string    `json:"creator"`
	CreateDate  time.Time `json:"createDate"`
}

func (a *ArtifactShare) String() string {
	return fmt.Sprintf("ArtifactShare(ID: %s, Link: %s, Description: %s, Creator: %s, CreateDate: %s)",
		a.ID, a.Link, a.Description, a.Creator, a.CreateDate.Format(time.RFC3339))
}

////////////////////////////////////////////////////////////
// Attribute
////////////////////////////////////////////////////////////

type Attribute struct {
	Key    string   `json:"key"`
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

func (a *Attribute) String() string {
	return fmt.Sprintf("Attribute(Key: %s, Value: %s, Values: %v)", a.Key, a.Value, a.Values)
}

////////////////////////////////////////////////////////////
// Constant
////////////////////////////////////////////////////////////

type Constant struct {
	Key    string   `json:"key"`
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

func (c *Constant) String() string {
	return fmt.Sprintf("Constant(Key: %s, Value: %s, Values: %v)", c.Key, c.Value, c.Values)
}

////////////////////////////////////////////////////////////
// Converter
////////////////////////////////////////////////////////////

type Converter struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

func (c *Converter) String() string {
	return fmt.Sprintf("Converter(ID: %s, Version: %s)", c.ID, c.Version)
}

////////////////////////////////////////////////////////////
// DeleteStatus
////////////////////////////////////////////////////////////

type DeleteStatus struct {
	Status          string `json:"status"`
	DetailedMessage string `json:"detailedMessage"`
}

func (d *DeleteStatus) String() string {
	return fmt.Sprintf("DeleteStatus(Status: %s, DetailedMessage: %s)", d.Status, d.DetailedMessage)
}

////////////////////////////////////////////////////////////
// Depository
////////////////////////////////////////////////////////////

type Depository struct {
	ID            string `json:"id"`
	ProjectId     int    `json:"projectId"`
	ActiveStorage int    `json:"activeStorage"`
	Name          string `json:"name"`
}

func (d *Depository) String() string {
	return fmt.Sprintf("Depository(Id: %s, ProjectId: %d, ActiveStorage: %d, Name: %s)", d.ID, d.ProjectId, d.ActiveStorage, d.Name)
}

////////////////////////////////////////////////////////////
// DepositoryIdResponse
////////////////////////////////////////////////////////////

type DepositoryIdResponse struct {
	ID string `json:"id"`
}

func (d *DepositoryIdResponse) String() string {
	return fmt.Sprintf("DepositoryIdResponse(ID: %s)", d.ID)
}

////////////////////////////////////////////////////////////
// FileReference
////////////////////////////////////////////////////////////

type FileReference struct {
	ID          int64     `json:"id"`
	Filename    string    `json:"filename"`
	RelPath     string    `json:"relPath"`
	UploadDate  time.Time `json:"uploadDate"`
	DownloadURL string    `json:"downloadUrl"`
	FileSize    int64     `json:"fileSize"`
	FileHash    string    `json:"fileHash"`
	FilePath    string    `json:"filePath"`
}

func (f *FileReference) String() string {
	return fmt.Sprintf("FileReference(ID: %d, Filename: %s, RelPath: %s, UploadDate: %s, DownloadURL: %s, FileSize: %d, FileHash: %s, FilePath: %s)",
		f.ID, f.Filename, f.RelPath, f.UploadDate.Format(time.RFC3339), f.DownloadURL, f.FileSize, f.FileHash, f.FilePath)
}

////////////////////////////////////////////////////////////
// Filter
////////////////////////////////////////////////////////////

type Filter struct {
	FilterId    int64             `json:"filterId"`
	Name        string            `json:"name"`
	Category    string            `json:"category,omitempty"`
	Description string            `json:"description,omitempty"`
	Parameters  *FilterParameters `json:"parameters,omitempty"`
}

func (f *Filter) String() string {
	return fmt.Sprintf("Filter(FilterId: %d, Name: %s, Category: %s, Description: %s, Parameters: %v)", f.FilterId, f.Name, f.Category, f.Description, f.Parameters)
}

////////////////////////////////////////////////////////////
// FilterInformation
////////////////////////////////////////////////////////////

type FilterInformation struct {
	FilterId    int64  `json:"filterId"`
	Name        string `json:"name"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
}

func (f *FilterInformation) String() string {
	return fmt.Sprintf("FilterInformation(FilterId: %d, Name: %s, Category: %s, Description: %s)", f.FilterId, f.Name, f.Category, f.Description)
}

////////////////////////////////////////////////////////////
// FilterParameters
////////////////////////////////////////////////////////////

type FilterParameters struct {
	TestCaseTagSetID       *int64              `json:"testCaseTagSetId,omitempty"`
	TestSuiteName          []string            `json:"testSuiteName,omitempty"`
	TestCaseName           []string            `json:"testCaseName,omitempty"`
	ParameterSetName       []string            `json:"parameterSetName,omitempty"`
	TestEnvironments       []string            `json:"testEnvironments,omitempty"`
	Attributes             []*KeyValuesFilter  `json:"attributes,omitempty"`
	Constants              []*KeyValuesFilter  `json:"constants,omitempty"`
	ExecutionTimeMin       *int                `json:"executionTimeMin,omitempty"`
	ExecutionTimeMax       *int                `json:"executionTimeMax,omitempty"`
	PlannedTestCaseFolder  []string            `json:"plannedTestCaseFolder,omitempty"`
	DateFrom               *time.Time          `json:"dateFrom,omitempty"`
	DateTo                 *time.Time          `json:"dateTo,omitempty"`
	ArchiveFiles           []string            `json:"archiveFiles,omitempty"`
	TestArgumentExpr       string              `json:"testArgumentExpr,omitempty"`
	TestArgumentDirections []Direction         `json:"testArgumentDirections,omitempty"`
	AtxIds                 []int64             `json:"atxIds,omitempty"`
	Verdicts               []Verdict           `json:"verdicts,omitempty"`
	ReviewExists           string              `json:"reviewExists,omitempty"`
	IncludeObsoleteReviews *bool               `json:"includeObsoleteReviews,omitempty"`
	ReviewAuthor           string              `json:"reviewAuthor,omitempty"`
	ReviewComment          string              `json:"reviewComment,omitempty"`
	ReviewSummary          string              `json:"reviewSummary,omitempty"`
	ReviewVerdicts         []ReviewVerdict     `json:"reviewVerdicts,omitempty"`
	InvalidRuns            *ValidityConstraint `json:"invalidRuns,omitempty"`
	ReviewDefectClass      string              `json:"reviewDefectClass,omitempty"`
	ReviewDefectPriority   string              `json:"reviewDefectPriority,omitempty"`
	ReviewTags             []string            `json:"reviewTags,omitempty"`
	ReviewCustomEvaluation string              `json:"reviewCustomEvaluation"`
	ReviewTickets          []string            `json:"reviewTickets"`
}

////////////////////////////////////////////////////////////
// KeyValuesFilter
////////////////////////////////////////////////////////////

type KeyValuesFilter struct {
	Key     string   `json:"key"`
	Values  []string `json:"values"`
	Negated *bool    `json:"negated,omitempty"`
}

func (k *KeyValuesFilter) String() string {
	return fmt.Sprintf("KeyValuesFilter(Key: %s, Values: %v, Negated: %v)", k.Key, k.Values, k.Negated)
}

////////////////////////////////////////////////////////////
// LockedArtifactGroup
////////////////////////////////////////////////////////////

type LockedArtifactGroup struct {
	Name string `json:"name"`
}

func (l *LockedArtifactGroup) String() string {
	return fmt.Sprintf("LockedArtifactGroup(Name: %s)", l.Name)
}

////////////////////////////////////////////////////////////
// Project
////////////////////////////////////////////////////////////

type Project struct {
	ID          int                 `json:"projectId"`
	Name        string              `json:"projectName"`
	Description string              `json:"projectDescription"`
	IsActive    bool                `json:"isActive"`
	Deleted     ProjectDeletedState `json:"deleted"`
}

func (p *Project) String() string {
	return fmt.Sprintf("Project(ID: %d, Name: %s, Description: %s, IsActive: %t, Deleted: %s)", p.ID, p.Name, p.Description, p.IsActive, p.Deleted)
}

////////////////////////////////////////////////////////////
// ProjectRole
////////////////////////////////////////////////////////////

type ProjectRole struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (o *ProjectRole) String() string {
	return fmt.Sprintf("ProjectRole(ID: %d, Name: %s)", o.ID, o.Name)
}

////////////////////////////////////////////////////////////
// Recording
////////////////////////////////////////////////////////////

type Recording struct {
	Name      string    `json:"name"`
	Direction Direction `json:"direction"`
	FileHash  string    `json:"fileHash"`
}

func (r *Recording) String() string {
	return fmt.Sprintf("Recording(Name: %s, Direction: %s, FileHash: %s)", r.Name, r.Direction, r.FileHash)
}

////////////////////////////////////////////////////////////
// ReportHistoryItem
////////////////////////////////////////////////////////////

type ReportHistoryItem struct {
	ReportID      int64        `json:"reportId"`
	TestPlanName  string       `json:"testPlanName"`
	Status        ReportStatus `json:"status"`
	UploadDate    time.Time    `json:"uploadDate"`
	ExecutionDate time.Time    `json:"executionDate"`
	FileSize      int64        `json:"fileSize"`
}

func (r *ReportHistoryItem) String() string {
	return fmt.Sprintf("ReportHistoryItem(ReportID: %d, TestPlanName: %s, Status: %s, UploadDate: %s, ExecutionDate: %s, FileSize: %d)",
		r.ReportID, r.TestPlanName, r.Status, r.UploadDate.Format(time.RFC3339), r.ExecutionDate.Format(time.RFC3339), r.FileSize)
}

////////////////////////////////////////////////////////////
// Review
////////////////////////////////////////////////////////////

type Review struct {
	ID               int64            `json:"id"`
	ProjectID        int              `json:"projectId"`
	Attachments      []*FileReference `json:"attachments"`
	Summary          string           `json:"summary,omitempty"`
	Comment          string           `json:"comment"`
	Verdict          Verdict          `json:"verdict,omitempty"`
	CustomEvaluation string           `json:"customEvaluation,omitempty"`
	Reviewer         string           `json:"reviewer"`
	ReviewDate       time.Time        `json:"reviewDate"`
	Contacts         []string         `json:"contacts,omitempty"`
	Tickets          []string         `json:"tickets,omitempty"`
	Tags             []string         `json:"tags,omitempty"`
	DefectClass      string           `json:"defectClass,omitempty"`
	DefectPriority   string           `json:"defectPriority,omitempty"`
	InvalidRun       bool             `json:"invalidRun,omitempty"`
}

func (r *Review) String() string {
	return fmt.Sprintf("Review(ID: %d, ProjectID: %d, Verdict: %s, Reviewer: %s, ReviewDate: %s, DefectClass: %s, DefectPriority: %s, InvalidRun: %t)",
		r.ID, r.ProjectID, r.Verdict, r.Reviewer, r.ReviewDate.Format(time.RFC3339), r.DefectClass, r.DefectPriority, r.InvalidRun)
}

////////////////////////////////////////////////////////////
// TaskRef
////////////////////////////////////////////////////////////

type TaskRef struct {
	TaskID string `json:"taskId"`
}

func (t *TaskRef) String() string {
	return fmt.Sprintf("TaskRef(TaskID: %s)", t.TaskID)
}

////////////////////////////////////////////////////////////
// TestCaseExecution
////////////////////////////////////////////////////////////

type TestCaseExecution struct {
	ID                 int64              `json:"id"`
	ProjectID          int                `json:"projectId"`
	ReportID           int64              `json:"reportId"`
	TestSuiteName      string             `json:"testSuiteName"`
	TestCaseName       string             `json:"testCaseName"`
	ExecutionTimestamp time.Time          `json:"executionTimestamp"`
	Verdict            Verdict            `json:"verdict"`
	EffectiveVerdict   Verdict            `json:"effectiveVerdict"`
	TestEnvironments   []*TestEnvironment `json:"testEnvironments"`
	Attributes         []*Attribute       `json:"attributes"`
	Constants          []*Constant        `json:"constants"`
	Arguments          []*Argument        `json:"arguments"`
	Recordings         []*Recording       `json:"recordings"`
	ParameterSet       string             `json:"parameterSet"`
	TestSteps          *TestSteps         `json:"testSteps"`
	LastReview         *Review            `json:"lastReview"`
	Artifacts          []*FileReference   `json:"artifacts"`
	ExecutionTime      int                `json:"executionTime"`
	Releases           []int64            `json:"releases"`
}

func (t *TestCaseExecution) String() string {
	return fmt.Sprintf("TestCaseExecution(ID: %d, ProjectID: %d, ReportID: %d, Verdict: %s)",
		t.ID, t.ProjectID, t.ReportID, t.Verdict)
}

////////////////////////////////////////////////////////////
// TestCaseExecutionLink
////////////////////////////////////////////////////////////

type TestCaseExecutionLink struct {
	TceID int64  `json:"tceId"`
	Rel   string `json:"rel"`
	Href  string `json:"href"`
}

func (l *TestCaseExecutionLink) String() string {
	return fmt.Sprintf("TestCaseExecutionLink(TceID: %d, Rel: %s, Href: %s)", l.TceID, l.Rel, l.Href)
}

////////////////////////////////////////////////////////////
// TestEnvironment
////////////////////////////////////////////////////////////

type TestEnvironment struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc,omitempty"`
}

func (t *TestEnvironment) String() string {
	return fmt.Sprintf("TestEnvironment(Key: %s, Value: %s, Desc: %s)", t.Key, t.Value, t.Desc)
}

////////////////////////////////////////////////////////////
// TestSteps
////////////////////////////////////////////////////////////

type TestSteps struct {
	Setup     []IAbstractTestStep `json:"setup"`
	Execution []IAbstractTestStep `json:"execution"`
	Teardown  []IAbstractTestStep `json:"teardown"`
}

func (t *TestSteps) String() string {
	return fmt.Sprintf("TestSteps(Setup: %d, Execution: %d, Teardown: %d)",
		len(t.Setup), len(t.Execution), len(t.Teardown))
}

// Custom unmarshal function to handle different types of test steps.
func (t *TestSteps) UnmarshalJSON(data []byte) error {
	type Alias TestSteps
	var raw struct {
		Setup     []json.RawMessage `json:"setup"`
		Execution []json.RawMessage `json:"execution"`
		Teardown  []json.RawMessage `json:"teardown"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if steps, err := unmarshalRawTestStep(raw.Setup); err != nil {
		return err
	} else {
		t.Setup = steps
	}

	if steps, err := unmarshalRawTestStep(raw.Execution); err != nil {
		return err
	} else {
		t.Execution = steps
	}

	if steps, err := unmarshalRawTestStep(raw.Teardown); err != nil {
		return err
	} else {
		t.Teardown = steps
	}

	return nil
}

// Helper method to unmarshal raw test steps into the appropriate types.
func unmarshalRawTestStep(raw []json.RawMessage) ([]IAbstractTestStep, error) {
	ret := make([]IAbstractTestStep, len(raw))
	for i, step := range raw {
		// Get the type of the step
		var stepType struct {
			DType TestStepType `json:"dType"`
		}
		if err := json.Unmarshal(step, &stepType); err != nil {
			return nil, err
		}

		// Handle the different types
		switch stepType.DType {
		case TEST_STEP_TYPE_TEST_STEP:
			var ts TestStep
			if err := json.Unmarshal(step, &ts); err != nil {
				return nil, err
			}
			ret[i] = &ts
		case TEST_STEP_TYPE_TEST_STEP_FOLDER:
			var folder TestStepFolder
			if err := json.Unmarshal(step, &folder); err != nil {
				return nil, err
			}
			ret[i] = &folder
		default:
			return nil, fmt.Errorf("unknown type: %s", stepType.DType)
		}
	}
	return ret, nil
}

////////////////////////////////////////////////////////////
// TestStepFolder
////////////////////////////////////////////////////////////

type TestStepFolder struct {
	Name           string              `json:"name"`
	Description    string              `json:"description,omitempty"`
	Verdict        Verdict             `json:"verdict,omitempty"`
	ExpectedResult string              `json:"expectedResult,omitempty"`
	TestSteps      []IAbstractTestStep `json:"teststeps"`
}

func (f *TestStepFolder) String() string {
	return fmt.Sprintf("TestStepFolder(Name: %s, Verdict: %s, ExpectedResult: %s, TestSteps: %d)",
		f.Name, f.Verdict, f.ExpectedResult, len(f.TestSteps))
}

func (f *TestStepFolder) GetType() TestStepType {
	return TEST_STEP_TYPE_TEST_STEP_FOLDER
}

func (f *TestStepFolder) AsTestStep() *TestStep {
	return nil
}

func (f *TestStepFolder) AsTestStepFolder() *TestStepFolder {
	return f
}

// Custom unmarshal function to handle different types of test steps.
func (t *TestStepFolder) UnmarshalJSON(data []byte) error {
	type Alias TestStepFolder
	var raw struct {
		TestSteps []json.RawMessage `json:"teststeps"`
		*Alias
	}
	raw.Alias = (*Alias)(t)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if steps, err := unmarshalRawTestStep(raw.TestSteps); err != nil {
		return err
	} else {
		t.TestSteps = steps
	}

	return nil
}

func (f *TestStepFolder) MarshalJSON() ([]byte, error) {
	type Alias TestStepFolder
	return json.Marshal(&struct {
		Type string `json:"@type"`
		*Alias
	}{
		Type:  strings.ToLower(string(f.GetType())),
		Alias: (*Alias)(f),
	})
}

////////////////////////////////////////////////////////////
// TestStep
////////////////////////////////////////////////////////////

type TestStep struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Verdict        string `json:"verdict,omitempty"`
	ExpectedResult string `json:"expectedResult,omitempty"`
}

func (f *TestStep) String() string {
	return fmt.Sprintf("TestStep(Name: %s, Verdict: %s, ExpectedResult: %s)",
		f.Name, f.Verdict, f.ExpectedResult)
}

func (f *TestStep) GetType() TestStepType {
	return TEST_STEP_TYPE_TEST_STEP
}

func (f *TestStep) AsTestStep() *TestStep {
	return f
}
func (f *TestStep) AsTestStepFolder() *TestStepFolder {
	return nil
}

func (f *TestStep) MarshalJSON() ([]byte, error) {
	type Alias TestStep
	return json.Marshal(&struct {
		Type string `json:"@type"`
		*Alias
	}{
		Type:  strings.ToLower(string(f.GetType())),
		Alias: (*Alias)(f),
	})
}

////////////////////////////////////////////////////////////
// User
////////////////////////////////////////////////////////////

type User struct {
	ID                         int64                 `json:"id"`
	UserName                   string                `json:"userName"`
	DisplayName                string                `json:"displayName"`
	Email                      string                `json:"email"`
	Company                    string                `json:"company"`
	AssociatedProjects         []*UserProjectContext `json:"associatedProjects"`
	SystemGroups               []string              `json:"systemGroups"`
	GlobalPermissions          []string              `json:"globalPermissions"`
	UserType                   UserType              `json:"userType"`
	ConfirmedDisclaimerVersion int                   `json:"confirmedDisclaimerVersion"`
	// An approximate indication of when the user last logged into web interface. Scaled in weeks, months etc. for privacy reasons.
	LastSeen string `json:"lastSeen"`
}

func (u *User) String() string {
	return fmt.Sprintf("User(ID: %d, UserName: %s, DisplayName: %s, UserType: %s, LastSeen: %s)",
		u.ID, u.UserName, u.DisplayName, u.UserType, u.LastSeen)
}

////////////////////////////////////////////////////////////
// UserProjectContext
////////////////////////////////////////////////////////////

type UserProjectContext struct {
	ProjectID               int                  `json:"projectId"`
	ActivationStatus        UserActivationStatus `json:"activationStatus"`
	IndividualPermissions   []string             `json:"individualPermissions"`
	ProjectRoles            []int64              `json:"projectRoles"`
	IndividualProjectRoles  []int64              `json:"individualProjectRoles"`
	SystemGroupProjectRoles []int64              `json:"systemGroupProjectRoles"`
	EffectivePermissions    []string             `json:"effectivePermissions"`
}

func (u *UserProjectContext) String() string {
	return fmt.Sprintf("UserProjectContext(ProjectID: %d, ActivationStatus: %s, EffectivePermissions: %v)",
		u.ProjectID, u.ActivationStatus, u.EffectivePermissions)
}

////////////////////////////////////////////////////////////
// UploadStatus
////////////////////////////////////////////////////////////

type UploadStatus struct {
	Status       string `json:"status"`
	UploadResult struct {
		UploadReturnCode int      `json:"uploadReturnCode"`
		ReportID         int      `json:"reportId"`
		ResultMessages   []string `json:"resultMessages"`
		IsDoubleUpload   bool     `json:"isDoubleUpload"`
	} `json:"uploadResult"`
}

func (u *UploadStatus) String() string {
	return fmt.Sprintf("UploadStatus(Status: %s, UploadReturnCode: %d, ReportID: %d, IsDoubleUpload: %t, ResultMessages: %s)",
		u.Status, u.UploadResult.UploadReturnCode, u.UploadResult.ReportID, u.UploadResult.IsDoubleUpload, strings.Join(u.UploadResult.ResultMessages, "|"))
}
