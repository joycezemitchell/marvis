import React, { useContext, useState } from 'react';
import { TextField, Button, Box, Typography } from '@mui/material';
import ApiContext from "../../ApiContext";

const LoginPage = ({ onLogin }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const apiUrl = useContext(ApiContext);

    const handleLogin = async () => {
        try {
            const response = await fetch(`${apiUrl}login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const data = await response.json();
            const token = data.token;

            onLogin(username, token);
        } catch (error) {
            console.error('Error:', error);
            setErrorMessage('Invalid username or password');
        }
    };

    return (
        <Box
            display="flex"
            flexDirection="column"
            justifyContent="center"
            alignItems="center"
            minHeight="100vh"
        >
            <Typography variant="h4" component="h2" gutterBottom>
                Sign in
            </Typography>
            <Box
                component="form"
                display="flex"
                flexDirection="column"
                gap={2}
                onSubmit={(e) => {
                    e.preventDefault();
                    handleLogin();
                }}
            >
                <TextField
                    label="Username"
                    variant="outlined"
                    onChange={(e) => setUsername(e.target.value)}
                />
                <TextField
                    type="password"
                    label="Password"
                    variant="outlined"
                    onChange={(e) => setPassword(e.target.value)}
                />
                <Button variant="contained" type="submit">
                    Login
                </Button>
                {errorMessage && (
                    <Typography variant="body2" color="error">
                        {errorMessage}
                    </Typography>
                )}
            </Box>
        </Box>
    );
};

export default LoginPage;