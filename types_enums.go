package gotestguide

////////////////////////////////////////////////////////////
// Direction
////////////////////////////////////////////////////////////

type Direction string

const (
	DIRECTION_IN    Direction = "IN"
	DIRECTION_OUT   Direction = "OUT"
	DIRECTION_INOUT Direction = "INOUT"
)

////////////////////////////////////////////////////////////
// ProjectDeletedState
////////////////////////////////////////////////////////////

type ProjectDeletedState string

func (s ProjectDeletedState) String() string {
	if s == PROJECT_DELETED_STATE_ACTIVE {
		return "ACTIVE"
	}
	return string(s)
}

const (
	PROJECT_DELETED_STATE_ACTIVE      ProjectDeletedState = ""
	PROJECT_DELETED_STATE_IN_PROGRESS ProjectDeletedState = "IN_PROGRESS"
	PROJECT_DELETED_STATE_FINISHED    ProjectDeletedState = "FINISHED"
)

////////////////////////////////////////////////////////////
// ReportStatus
////////////////////////////////////////////////////////////

type ReportStatus string

const (
	REPORT_STATUS_PENDING             ReportStatus = "PENDING"
	REPORT_STATUS_COMPLETE            ReportStatus = "COMPLETE"
	REPORT_STATUS_COMPLETE_WITH_ERROR ReportStatus = "COMPLETE_WITH_ERROR"
)

////////////////////////////////////////////////////////////
// TestStepType
////////////////////////////////////////////////////////////

type TestStepType string

const (
	TEST_STEP_TYPE_TEST_STEP        TestStepType = "TestStep"
	TEST_STEP_TYPE_TEST_STEP_FOLDER TestStepType = "TestStepFolder"
)

////////////////////////////////////////////////////////////
// UserActivationStatus
////////////////////////////////////////////////////////////

type UserActivationStatus string

const (
	USER_ACTIVATION_STATUS_ACTIVATED   UserActivationStatus = "ACTIVATED"
	USER_ACTIVATION_STATUS_IN_PROGRESS UserActivationStatus = "IN_PROGRESS"
	USER_ACTIVATION_STATUS_DEACTIVATED UserActivationStatus = "DEACTIVATED"
)

////////////////////////////////////////////////////////////
// UserType
////////////////////////////////////////////////////////////

type UserType string

const (
	USER_TYPE_REGULAR   UserType = "REGULAR"
	USER_TYPE_TECHNICAL UserType = "TECHNICAL"
)

////////////////////////////////////////////////////////////
// Verdict
////////////////////////////////////////////////////////////

type Verdict string

const (
	VERDICT_NONE         Verdict = "NONE"
	VERDICT_PASSED       Verdict = "PASSED"
	VERDICT_INCONCLUSIVE Verdict = "INCONCLUSIVE"
	VERDICT_FAILED       Verdict = "FAILED"
	VERDICT_ERROR        Verdict = "ERROR"
)
