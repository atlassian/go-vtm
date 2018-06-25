package stingray

import (
	"encoding/json"
	"net/http"
)

type Protection struct {
	jsonConfigResource         `json:"-"`
	ProtectionProperties `json:"properties"`
}

type ProtectionProperties struct {
	Basic struct {
		Debug                     *bool   `json:"debug,omitempty"`
		Enabled                   *bool   `json:"enabled,omitempty"`
		LogTime                   *int    `json:"log_time,omitempty"`
		Note                      *string `json:"note,omitemtpy"`
		PerProcessConnectionCount *bool   `json:"per_process_connection_count,omitemtpy"`
		Rule                      *string `json:"rule,omitempty"`
		Testing                   *bool   `json:"testing,omitempty"`
	} `json:"basic"`
	ConnectionLimiting struct {
		Max10Connections  *int `json:"max_10_connections,omitempty"`
		Max1Connections   *int `json:"max_1_connections,omitempty"`
		MaxConnectionRate *int `json:"max_connection_rate,omitempty"`
		MinConnections    *int `json:"min_connections,omitempty"`
		RateTimer         *int `json:"rate_timer,omitempty"`
	} `json:"connection_limiting"`
	Http struct {
		CheckRfc2396     *bool `json:"check_rfc2396,omitempty"`
		MaxBodyLength    *int  `json:"max_body_length,omitempty"`
		MaxHeaderLength  *int  `json:"max_header_length,omitempty"`
		MaxRequestLength *int  `json:"max_request_length,omitempty"`
		MaxUrlLength     *int  `json:"max_url_length,omitempty"`
		RejectBinary     *bool `json:"reject_binary,omitempty"`
		SendErrorPage    *bool `json:"send_error_page,omitempty"`
	} `json:"http"`
	AccessRestriction struct {
		Allowed *[]string `json:"allowed,omitempty"`
		Banned  *[]string `json:"banned,omitempty"`
	} `json:"access_restriction"`
}

func (r *Protection) endpoint() string {
	return "protection"
}

//String will return back the json as a string
func (r *Protection) String() string {
	b := r.Bytes()
	return string(b)
}

//Bytes will return back just the bytes
func (r *Protection) Bytes() []byte {
	b, _ := jsonMarshal(r)
	return b
}

func (r *Protection) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewProtection(name string) *Protection {
	r := new(Protection)
	r.setName(name)
	return r
}

func (c *Client) GetProtection(name string) (*Protection, *http.Response, error) {
	r := NewProtection(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListProtections() ([]string, *http.Response, error) {
	return c.List(&Protection{})
}
