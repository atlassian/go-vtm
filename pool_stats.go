package stingray

import (
	"encoding/json"
	"net/http"
)

// A PoolStats is a Stingray pool.
type PoolStats struct {
	jsonStatsResource `json:"-"`
	PoolStatistics    `json:"statistics"`
}

type PoolStatistics struct {
	Algorithm          string `json:"algorithm"`
	BytesIn            int64  `json:"bytes_in"`
	BytesInHigh        uint32  `json:"bytes_in_hi"`
	BytesInLow         uint32  `json:"bytes_in_low"`
	BytesOut           int64  `json:"bytes_out"`
	BytesOutHigh       uint32  `json:"bytes_out_hi"`
	BytesOutLow        uint32  `json:"bytes_out_low"`
	ConnectionsQueued  int    `json:"conns_queued"`
	DisabledNodeCount  int    `json:"disabled"`
	DrainingNodeCount  int    `json:"draining"`
	MaxQueueTime       int    `json:"max_queue_time"`
	MeanQueueTime      int    `json:"mean_queue_time"`
	MinQueueTime       int    `json:"min_queue_time"`
	NodeCount          int    `json:"nodes"`
	SessionPersistence string `json:"persistence"`
	QueueTimeouts      int    `json:"queue_timeouts"`
	SessionsMigrated   int    `json:"session_migrated"`
	State              string `json:"state"`
	TotalConnections   int    `json:"total_conn"`
}

func (r *PoolStats) endpoint() string {
	return "pools"
}


//String will return back the json as a string
func (r *PoolStats) String() string {
	b := r.Bytes()
	return string(b)
}

//Bytes will return back just the bytes
func (r *PoolStats) Bytes() []byte {
	b, _ := jsonMarshal(r)
	return b
}

func (r *PoolStats) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewPoolStats(name string) *PoolStats {
	r := new(PoolStats)
	r.setName(name)
	return r
}

func (c *Client) GetPoolStats(name string) (*PoolStats, *http.Response, error) {
	r := NewPoolStats(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
