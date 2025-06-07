### Flashcard app
- This is a simple flashcard application mainly made to understand the API implementation on golang
- I have implemented multiple paths for different operations in this
- `POST /card` : will add a flashcard
- `GET /quiz` : will randmonly get a flashcard's question 


## How to run the application
1. go run main.go (and we will get a response like the application is running)
2. use different methods like POST and GET to view different results
3. before using the GET method, populate the map with different flashcards.

- Example usecase for adding a card:
```bash
curl -X POST \
-d {"Question":"Who is the Prime Minister of India","Answer":"Narendra Modi"} \
localhost:8080/card
```
- Example usecase for getting a random card:
```bash
curl localhost:8080/quiz
```