package contracts

import "time"

type Item struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type Batch struct {
	Items []Item `json:"items"`
}

type Limits struct {
	MaxItems int           `json:"maxItems"`
	Interval time.Duration `json:"interval"`
}

type ProcessRequest struct {
	Batch Batch `json:"batch"`
}

type HTTPError struct {
	Message    string `json:"message"`
	IncidentID string `json:"incident_id,omitempty"`
}
