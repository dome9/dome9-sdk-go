package client

import (
	"bytes"
	"dome9"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//var client *dome9.Config


type Client struct {
	Config *dome9.Config
}

//func init()  {
//	client = dome9.DefaultConfig()
//}

// NewClient returns a new client for the specified apiKey.
func NewClient(config *dome9.Config) (c *Client) {
	if config == nil {
		config = dome9.DefaultConfig()
	}
	c = &Client{Config: config}
	//c.IpLists = &services.IpListService{client: c}
	return
}

func (client *Client) Test(){
	fmt.Println("Success")
}

func (client *Client) NewRequestDo(method, url string, body, v interface{}) (*http.Response, error) {
	// E.g. "resp, err := client.NewRequestDo("POST", "iplist/", ipList, &v)"
	// We may add query options handling in var o
	req, err := client.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	client.logRequest(req)
	return client.do(req, v)
}

// Generating the Http request
func (client *Client) newRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Join the path to the base-url
	u := *client.Config.BaseURL
	unescaped, err := url.PathUnescape(urlPath)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = client.Config.BaseURL.Path + urlPath
	u.Path = client.Config.BaseURL.Path + unescaped

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.Config.AccessID, client.Config.SecretKey)

	req.Header.Add("Accept", "application/json")
	// TODO maybe add only if body isn't nil
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (client *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := client.Config.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if err := checkErrorInResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		if err := decodeJSON(resp, v); err != nil {
			return resp, err
		}
	}
	client.logResponse(resp)

	return resp, nil
}

func decodeJSON(res *http.Response, v interface{}) error {
	return json.NewDecoder(res.Body).Decode(v)
}

func checkErrorInResponse(res *http.Response) error {
	if c := res.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: res}
	errorMessage, err := ioutil.ReadAll(res.Body)
	if err == nil && len(errorMessage) > 0 {
		errorResponse.Message = string(errorMessage)
	}
	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response
	Message  string
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("FAILED: %v, %v, %d, %v, %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Response.Status, r.Message)
}

func (client *Client) logf(format string, args ...interface{}) {
	if client.Config.Logger != nil {
		client.Config.Logger.Printf(format, args...)
	}
}

const logReqMsg = `Request "%s %s" details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

func (client *Client) logRequest(req *http.Request) {
	if client.Config.Logger != nil && req != nil {
		out, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			client.logf(logReqMsg, req.Method, req.URL, string(out))
		}
	}
}

const logRespMsg = `Response "%s %s" details:
---[ RESPONSE ]----------------------------------------
%s
-------------------------------------------------------`

func (client *Client) logResponse(resp *http.Response) {
	if client.Config.Logger != nil && resp != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			client.logf(logRespMsg, resp.Request.Method, resp.Request.URL, string(out))
		}
	}
}
