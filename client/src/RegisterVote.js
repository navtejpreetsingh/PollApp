import { useState, useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

const API_BASE_URL = "/";
const VIEW_POLL_URL = "view_poll"


const RegisterVote = () => {

    const [ option,setOption ] = useState(null);
    const [ isPending,setIsPending ] = useState(true);
    const [ questions, setQuestions ] =  useState(null);
    const [ currentQuestion, setCurrentQuestion ] = useState(0);
    let navigate = useNavigate();

    useEffect(() => {
    console.log("Use effect ran!");
    console.log(questions);
    fetch(API_BASE_URL+VIEW_POLL_URL)
      .then(res => {
        if (!res.ok){
          throw Error('Could not fetch the requested data.');
        }
        return res.json();
      })
      .then(data => {
        console.log("dataFromServer => ", data);
        // sorting and storing the data
        data.sort((a,b)=> (a.qid - b.qid));
        data.map( (ques) => (ques.options.sort((a,b) => (a.option_id - b.option_id))));
        console.log("data ???? ", data[0])
        setQuestions(data);
      })
      .catch(err => {
        console.log(err.message);
      })
    },[]);

    // effect to check whether the 
    useEffect(() => {
        if (option && Object.keys(option).length === questions.length){
            console.log("Length Matches!");
            setIsPending(false);
            // handleSubmit();
        } else if (option){
            console.log("Length does not match!", Object.keys(option).length);
        } else {
            console.log("Length does not match! , currLen = NULL");
        }
    },[option]);

    const convertData = (option) => {
        let arr = [];
        let keys = Object.keys(option).map((element) => (parseInt(element)));
        keys.forEach((qid) => {
            arr.push({"qid":qid,"option_id":parseInt(option[qid])})
        });
        return {"poll_votes" : arr}
    };

    const handleChangeOption = (e) => {
        let targetString = e.target.value;
        const myArray = targetString.split("_")
        setOption({ ...option, [myArray[0]]: myArray[1]});
        const nextQuestion = currentQuestion + 1;
		if (nextQuestion < questions.length) {
			setCurrentQuestion(nextQuestion);
		} else {
            setIsPending(false);
		}
    };

    const handleSubmit = async (e) => {
        var requestData = convertData(option);
        console.log("requestData here > ", requestData);

        let res = await axios({
            url: API_BASE_URL + "participate",
            method: 'POST',
            headers: {
                'Accept': 'application/json', 
                'content-type': 'application/json',
            },
            data: JSON.stringify(requestData)
        })
        .then((res) => {
            console.log(res)
            if (res.status === 201) {
                console.log("RESPONSE" ,res.data , " type:", typeof(res.data));
            } else {
                throw Error("Some error occured");
            }
        })
        .catch((err) => console.log("ERROR:", err));
        alert("You have successfully participated in the poll! Redirecting you to the home Page...")
        navigate("/");
    };

    const handleGoback = (e) => {
        navigate("/");
    };

    return (
        <div>
            {/* <div className="Question-List">
            <h4>Select the most appropriate option</h4>
                <FormControl>
                {questions && questions.map((ques) => (
                    <div className="Question" key={ques.qid}>
                        <h5>{ques.qid} .) {ques.question}</h5>
                        <RadioGroup onChange={handleChangeOption}>
                        {ques.options.map((opt) => (
                            <FormControlLabel value={opt.qid+"_"+opt.option_id} key={opt.qid+"_"+opt.option_id} control={<Radio/> } label={opt.option} />
                        ))}
                        </RadioGroup>
                    </div>
                ))}
                </FormControl>
            </div>
            
            <Button variant="contained" onClick={handleSubmit} disabled={isPending} style={{alignSelf:'right', margin: '25px'}}>Submit</Button> */}

                {questions && isPending && 
                    <div className='Question'>
                    <div className='question-section'>
                        <div className='question-count'>
                            <span>Question {currentQuestion + 1}</span>/{questions.length}
                        </div>
                        <div className='question-text'>{questions[currentQuestion].question}</div>
                    </div>
                    <div className='answer-section'>
                        {questions[currentQuestion].options.map((opt) => (
                            <button value={opt.qid+"_"+opt.option_id} key={opt.qid+"_"+opt.option_id} onClick={handleChangeOption}>{opt.option}</button>
                        ))}
                    </div>
                    </div>
                }
                { !isPending && 
                    <div className='Question'>
                        <div className='question-section'>
                        <div className='question-count'>
                            <span>Do you want to proceed to submit your responses?</span>
                        </div>
                        <div className='answer-section'>
                        <button className='primary' onClick={handleSubmit}>
                            Yes
                        </button>
                        
                        <button className='secondary' onClick={handleGoback}>
                            No, go back
                        </button>
                        </div>
                        </div>
                    </div>
                }
        
        </div>
    );
};

export default RegisterVote;