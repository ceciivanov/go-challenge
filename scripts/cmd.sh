
curl -X GET http://localhost:8080/users/1/favorites

curl -X DELETE http://localhost:8080/users/1/favorites/1

curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d @json/insight.json

curl -X PUT http://localhost:8080/users/1/favorites/1 \
     -H "Content-Type: application/json" \
     -d @json/insight.json 

curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 100,
          "type": "Audience",
          "description": "This audience is a 40 year old",
          "age": 40,
          "ageGroup": "25-45",
          "gender": "Male",
          "birthCountry": "USA",
          "hoursSpentOnMedia": 4,
          "numberOfPurchases": 10
     }'

curl -X PUT http://localhost:8080/users/1/favorites/1 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 1,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
     }'


# Tests

go test -v ./... -cover

# Benchmark

go test -bench=. -benchtime=5s ./.../tests