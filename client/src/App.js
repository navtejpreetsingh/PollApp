import './App.css';
import {BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './Home';
import RegisterVote from './RegisterVote';

function App() {
  return (
    <Router>
    <div className="App">
      <header className="App-header">Poll App
        <Routes>
          <Route exact path="/" element={<Home className="Inner-Block"/>} />
          <Route path="/vote" element={<RegisterVote />} />
        </Routes>  
      </header>
    </div>
    </Router>
  );
}

export default App;