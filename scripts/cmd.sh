
# GET
curl -X GET http://localhost:8080/users/1/favorites

# POST
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d @json/insight.json

# DELETE
curl -X DELETE http://localhost:8080/users/3/favorites/2

# PUT
curl -X PUT http://localhost:8080/users/3/favorites/66 \
     -H "Content-Type: application/json" \
     -d @json/insight.json 