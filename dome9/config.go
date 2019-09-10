package dome9

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	defaultBaseURL = "https://api.dome9.com/v2/"
	defaultTimeout = 20 * time.Second
	loggerPrefix   = "Dome9-logger: "
)

// Config contains all the configuration data for the API client
type Config struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	// The logger writer interface to write logging messages to. Defaults to standard out.
	Logger *log.Logger
	// Credentials for basic authentication.
	AccessID, SecretKey string
}

/*
NewConfig returns a default configuration for the client.
By default it will try to read the access and te secret from the environment variables.
*/

// TODO Add healthCheck method to NewConfig
func NewConfig(accessID, secretKey, rawUrl string) (*Config, error) {
	if accessID == "" || secretKey == "" {
		accessID = os.Getenv("DOME9_ACCESS_ID")
		secretKey = os.Getenv("DOME9_SECRET_KEY")
	}
	if rawUrl == "" {
		rawUrl = defaultBaseURL
	}

	baseURL, err := url.Parse(rawUrl)
	return &Config{
		BaseURL:    baseURL,
		HTTPClient: getDefaultHTTPClient(),
		Logger:     getDefaultLogger(), // TODO default should be nil and should add a setter for logger
		AccessID:   accessID,
		SecretKey:  secretKey,
	}, err
}

func getDefaultHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultTimeout}
}

func getDefaultLogger() *log.Logger {
	return log.New(os.Stdout, loggerPrefix, log.LstdFlags|log.Lshortfile)
}