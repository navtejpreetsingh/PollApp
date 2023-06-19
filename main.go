package main

import (
	"PollApp/models"
	"PollApp/views"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var DB_error error

type httpResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func viewPollHandler(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	if r.URL.Path != "/view_poll" {
		http.Error(w, "{'code': '' }", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { //we only support GET request
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	pollQuestions := views.GetPoll(DB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pollQuestions)
}

func participateHandler(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	if r.URL.Path != "/participate" { // https://pkg.go.dev/net/http#Error
		http.Error(w, "CODE 404 not found!!", http.StatusNotFound)
		return
	}
	if r.Method != "POST" { // https://pkg.go.dev/net/http#Error
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	var pollResponses models.PollParticipation
	err := json.NewDecoder(r.Body).Decode(&pollResponses)

	if err != nil {
		log.Println("Error decoding JSON:", err)
		// fmt.Fprintf(w, "ERROR: %v", err.Error)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("\n\tpassed %v", pollResponses)

	for _, pr := range pollResponses.PollVotes {
		fmt.Printf("for qid %v voted %v", pr.Qid, pr.Option_id)
		views.RegisterVote(DB, pr.Qid, pr.Option_id)
	}

	// fmt.Printf("for qid %v voted %v", pollResponses.Qid, pollResponses.Option_id)

	w.Header().Set("Content-Type", "application/json")
	response := httpResponse{Status: "created", Code: 201}
	// fmt.Println("response participateHandler: ", response)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func addQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" { // https://pkg.go.dev/net/http#Error
		http.Error(w, "CODE 404 not found!!", http.StatusNotFound)
		return
	}
	if r.Method != "POST" { // https://pkg.go.dev/net/http#Error
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var pollQuestion models.PollQuestion
	err := json.NewDecoder(r.Body).Decode(&pollQuestion)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		// fmt.Fprintf(w, "ERROR: %v", err.Error)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println("inside the handler :: ", pollQuestion)
	question := pollQuestion.Question
	options := []string{}
	for _, opt := range pollQuestion.Options {
		options = append(options, opt.Option)
	}

	qid := views.AddQuestion(DB, question, options)
	response := httpResponse{Status: strconv.Itoa(qid), Code: 201}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func deleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the question ID from the URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}
	qid, _ := strconv.Atoi(parts[2])

	// Assuming you have a function to delete the question by ID
	views.DeleteQuestion(DB, qid)

	response := httpResponse{Status: "Question deleted successfully", Code: http.StatusOK}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, origin, Cache-Control, X-Requested-With")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
// }

func ConnectDB() (*sql.DB, error) {
	connectionStr := "user=postgres dbname=poll password=12345 host=localhost sslmode=disable"
	DB, DB_error := sql.Open("postgres", connectionStr)
	if DB_error != nil {
		panic(DB_error)
	}
	DB_error = DB.Ping()
	if DB_error != nil {
		panic(DB_error)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return DB, DB_error
}

func main() {

	// connecting database
	DB, DB_error = ConnectDB()

	// routes
	http.HandleFunc("/view_poll", viewPollHandler)
	http.HandleFunc("/participate", participateHandler)
	http.HandleFunc("/add", addQuestionHandler)
	http.HandleFunc("/delete/", deleteQuestionHandler)

	fmt.Println("Start the server on port 8080")

	// start server
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
