POST http://localhost:8000/pessoas
Content-Type: application/json
{
	"name": "Ana Barbosa",
	"nickname": "ana",
	"birthdate": "1985-09-23T00:00:00Z",
	"stack": []
}
HTTP 200
[Captures]
person_id: jsonpath "$.id"


GET http://localhost:8000/pessoas/{{person_id}}
HTTP 200
[Asserts]
header   "Content-Type" == "application/json"
jsonpath "$.name" == "Ana Barbosa"
jsonpath "$.nickname" == "ana"
jsonpath "$.birthdate" == "1985-09-23T00:00:00Z"
jsonpath "$.stack" count == 0

POST http://localhost:8000/pessoas
Content-Type: application/json
{
	"nickname" : "josé",
	"name" : "José Roberto",
	"birthdate": "2000-10-01T00:00:00Z",
	"stack": ["C#", "Node", "Oracle"]
}
HTTP 200
[Captures]
person_id: jsonpath "$.id"


GET http://localhost:8000/pessoas/{{person_id}}
HTTP 200
[Asserts]
header   "Content-Type" == "application/json"
jsonpath "$.name" == "José Roberto"
jsonpath "$.nickname" == "josé"
jsonpath "$.birthdate" == "2000-10-01T00:00:00Z"
jsonpath "$.stack" count == 3

