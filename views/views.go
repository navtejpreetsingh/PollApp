package views

import (
	"PollApp/models"
	"database/sql"
	"fmt"
)

// returns the list of unique qids that are present in the database
func getUniqueQuestions(DB *sql.DB) []int {
	rows, DB_error := DB.Query("SELECT DISTINCT(qid) FROM questions_table WHERE questions_table.enabled is true ORDER BY qid;")
	if DB_error != nil {
		panic(DB_error)
	}
	defer rows.Close()
	unique_qids := []int{}

	for rows.Next() {
		var temp_qid int
		DB_error := rows.Scan(&temp_qid)
		unique_qids = append(unique_qids, temp_qid)
		if DB_error != nil {
			panic(DB_error)
		}
	}
	DB_error = rows.Err()
	if DB_error != nil {
		panic(DB_error)
	}
	return unique_qids
}

// returns the list of questions along with its options
func GetPoll(DB *sql.DB) []*models.PollQuestion {
	poll := []*models.PollQuestion{}
	question_ids := getUniqueQuestions(DB)
	fmt.Println("unique qids: ", question_ids)
	for _, question_id := range question_ids {
		// query one question at a time
		query := fmt.Sprintf("SELECT q.qid, q.question, o.option_id, o.option, o.votes FROM questions_table as q INNER JOIN options_table as o ON q.qid = o.qid WHERE q.qid = %v AND q.qid = o.qid ORDER BY o.option_id;", question_id)
		rows, DB_error := DB.Query(query)
		if DB_error != nil {
			panic(DB_error)
		}
		// defer rows.Close()
		questionRow := models.PollQuestion{Qid: question_id, Question: "DEFAULT QUESTION"}
		optionsList := []models.PollOption{}
		for rows.Next() {
			tempOption := models.PollOption{}
			DB_error := rows.Scan(&questionRow.Qid, &questionRow.Question, &tempOption.Option_id, &tempOption.Option, &tempOption.Votes)
			tempOption.Qid = question_id
			if DB_error != nil {
				panic(DB_error)
			}
			// fmt.Println("\n", questionRow.qid, questionRow.question, tempOption)
			optionsList = append(optionsList, tempOption)
		}

		questionRow.Options = optionsList
		poll = append(poll, &questionRow)

		DB_error = rows.Err()
		if DB_error != nil {
			panic(DB_error)
		}
		rows.Close()
	}

	return poll
}

// resgisters response from anonymous user
func RegisterVote(DB *sql.DB, qid int, option_id int) {
	query := fmt.Sprintf("UPDATE options_table SET votes = votes + 1 WHERE option_id = %v AND qid = %v; ", option_id, qid)
	res, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("votes updated\n")
}

// adds a new question
func AddQuestion(DB *sql.DB, question string, options []string) int {

	addQuestionQuery := fmt.Sprintf("INSERT INTO questions_table(question, enabled) VALUES('%v',true);", question)
	res, err := DB.Exec(addQuestionQuery)
	if err != nil {
		panic(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	var qid int

	getQidQuery := fmt.Sprintf("SELECT qid FROM questions_table WHERE question = '%v';", question)

	if err := DB.QueryRow(getQidQuery).Scan(&qid); err != nil {
		fmt.Println("executed")
		if err == sql.ErrNoRows {
			fmt.Printf("Unknown question")
		} else {
			fmt.Errorf("SQL error: %v", err)
		}
		return -1
	}

	for option_id, option := range options {
		res, err := DB.Exec(fmt.Sprintf("insert into options_table values(%v,%v,'%v',0);", qid, option_id, option))
		if err != nil {
			panic(err)
		}
		_, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}
	}
	return qid
}

// deletes a question
func DeleteQuestion(DB *sql.DB, qid int) {
	deleteFromOptionsTable := fmt.Sprintf("DELETE FROM options_table WHERE qid=%v", qid)
	res, err := DB.Exec(deleteFromOptionsTable)
	if err != nil {
		panic(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	deleteFromQuestionsTable := fmt.Sprintf("DELETE FROM questions_table WHERE qid=%v", qid)
	res, err = DB.Exec(deleteFromQuestionsTable)
	if err != nil {
		panic(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
}
