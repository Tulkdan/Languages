GET http://localhost:8000/pessoas
HTTP 400
[Asserts]
body contains "Missing query param 't' in request"

GET http://localhost:8000/pessoas
[QueryStringParams]
t: invalid
HTTP 200
[Asserts]
header   "Content-Type" == "application/json"
jsonpath "$" count == 0

GET http://localhost:8000/pessoas
[QueryStringParams]
t: berto
HTTP 200
[Asserts]
header   "Content-Type" == "application/json"
jsonpath "$" count == 1
jsonpath "$[0].nickname" == "josé"

GET http://localhost:8000/pessoas
[QueryStringParams]
t: node
HTTP 200
[Asserts]
header   "Content-Type" == "application/json"
jsonpath "$" count == 1
jsonpath "$[0].nickname" == "josé"

