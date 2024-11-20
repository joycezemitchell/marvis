import React, { useState } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Switch } from 'react-router-dom';
import ChatBox from "./components/ChatBox/ChatBox";
import Header from "./components/Header/Header";
import LoginPage from "./components/Login/Login";
import ApiContext from './ApiContext';
import './App.css';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [username, setUsername] = useState('');
    const [token, setToken] = useState('');
    const [apiUrl, setApiUrl] = useState(process.env.REACT_APP_API_URL);

    const handleLogin = (username, token) => {
        setIsLoggedIn(true);
        setUsername(username);
        setToken(token);
    };

    return (
        <ApiContext.Provider value={apiUrl}>
            <Router>
                <div className="App">
                    <Header />
                    <Switch>
                        <Route exact path="/">
                            {isLoggedIn ? (
                                <ChatBox username={username} token={token} />
                            ) : (
                                <LoginPage onLogin={handleLogin} />
                            )}
                        </Route>
                    </Switch>
                </div>
            </Router>
        </ApiContext.Provider>
    );
}

export default App;