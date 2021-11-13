package tests

import (
	"GoDatabaseAssignment/pairs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSavePairs(t *testing.T) {
	record := pairs.KeyValue{
		Key:   "test",
		Value: "test_key_value",
	}
	response, err := pairs.SavePair(record)
	assert.Nil(t, err)
	assert.Equal(t, record.Key, response.Key)

}

func TestGetKeyValuePairs(t *testing.T) {
	key := "test"
	response, err := pairs.GetKeyValuePair(key)
	assert.Nil(t, err)
	assert.Equal(t, "test_key_value", response.Value)
}
