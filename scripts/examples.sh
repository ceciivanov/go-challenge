# GET valid user favorites
curl -X GET http://localhost:8080/users/1/favorites

# GET invalid user favorites
curl -X GET http://localhost:8080/users/999999/favorites

# DELETE valid user favorite
curl -X DELETE http://localhost:8080/users/1/favorites/1

# DELETE invalid user favorite
curl -X DELETE http://localhost:8080/users/1/favorites/999999

# ADD valid user favorite
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

# ADD existing user favorite
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 1,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
        }'

# ADD invalid user favorite
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 200,
          "type": "INVALIDTYPE",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
          }'

# EDIT valid user favorite
curl -X PUT http://localhost:8080/users/1/favorites/1 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 1,
          "type": "Insight",
          "description": "Updated Insight for testing",
          "text": "Updated Insight"
     }'

# EDIT user favorite with mismatched id
curl -X PUT http://localhost:8080/users/1/favorites/10 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 1,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
     }'