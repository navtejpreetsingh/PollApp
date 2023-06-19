curl http://localhost:8080/view_poll

curl -X POST -H "Content-Type: application/json" -d '{"poll_votes": [{"qid":1, "option_id":2},{"qid":2, "option_id":4}]}' http://localhost:8080/participate

curl -X POST -H "Content-Type: application/json" -d '{"qid":1, "option_id":2}' http://localhost:8080/participate