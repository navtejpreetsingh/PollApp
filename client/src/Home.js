import { Button } from '@mui/base';
import QuestionList from './QuestionList';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from "axios";


const API_BASE_URL = "/";
const VIEW_POLL_URL = "view_poll"

const Home = () => {
    const [ questions, setQuestions ] =  useState(null);
    let navigate = useNavigate();

    useEffect(() => {
    console.log("Use effect ran!");
    console.log(questions);
    axios.get(API_BASE_URL+VIEW_POLL_URL)
      .then(res => {
        console.log(res);
        if (res.status != 200){
          throw Error('Could not fetch the requested data.');
        }
        console.log("dataFromServer => ", res.data);
        // sorting and storing the data
        res.data.sort((a,b)=> (a.qid - b.qid));
        res.data.map( (ques) => (ques.options.sort((a,b) => (a.option_id - b.option_id))));
        
        setQuestions(res.data);
      })
      .catch(err => {
        console.log(err.message);
      })
  },[]);

    const handleParticipateNow = () => {
      navigate("/vote");
    };

    return (
    <div className='action-section'>
        {questions && <QuestionList className="QuestionList" voting={"false"} questions={questions} title="Here's the questionaire of the Poll"/>}
        {/* <Link to="/vote" elem> Participate Now </Link> */}
        <Button className='primary' onClick={handleParticipateNow}>
          {/* <a href="/vote"> */}
            Participate Now
          {/* </a>*/}
          </Button> 
    </div>
    );
};

export default Home;