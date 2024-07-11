
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
          "id": 50,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
     }'

curl -X PUT http://localhost:8080/users/1/favorites/1 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 50,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
     }'


# Tests

go test -v ./.../tests -cover

# Benchmark

go test -bench=. -benchtime=5s ./.../tests