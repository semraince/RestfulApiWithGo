package records

import (
	"GoDatabaseAssignment/database"
	"GoDatabaseAssignment/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecords(data RequestRecord) RecordResponse {
	startDateParsed := utils.ParseDate(data.StartDate)
	endDateParsed := utils.ParseDate(data.EndDate)
	var searchList []RecordDto

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
	var response RecordResponse
	database.Aggregate("records", pipeline, &searchList)
	if searchList == nil {
		return response
	}
	var recordResults []*Record
	for _, r := range searchList {
		recordResults = append(recordResults, &Record{
			Key:        r.Key,
			TotalCount: r.TotalCount,
			CreatedAt:  r.CreatedAt,
		})
	}
	response = RecordResponse{
		Code: 0,
		Msg:  "success",
		Data: map[string]interface{}{"records": recordResults},
	}

	return response
}
