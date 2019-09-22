package cloudaccounts

const (
	RESTfulPathAWS   = "cloudaccounts"
	RESTfulPathAzure = "AzureCloudAccount"
	RESTfulPathGCP   = "GoogleCloudAccount"
)

// AWS service paths
const (
	RESTfulServicePathAWSName         = "name"
	RESTfulServicePathAWSRegionConfig = "region-conf"
)

type QueryParameters struct {
	ID string
}
