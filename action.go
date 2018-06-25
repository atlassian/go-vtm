package stingray

import (
	"encoding/json"
	"net/http"
)

type Action struct {
	jsonConfigResource     `json:"-"`
	ActionProperties `json:"properties"`
}

type ProgramArgument struct {
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ActionProperties struct {
	Basic struct {
		Note              *string `json:"note,omitempty"`
		SyslogMsgLenLimit *int    `json:"syslog_msg_len_limit,omitempty"`
		Timeout           *int    `json:"timeout,omitempty"`
		Type              *string `json:"type,omitempty"`
		Verbose           *bool   `json:"verbose,omitempty"`
	} `json:"basic"`
	Email struct {
		Server *string   `json:"server,omitempty"`
		To     *[]string `json:"to,omitempty"`
	} `json:"email"`
	Log struct {
		File *string `json:"file,omitempty"`
		From *string `json:"from,omitempty"`
	} `json:"log"`
	Program struct {
		Arguments []ProgramArgument `json:"arguments,omitempty"`
		Program   *string           `json:"program,omitempty"`
	} `json:"program"`
	SOAP struct {
		AdditionalData *string `json:"additional_data,omitempty"`
		Password       *string `json:"password,omitempty"`
		Proxy          *string `json:"proxy,omitempty"`
		Username       *string `json:"username,omitempty"`
	} `json:"soap"`
	Syslog struct {
		SyslogHost *string `json:"sysloghost,omitempty"`
	} `json:"syslog"`
	Trap struct {
		AuthPassword  *string `json:"auth_password,omitempty"`
		Community     *string `json:"community,omitempty"`
		HashAlgorithm *string `json:"hash_algorithm,omitempty"`
		PrivPassword  *string `json:"priv_password,omitempty"`
		TrapHost      *string `json:"traphost,omitempty"`
		Username      *string `json:"username,omitempty"`
		Version       *string `json:"version,omitempty"`
	} `json:"trap"`
}

func (r *Action) endpoint() string {
	return "actions"
}

//String will return back the json as a string
func (r *Action) String() string {
	b := r.Bytes()
	return string(b)
}

//Bytes will return back just the bytes
func (r *Action) Bytes() []byte {
	b, _ := jsonMarshal(r)
	return b
}

func (r *Action) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewAction(name string) *Action {
	r := new(Action)
	r.setName(name)
	return r
}

func (c *Client) GetAction(name string) (*Action, *http.Response, error) {
	r := NewAction(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListActions() ([]string, *http.Response, error) {
	return c.List(&Action{})
}
