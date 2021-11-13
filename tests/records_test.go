package tests

import (
	"GoDatabaseAssignment/records"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFromDb(t *testing.T) {
	record := records.RequestRecord{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	response := records.GetRecords(record)
	assert.Equal(t, response.Code, records.SuccessCode)
	record.StartDate = "2016-13-26"
	response = records.GetRecords(record)
	assert.Equal(t, response.Code, records.DateParseCode)
}
