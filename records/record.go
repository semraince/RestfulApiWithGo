package records

import "time"

type RequestRecord struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  uint   `json:"minCount"`
	MaxCount  uint   `json:"maxCount"`
}

type RecordDto struct {
	Key        string
	Value      string
	CreatedAt  time.Time
	Counts     []int
	TotalCount int
}

type Record struct {
	Key        string    `json:"key"`
	TotalCount int       `json:"totalCount"`
	CreatedAt  time.Time `json:"createdAt"`
}

type RecordResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
