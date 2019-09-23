package cloudaccounts

const (
	RESTfulPathAWS   = "cloudaccounts"
	RESTfulPathAzure = "AzureCloudAccount"
	RESTfulPathGCP   = "GoogleCloudAccount"
)

// AWS service paths
const (
	RESTfulServicePathAWSName               = "name"
	RESTfulServicePathAWSRegionConfig       = "region-conf"
	RESTfulServicePathAWSOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathAWSCredentials        = "credentials"
)

// Azure service paths
const (
	RESTfulServicePathAzureName               = "AccountName"
	RESTfulServicePathAzureOperationMode      = "OperationMode"
	RESTfulServicePathAzureOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathAzureCredentials        = "credentials"
)

type QueryParameters struct {
	ID string
}
