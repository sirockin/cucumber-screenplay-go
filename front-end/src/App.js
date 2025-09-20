import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import SignUp from './components/SignUp';
import Login from './components/Login';
import Account from './components/Account';
import Activate from './components/Activate';
import Projects from './components/Projects';
import Clear from './components/Clear';

function App() {
  return (
    <Router>
      <div className="container">
        <nav className="nav">
          <Link to="/signup">Sign Up</Link>
          <Link to="/login">Login</Link>
          <Link to="/admin/clear">Clear Data</Link>
        </nav>

        <Routes>
          <Route path="/signup" element={<SignUp />} />
          <Route path="/login" element={<Login />} />
          <Route path="/account/:name" element={<Account />} />
          <Route path="/activate/:name" element={<Activate />} />
          <Route path="/account/:name/projects" element={<Projects />} />
          <Route path="/admin/clear" element={<Clear />} />
          <Route path="/" element={<SignUp />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;