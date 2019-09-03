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
)

// Config contains all the configuration data for the API Client
type Config struct {
	BaseURL    *url.URL
	HTTPClient *http.Client

	// The logger writer interface to write logging messages to. Defaults to standard out.
	Logger *log.Logger

	// Credentials for basic authentication.
	AccessID, SecretKey string
}

// DefaultConfig returns a default configuration for the client.
func DefaultConfig() *Config {
	accessID, secretKey := DefaultKeys()
	return &Config{
		BaseURL:    DefaultBaseURL(),
		HTTPClient: DefaultHTTPClient(),
		Logger:     DefaultLogger(), // TODO default should be nil and should add a setter for logger
		AccessID:   accessID,
		SecretKey:  secretKey,
	}
}

func DefaultBaseURL() *url.URL {
	baseURL, _ := url.Parse(defaultBaseURL)
	return baseURL
}

func DefaultHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultTimeout}
}

func DefaultLogger() *log.Logger {
	return log.New(os.Stdout, "D9-logger: ", log.LstdFlags|log.Lshortfile)
}

func (c *Config) SetBaseURL(rawUrl string) (err error) {
	baseURL, err := url.Parse(rawUrl)
	c.BaseURL = baseURL
	return
}

func (c *Config) SetKeys(accessID, secretKey string) {
	c.AccessID = accessID
	c.SecretKey = secretKey
}

func DefaultKeys() (accessID string, secretKey string) {
	accessID = os.Getenv("accessID")
	secretKey = os.Getenv("secretKey")
	return
}
