# REST API Endpoints

1. records(POST) fetch data from mongodb
2. in-memory(POST) create key value pair in-memory
3. in-memory/:key (GET) key valeu pair from in-memory

# Set-Up

1. git clone https://github.com/semraince/RestfulApiWithGo.git
2. cd RestfulApiWithGo
3. go run main.go

Default port is 9090, it can be changed from conf.yaml

# Heroku End-Points

Postman request examples can be found in GoDatabaseAssignment.postman_collection.json file

### MongoDb

The request payload of the first endpoint will include a JSON with 4 fields.
“startDate” and “endDate” fields will contain the date in a “YYYY-MM-DD” format.
Data is filtered by using “createdAt”
“minCount” and “maxCount” are for filtering the data. Sum of the “count” array in
the documents should be between “minCount” and “maxCount”.

```jsx
https://sleepy-hollows-06250.herokuapp.com/records
```

```jsx
{
  "startDate": "2016-01-26",
  "endDate": "2018-02-02",
  "minCount": 2700,
  "maxCount": 3000
}
```

##### Response Payload

```jsx
{
  "code":0,
  "msg":"Success",
  "records":[
    {
    "key":"TAKwGc6Jr4i8Z487",
    "createdAt":"2017-01-28T01:22:14.398Z",
    "totalCount":2800
    },
    {
    "key":"NAeQ8eX7e5TEg7oH",
    "createdAt":"2017-01-27T08:19:14.135Z",
    "totalCount":2900
    }
  ]
}
```

##### Error Responses

| Status | Response                           |
| ------ | ---------------------------------- |
| 502    | `{ "msg": "Db Error Has Occured"}` |
| 503    | `{ "message": "Date Parse Error"}` |

### InMemory Db

##### POST Request

```jsx
https://sleepy-hollows-06250.herokuapp.com/in-memory
```

The request payload of POST endpoint will include a JSON with 2 fields.
“key” fields holds the key (any key in string type)
“value” fields holds the value (any value in string type)

##### Request Payload

```jsx
{
"key": "active-tabs",
"value": "getir"
}
```

Response payload return echo of the request or if any error occured

##### GET Request

```jsx
https://sleepy-hollows-06250.herokuapp.com/in-memory?key=active-tabs
```

The request payload of GET endpoint will include 1 query parameter. That is “key” param holds the key (any key in string type)

Response payload of GET endpoint should return a JSON with 2 fields or error in message field.
