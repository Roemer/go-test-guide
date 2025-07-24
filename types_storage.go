package gotestguide

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////
// storages
////////////////////////////////////////////////////////////

// Wrapper object for storages (for unmarshalling).
type storages struct {
	Storages []IStorage
}

func (s *storages) UnmarshalJSON(data []byte) error {
	// Raw unmarshall the storages into a slice
	var rawStorages []json.RawMessage
	if err := json.Unmarshal(data, &rawStorages); err != nil {
		return err
	}
	// Iterate over the raw messages and unmarshal them into the correct storage type
	s.Storages = make([]IStorage, len(rawStorages))
	for i, raw := range rawStorages {
		// Unmarshal the storage type
		var storageType struct {
			StorageType StorageType `json:"storageType"`
		}
		if err := json.Unmarshal(raw, &storageType); err != nil {
			return err
		}
		// Create the appropriate storage type based on the type
		var storage IStorage
		switch storageType.StorageType {
		case STORAGE_TYPE_FILE:
			storage = &StorageFile{}
		case STORAGE_TYPE_SMB:
			storage = &StorageSmb{}
		case STORAGE_TYPE_ARTIFACTORY:
			storage = &StorageArtifactory{}
		case STORAGE_TYPE_AWSS3:
			storage = &StorageAwsS3{}
		case STORAGE_TYPE_SFTP:
			storage = &StorageSftp{}
		case STORAGE_TYPE_AZUREBLOB:
			storage = &StorageAzureBlob{}
		default:
			return fmt.Errorf("unknown storage type: %s", storageType.StorageType)
		}
		// Unmarshal the raw data into the storage type
		if err := json.Unmarshal(raw, &storage); err != nil {
			return err
		}
		// Assign the storage to the slice
		s.Storages[i] = storage
	}
	return nil
}

////////////////////////////////////////////////////////////
// StorageBase
////////////////////////////////////////////////////////////

type IStorage interface {
	GetType() StorageType
	AsFileStorage() *StorageFile
	AsSmbStorage() *StorageSmb
	AsArtifactoryStorage() *StorageArtifactory
	AsAwsS3Storage() *StorageAwsS3
	AsSftpStorage() *StorageSftp
	AsAzureBlobStorage() *StorageAzureBlob
}

// Base structure for all storage types.
type StorageBase struct {
	StorageType                           StorageType                   `json:"storageType"`
	StorageNumber                         int                           `json:"storageNumber"`
	Name                                  string                        `json:"name"`
	KeepFileInStorageWhenDeletingArtifact bool                          `json:"keepFileInStorageWhenDeletingArtifact"`
	MigrationRole                         string                        `json:"migrationRole"`
	DeletionState                         string                        `json:"deletionState"`
	Quota                                 *StorageQuota                 `json:"quota"`
	ConnectionCheck                       *StorageConnectionCheckConfig `json:"connectionCheck"`
}

func (s *StorageBase) GetType() StorageType {
	return s.StorageType
}

func (s *StorageBase) AsFileStorage() *StorageFile {
	return nil
}

func (s *StorageBase) AsSmbStorage() *StorageSmb {
	return nil
}

func (s *StorageBase) AsArtifactoryStorage() *StorageArtifactory {
	return nil
}

func (s *StorageBase) AsAwsS3Storage() *StorageAwsS3 {
	return nil
}

func (s *StorageBase) AsSftpStorage() *StorageSftp {
	return nil
}

func (s *StorageBase) AsAzureBlobStorage() *StorageAzureBlob {
	return nil
}

////////////////////////////////////////////////////////////
// StorageConnectionCheckConfig
////////////////////////////////////////////////////////////

type StorageConnectionCheckConfig struct {
	CronExpression                 string `json:"cronExpression"`
	TimeZone                       string `json:"timeZone"`
	NotifyProjectManagersOnFailure bool   `json:"notifyProjectManagersOnFailure"`
}

////////////////////////////////////////////////////////////
// StorageQuota
////////////////////////////////////////////////////////////

type StorageQuota struct {
	LimitInGiB                                int  `json:"limitInGiB"`
	RejectUpload                              bool `json:"rejectUpload"`
	NotifyDepositoryManagerThresholdInPercent int  `json:"notifyDepositoryManagerThresholdInPercent"`
	NotifyDepositoryUsersThresholdInPercent   int  `json:"notifyDepositoryUsersThresholdInPercent"`
}

////////////////////////////////////////////////////////////
// StorageFile
////////////////////////////////////////////////////////////

// Depository backing storage in local files.
type StorageFile struct {
	*StorageBase
	// Path to local folder where the files should be stored.
	Folder string `json:"folder,omitempty"`
}

