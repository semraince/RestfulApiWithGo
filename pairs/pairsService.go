package pairs

import (
	"GoDatabaseAssignment/inmemory"
	"errors"
)

const (
	DBERROR     = "An Error Occured!"
	KEYNOTFOUND = "Key Not Found"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetKeyValuePair(key string) (KeyValue, error) {
	var pairResponse KeyValue
	result, err := inmemory.Read(key)
	if err != nil {
		return pairResponse, errors.New(DBERROR)
	}
	if result == nil {
		return pairResponse, errors.New(KEYNOTFOUND)
	}
	pairResponse = KeyValue{
		Key:   result.Key,
		Value: result.Value,
	}
	return pairResponse, nil
}

func SavePair(data KeyValue) (KeyValue, error) {
	pair := inmemory.KeyValue{
		Key:   data.Key,
		Value: data.Value,
	}
	err := inmemory.Write(&pair)
	if err != nil {
		return data, err
	}
	return data, nil
}
