import axios from 'axios';
import { useEffect, useState } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import './App.css';
import Nav from './components/Nav';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  const [user, setUser] = useState(null);
  const [loggedIn, setLoggedIn] = useState(false);
  useEffect(() => {
    (async () => {
      try {
        const response = await axios.get('user');
        setUser(response.data);
      } catch (error) {
        setUser(null);
      }
    })();
  }, [loggedIn]);
  return (
    <div className="App">
      <BrowserRouter>
        <Nav user={user} setLoggedIn={() => setLoggedIn(false)} />
        <Route path="/" exact component={() => <Home user={user} />} />
        <Route
          path="/login"
          component={() => <Login setLoggedIn={() => setLoggedIn(true)} />}
        />
        <Route path="/register" component={Register} />
      </BrowserRouter>
    </div>
  );
}

export default App;