func (s *StorageFile) AsFileStorage() *StorageFile {
	return s
}

////////////////////////////////////////////////////////////
// StorageSmb
////////////////////////////////////////////////////////////

// Depository backing storage on the basis of Server Message Block (SMB).
type StorageSmb struct {
	*StorageBase
	UserName                   string     `json:"userName,omitempty"`
	Password                   string     `json:"password,omitempty"`
	Domain                     string     `json:"domain,omitempty"`
	Host                       string     `json:"host,omitempty"`
	Port                       int        `json:"port,omitempty"`
	Share                      string     `json:"share,omitempty"`
	FolderPath                 string     `json:"folderPath,omitempty"`
	DfsEnabled                 bool       `json:"dfsEnabled,omitempty"`
	Dialect                    SmbDialect `json:"dialect,omitempty"`
	TransportEncryptionEnabled bool       `json:"transportEncryptionEnabled,omitempty"`
}

func (s *StorageSmb) AsSmbStorage() *StorageSmb {
	return s
}

////////////////////////////////////////////////////////////
// StorageArtifactory
////////////////////////////////////////////////////////////

// Depository backing storage on the basis of JFrog Artifactory.
type StorageArtifactory struct {
	*StorageBase
	URL               string `json:"url"`
	RepoKey           string `json:"repoKey"`
	UserName          string `json:"userName"`
	ApiKey            string `json:"apiKey"`
	ConnectionTimeout int    `json:"connectionTimeout"`
	SocketTimeout     int    `json:"socketTimeout"`
}

func (s *StorageArtifactory) AsArtifactoryStorage() *StorageArtifactory {
	return s
}

////////////////////////////////////////////////////////////
// StorageAwsS3
////////////////////////////////////////////////////////////

// Depository backing storage on the basis of AWS S3.
type StorageAwsS3 struct {
	*StorageBase
	BucketName        string            `json:"bucketName"`
	CustomEndpoint    string            `json:"customEndpoint"`
	UserName          string            `json:"userName"`
	Password          string            `json:"password"`
	ObjectKeyPrefix   string            `json:"objectKeyPrefix"`
	StorageClass      AwsS3StorageClass `json:"storageClass"`
	AwsRegion         string            `json:"awsRegion"`
	ConnectionTimeout int               `json:"connectionTimeout"`
	SocketTimeout     int               `json:"socketTimeout"`
}

func (s *StorageAwsS3) AsAwsS3Storage() *StorageAwsS3 {
	return s
}

////////////////////////////////////////////////////////////
// StorageSftp
////////////////////////////////////////////////////////////

// Depository backing storage on the basis of Secure File Transfer Protocol (SFTP).
type StorageSftp struct {
	*StorageBase
	Host               string                  `json:"host"`
	Port               int                     `json:"port"`
	AuthenticationInfo ISftpAuthenticationInfo `json:"authenticationInfo"`
	FolderPath         string                  `json:"folderPath"`
}

func (s *StorageSftp) AsSftpStorage() *StorageSftp {
	return s
}

