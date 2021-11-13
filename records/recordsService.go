package records

import (
	"GoDatabaseAssignment/database"
	"GoDatabaseAssignment/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SuccessCode      = 0
	SuccessMessage   = "Success"
	FailureCode      = 502
	FailureMessage   = "Db Error Has Occured"
	DateParseCode    = 503
	DateParseMessage = "Date Parse Error"
)

func GetRecords(data RequestRecord) RecordResponse {
	var searchList []RecordDto
	var response RecordResponse
	var recordResults []*Record

	startDateParsed, err := utils.ParseDate(data.StartDate)
	if err != nil {
		return generateResponse(DateParseCode, DateParseMessage, recordResults)
	}
	endDateParsed, err := utils.ParseDate(data.EndDate)
	if err != nil {
		return generateResponse(DateParseCode, DateParseMessage, recordResults)
	}

	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"createdAt", bson.D{
					{"$gte", startDateParsed},
					{"$lte", endDateParsed},
				}},
			}},
		},
		{
			{"$project", bson.D{
				{"key", 1},
				{"createdAt", 1},
				{"_id", 0},
				{"totalCount", bson.D{{"$sum", "$counts"}}},
			}},
		},
		{
			{"$match", bson.D{
				{"totalCount", bson.D{
					{"$lte", data.MaxCount},
					{"$gte", data.MinCount},
				}},
			}},
		},
	}
	err = database.Aggregate("records", pipeline, &searchList)
	if err != nil {
		return generateResponse(FailureCode, FailureMessage, recordResults)
	}
	if searchList == nil {
		return response
	}
	for _, r := range searchList {
		recordResults = append(recordResults, &Record{
			Key:        r.Key,
			TotalCount: r.TotalCount,
			CreatedAt:  r.CreatedAt,
		})
	}
	return generateResponse(SuccessCode, SuccessMessage, recordResults)

}

func generateResponse(code int, message string, recordResults []*Record) RecordResponse {
	response := RecordResponse{
		Code: code,
		Msg:  message,
		Data: map[string]interface{}{"records": recordResults},
	}
	return response
}
