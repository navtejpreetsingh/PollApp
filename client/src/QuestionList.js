// import Paper from '@material-ui/core/Paper';
import {
    ArgumentAxis,
    ValueAxis,
    Chart,
    BarSeries,
} from '@devexpress/dx-react-chart-material-ui';

const QuestionList = ({questions, title}) => {
    
    // const data = [
    //     { argument: 'Monday', value: 30 },
    //     { argument: 'Tuesday', value: 20 },
    //     { argument: 'Wednesday', value: 10 },
    //     { argument: 'Thursday', value: 50 },
    // ];

    return (
        <div className="Question-List">
            <h4>{title}</h4>
            {questions.map((ques) => (
                <div className="Question view-only" key={ques.qid}>
                    
                    {/* <h5>{ques.qid} .) {ques.question}</h5>
                    {ques.options.map((opt) => (
                        <div className="Option" key={opt.qid+"_"+opt.option_id}><h6>({opt.qid+"_"+opt.option_id}) Votes:{opt.votes} {opt.option}</h6></div>
                    ))} */}
                    <div className='question-section'>
                        <div className='question-count'>
                            <span>Question {ques.qid}</span>/{questions.length}
                        </div>
                        <div className='question-text'>{ques.question}</div>
                    </div>
                    <div className='answer-section'>
                        {ques.options.map((opt) => (
                            <div className='option-details' key={"div_"+opt.qid+"_"+opt.option_id}>
                                <span className='votes' key={"votes_"+opt.qid+"_"+opt.option_id}>{opt.votes}</span>
                                <span className='option' key={"option_"+opt.qid+"_"+opt.option_id}>{opt.option}</span>
                            </div>
                        ))}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default QuestionList;