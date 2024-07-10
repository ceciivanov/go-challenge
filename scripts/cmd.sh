
curl -X GET http://localhost:8080/users/1/favorites

curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d @json/insight.json

curl -X DELETE http://localhost:8080/users/3/favorites/2

curl -X PUT http://localhost:8080/users/1/favorites/1 \
     -H "Content-Type: application/json" \
     -d @json/insight.json 