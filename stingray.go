package stingray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"errors"
)

const (
	baseConfigPath = "/api/tm/3.5/config/active/"
	baseStatsPath = "/api/tm/3.5/status/local_tm/statistics/"
)

type pathType int
const (
	configPath pathType = iota
	statsPath
)

// A Client manages communication with the Stingray API.
type Client struct {
	// HTTP client used to communicate with the API.
	Client    *http.Client

	// API base URL for configuration namespace
	configURL *url.URL

	// API base URL for stats namespace
	statsURL  *url.URL

	// Username used for communicating with the API.
	Username  string

	// Password used for communicating with the API.
	Password  string
}

// NewClient returns a new Stingray API client, using the supplied
// URL, username, and password
func NewClient(httpClient *http.Client, urlStr, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	cu, _ := url.Parse(urlStr)
	su := cu
	var relStats *url.URL
	if cu.Path == "" {
		rel, _ := url.Parse(baseConfigPath)
		cu = cu.ResolveReference(rel)
		relStats, _ = url.Parse(baseStatsPath)
	} else {
		statsRelPath := "../../status/local_tm/statistics/"
		relStats, _ = url.Parse(statsRelPath)
	}
	su = su.ResolveReference(relStats)

	c := &Client{
		Client:   httpClient,
		configURL:  cu,
		statsURL:   su,
		Username: username,
		Password: password,
	}

	return c
}

// NewRequest creates a new request with the params
func (c *Client) NewRequest(method, urlStr string, body *[]byte) (*http.Request, error) {
	return c.doMakeRequest(configPath, method, urlStr, body)
}

func (c *Client) doMakeRequest(pathType pathType, method, urlStr string, body *[]byte) (*http.Request, error) {
	var bodyreader io.Reader
	var baseUrl *url.URL

	if body != nil {
		bodyreader = bytes.NewReader(*body)
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	switch pathType {
	case configPath:
		baseUrl = c.configURL
		break
	case statsPath:
		baseUrl = c.statsURL
		break
	default:
		return nil, errors.New("Tried to make request with an unknown path type")
	}

	u := baseUrl.ResolveReference(rel)
	req, err := http.NewRequest(method, u.String(), bodyreader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	return req, nil
}

// Do sends an API request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	return resp, nil
}

// Get retrieves a resource
func (c *Client) Get(r Resourcer) (*http.Response, error) {

	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	req, err := c.doMakeRequest(r.pathType(), "GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return resp, err
	}

	err = r.decode(body)

	return resp, err
}

// Set sets a resource
func (c *Client) Set(r Resourcer) (*http.Response, error) {
	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	data := r.Bytes()
	req, err := c.NewRequest("PUT", u, &data)

	req.Header.Add("Content-Type", r.contentType())
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Delete deletes a resource
func (c *Client) Delete(r Resourcer) (*http.Response, error) {
	u := fmt.Sprintf("%v/%v", r.endpoint(), r.Name())

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	return resp, err
}

// List lists resources of the specified type
func (c *Client) List(r Resourcer) ([]string, *http.Response, error) {
	req, err := c.NewRequest("GET", r.endpoint(), nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp, err
	}

	rl := &resourceList{}
	err = rl.decode(body)
	if err != nil {
		return nil, resp, err
	}

	return rl.names(), resp, nil
}

// ErrorResponse represents an error message returned by the Stingray API.
//
// See Chapter 2, Further Aspects of the Resource Model, Errors.
type ErrorResponse struct {
	Response  *http.Response // HTTP response that caused this error
	ID        string         `json:"error_id"`
	Text      string         `json:"error_text"`
	ErrorInfo interface{}    `json:"error_info"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ID, r.Text, r.ErrorInfo)
}

// CheckResponse checks the API response for errors, and returns them
// if present. A response is considered an error if it has a status
// code outside the 200 range. API error responses are expected to
// have either no response body, or a JSON response body that maps to
// ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

// jsonMarshal un-escapes certain "\uXXXX" escape sequences since the
// Stingray REST API does not decode these correctly. The
// json.Unmarshal function creates these escape sequences for &, <,
// and >.
func jsonMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)

	return b, err
}