func (s *StorageSftp) UnmarshalJSON(data []byte) error {
	type Alias StorageSftp
	aux := &struct {
		AuthRaw json.RawMessage `json:"authenticationInfo"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	// Unmarshal to extract raw "authenticationInfo" block
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Peek at the "type" field
	var typePeek struct {
		Type SftpAuthenticationType `json:"type"`
	}
	if err := json.Unmarshal(aux.AuthRaw, &typePeek); err != nil {
		return err
	}
	// Create the appropriate storage type based on the type
	var authInfo ISftpAuthenticationInfo
	switch typePeek.Type {
	case SFTP_AUTHENTICATION_TYPE_BASIC:
		authInfo = &SftpAuthenticationInfoBasic{}
	case SFTP_AUTHENTICATION_TYPE_SSH_KEY:
		authInfo = &SftpAuthenticationInfoSshKey{}
	default:
		return fmt.Errorf("unsupported authentication type: %s", typePeek.Type)
	}
	// Unmarshal the raw data into the storage type
	if err := json.Unmarshal(aux.AuthRaw, &authInfo); err != nil {
		return err
	}
	// Assign the storage to the slice
	s.AuthenticationInfo = authInfo
	return nil
}

type ISftpAuthenticationInfo interface {
	GetType() SftpAuthenticationType
	AsBasic() *SftpAuthenticationInfoBasic
	AsSshKey() *SftpAuthenticationInfoSshKey
}

type SftpAuthenticationInfo struct {
	Type     SftpAuthenticationType `json:"type"`
	UserName string                 `json:"userName"`
}

func (s *SftpAuthenticationInfo) GetType() SftpAuthenticationType {
	return s.Type
}

func (s *SftpAuthenticationInfo) AsBasic() *SftpAuthenticationInfoBasic {
	return nil
}

func (s *SftpAuthenticationInfo) AsSshKey() *SftpAuthenticationInfoSshKey {
	return nil
}

type SftpAuthenticationInfoBasic struct {
	*SftpAuthenticationInfo
	Password string `json:"password"`
}

func (s *SftpAuthenticationInfoBasic) AsBasic() *SftpAuthenticationInfoBasic {
	return s
}

type SftpAuthenticationInfoSshKey struct {
	*SftpAuthenticationInfo
	PrivateKey string `json:"privateKey"`
}

func (s *SftpAuthenticationInfoSshKey) AsSshKey() *SftpAuthenticationInfoSshKey {
	return s
}

////////////////////////////////////////////////////////////
// StorageAzureBlob
////////////////////////////////////////////////////////////

// Depository backing storage on the basis of Azure Blob Storage
type StorageAzureBlob struct {
	*StorageBase
	StorageAccount     string                   `json:"storageAccount"`
	ContainerName      string                   `json:"containerName"`
	BlobNamePrefix     string                   `json:"blobNamePrefix"`
	AuthenticationInfo IAzureAuthenticationInfo `json:"authenticationInfo"`
}

func (s *StorageAzureBlob) AsAzureBlobStorage() *StorageAzureBlob {
	return s
}

func (s *StorageAzureBlob) UnmarshalJSON(data []byte) error {
	type Alias StorageAzureBlob
	aux := &struct {
		AuthRaw json.RawMessage `json:"authenticationInfo"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	// Unmarshal to extract raw "authenticationInfo" block
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Peek at the "type" field
	var typePeek struct {
		Type AzureAuthenticationType `json:"type"`
	}
	if err := json.Unmarshal(aux.AuthRaw, &typePeek); err != nil {
		return err
	}
	// Create the appropriate storage type based on the type
	var authInfo IAzureAuthenticationInfo
	switch typePeek.Type {
	case AZURE_AUTHENTICATION_TYPE_BASIC:
		authInfo = &AzureBasicAuthenticationInfo{}
	case AZURE_AUTHENTICATION_TYPE_SAS:
		authInfo = &AzureSasAuthenticationInfo{}
	case AZURE_AUTHENTICATION_TYPE_SHARED_KEY:
		authInfo = &AzureSharedKeyAuthenticationInfo{}
	default:
		return fmt.Errorf("unsupported authentication type: %s", typePeek.Type)
	}
	// Unmarshal the raw data into the storage type
	if err := json.Unmarshal(aux.AuthRaw, &authInfo); err != nil {
		return err
	}
	// Assign the storage to the slice
	s.AuthenticationInfo = authInfo
	return nil
}

type IAzureAuthenticationInfo interface {
	GetType() AzureAuthenticationType
	AsBasic() *AzureBasicAuthenticationInfo
	AsSas() *AzureSasAuthenticationInfo
	AsSharedKey() *AzureSharedKeyAuthenticationInfo
}

type AzureAuthenticationInfo struct {
	Type AzureAuthenticationType `json:"type"`
}

func (s *AzureAuthenticationInfo) GetType() AzureAuthenticationType {
	return s.Type
}

func (s *AzureAuthenticationInfo) AsBasic() *AzureBasicAuthenticationInfo {
	return nil
}

func (s *AzureAuthenticationInfo) AsSas() *AzureSasAuthenticationInfo {
	return nil
}

func (s *AzureAuthenticationInfo) AsSharedKey() *AzureSharedKeyAuthenticationInfo {
	return nil
}

type AzureBasicAuthenticationInfo struct {
	*AzureAuthenticationInfo
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *AzureBasicAuthenticationInfo) AsBasic() *AzureBasicAuthenticationInfo {
	return s
}

type AzureSasAuthenticationInfo struct {
	*AzureAuthenticationInfo
	Token string `json:"token"`
}

func (s *AzureSasAuthenticationInfo) AsSas() *AzureSasAuthenticationInfo {
	return s
}

type AzureSharedKeyAuthenticationInfo struct {
	*AzureAuthenticationInfo
	AccountName string `json:"accountName"`
	AccountKey  string `json:"accountKey"`
}

func (s *AzureSharedKeyAuthenticationInfo) AsSharedKey() *AzureSharedKeyAuthenticationInfo {
	return s
}
