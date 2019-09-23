package cloudaccounts

const (
	RESTfulPathAWS   = "cloudaccounts"
	RESTfulPathAzure = "AzureCloudAccount"
	RESTfulPathGCP   = "GoogleCloudAccount"
)

// AWS service paths
const (
	RESTfulServicePathAWSName            = "name"
	RESTfulServicePathAWSRegionConfig    = "region-conf"
	RESTfulServicePathOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathCredentials        = "credentials"
)

type QueryParameters struct {
	ID string
}
