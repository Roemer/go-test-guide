package gotestguide

////////////////////////////////////////////////////////////
// AwsS3StorageClass
////////////////////////////////////////////////////////////

type AwsS3StorageClass string

const (
	AWS_S3_STORAGE_CLASS_STANDARD                   AwsS3StorageClass = "STANDARD"
	AWS_S3_STORAGE_CLASS_REDUCED_REDUNDANCY         AwsS3StorageClass = "REDUCED_REDUNDANCY"
	AWS_S3_STORAGE_CLASS_GLACIER                    AwsS3StorageClass = "GLACIER"
	AWS_S3_STORAGE_CLASS_STANDARD_INFREQUENT_ACCESS AwsS3StorageClass = "STANDARD_IA"
	AWS_S3_STORAGE_CLASS_ONE_ZONE_INFREQUENT_ACCESS AwsS3StorageClass = "ONEZONE_IA"
	AWS_S3_STORAGE_CLASS_INTELLIGENT_TIERING        AwsS3StorageClass = "INTELLIGENT_TIERING"
	AWS_S3_STORAGE_CLASS_DEEP_ARCHIVE               AwsS3StorageClass = "DEEP_ARCHIVE"
)

////////////////////////////////////////////////////////////
// AzureAuthenticationType
////////////////////////////////////////////////////////////

type AzureAuthenticationType string

const (
	AZURE_AUTHENTICATION_TYPE_BASIC      AzureAuthenticationType = "BASIC"
	AZURE_AUTHENTICATION_TYPE_SAS        AzureAuthenticationType = "SAS"
	AZURE_AUTHENTICATION_TYPE_SHARED_KEY AzureAuthenticationType = "SHARED_KEY"
)

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
// ReviewVerdict
////////////////////////////////////////////////////////////

type ReviewVerdict string

const (
	REVIEW_VERDICT_NONE         ReviewVerdict = "NONE"
	REVIEW_VERDICT_PASSED       ReviewVerdict = "PASSED"
	REVIEW_VERDICT_INCONCLUSIVE ReviewVerdict = "INCONCLUSIVE"
	REVIEW_VERDICT_FAILED       ReviewVerdict = "FAILED"
	REVIEW_VERDICT_ERROR        ReviewVerdict = "ERROR"
	REVIEW_VERDICT_NO_VERDICT   ReviewVerdict = "NO_VERDICT"
)

////////////////////////////////////////////////////////////
// SftpAuthenticationType
////////////////////////////////////////////////////////////

type SftpAuthenticationType string

const (
	SFTP_AUTHENTICATION_TYPE_BASIC   SftpAuthenticationType = "BASIC"
	SFTP_AUTHENTICATION_TYPE_SSH_KEY SftpAuthenticationType = "SSH_KEY"
)

////////////////////////////////////////////////////////////
// SmbDialect
////////////////////////////////////////////////////////////

type SmbDialect string

const (
	SMB_DIALECT_2_0_2 SmbDialect = "SMB_2_0_2"
	SMB_DIALECT_2_1   SmbDialect = "SMB_2_1"
	SMB_DIALECT_2XX   SmbDialect = "SMB_2XX"
	SMB_DIALECT_3_0   SmbDialect = "SMB_3_0"
	SMB_DIALECT_3_0_2 SmbDialect = "SMB_3_0_2"
	SMB_DIALECT_3_1_1 SmbDialect = "SMB_3_1_1"
)

////////////////////////////////////////////////////////////
// StorageType
////////////////////////////////////////////////////////////

type StorageType string

const (
	STORAGE_TYPE_FILE        StorageType = "fileStorage"
	STORAGE_TYPE_SMB         StorageType = "smbStorage"
	STORAGE_TYPE_ARTIFACTORY StorageType = "artifactoryStorage"
	STORAGE_TYPE_AWSS3       StorageType = "awsS3Storage"
	STORAGE_TYPE_SFTP        StorageType = "sftpStorage"
	STORAGE_TYPE_AZUREBLOB   StorageType = "azureBlobStorage"
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
// ValidityConstraint
////////////////////////////////////////////////////////////

type ValidityConstraint string

const (
	VALIDITY_CONSTRAINT_NO_CONSTRAINT ValidityConstraint = "NO_CONSTRAINT"
	VALIDITY_CONSTRAINT_ONLY_VALID    ValidityConstraint = "ONLY_VALID"
	VALIDITY_CONSTRAINT_ONLY_INVALID  ValidityConstraint = "ONLY_INVALID"
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
