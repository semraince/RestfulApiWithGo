package pairs

import (
	"GoDatabaseAssignment/inmemory"
	"errors"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetKeyValuePair(key string) (KeyValue, error) {
	var pairResponse KeyValue
	result, err := inmemory.Read(key)
	if err != nil {
		return pairResponse, errors.New("An Error Occured")
	}
	if result == nil {
		return pairResponse, nil
	}
	pairResponse = KeyValue{
		Key:   result.Key,
		Value: result.Value,
	}
	return pairResponse, err
}

func savePair(data KeyValue) error {
	pair := inmemory.KeyValue{
		Key:   data.Key,
		Value: data.Value,
	}
	err := inmemory.Write(&pair)
	return err
}