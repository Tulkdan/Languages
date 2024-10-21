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

GET http://localhost:8000/contagem-pessoas
HTTP 200
1
