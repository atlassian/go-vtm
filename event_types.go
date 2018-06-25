package stingray

import (
	"encoding/json"
	"net/http"
)

type EventType struct {
	jsonConfigResource        `json:"-"`
	EventTypeProperties `json:"properties"`
}

type EventObjects struct {
	Objects *[]string `json:"objects,omitempty"`
}

type EventTag struct {
	EventTags *[]string `json:"event_tags,omitempty"`
}

type EventTagObjects struct {
	EventTag
	EventObjects
}

type EventTypeProperties struct {
	Basic struct {
		BuiltIn *bool     `json:"built_in,omitempty"`
		Note    *string   `json:"note,omitempty"`
		Actions *[]string `json:"actions,omitempty"`
	} `json:"basic"`
	CloudCredentials EventTagObjects `json:"cloudcredentials"`
	Config           EventTag        `json:"config"`
	FaultTolerance   EventTag        `json:"faulttolerance"`
	General          EventTag        `json:"general"`
	GLB              EventTagObjects `json:"glb"`
	Java             EventTag        `json:"java"`
	LicenseKeys      EventTagObjects `json:"licensekeys"`
	Locations        EventTagObjects `json:"locations"`
	Monitors         EventTagObjects `json:"monitors"`
	Pools            EventTagObjects `json:"pools"`
	Protection       EventTagObjects `json:"protection"`
	Rules            EventTagObjects `json:"rules"`
	Slm              EventTagObjects `json:"slm"`
	SSL              EventTag        `json:"ssl"`
	SSLhw            EventTag        `json:"sslhw"`
	TrafficScript    EventTag        `json:"trafficscript"`
	Vservers         EventTagObjects `json:"vservers"`
	Zxtms            EventTagObjects `json:"zxtms"`
}

func (r *EventType) endpoint() string {
	return "event_types"
}

//String will return back the json as a string
func (r *EventType) String() string {
	b := r.Bytes()
	return string(b)
}

//Bytes will return back just the bytes
func (r *EventType) Bytes() []byte {
	b, _ := jsonMarshal(r)
	return b
}

func (r *EventType) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewEventType(name string) *EventType {
	r := new(EventType)
	r.setName(name)
	return r
}

func (c *Client) GetEventType(name string) (*EventType, *http.Response, error) {
	r := NewEventType(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListEventTypes() ([]string, *http.Response, error) {
	return c.List(&EventType{})
}
