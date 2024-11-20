import { useState } from 'react';
import { Box, TextField, Button } from '@mui/material';

const MessageForm = ({ onSendMessage }) => {
    const [message, setMessage] = useState('');

    const handleSendMessage = (event) => {
        event.preventDefault();
        if (message !== '') {
            onSendMessage(message);
            setMessage('');
        }
    };

    return (
        <form onSubmit={handleSendMessage} style={{ display: 'flex', marginTop: 'auto', padding: '1em 0' }}>
            <Box display="flex" justifyContent="space-between" px={2} width="100%">
                <TextField
                    multiline
                    maxRows={4}
                    variant="outlined"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    onKeyPress={(e) => {
                        if (e.key === 'Enter' && !e.shiftKey) {
                            handleSendMessage(e);
                        }
                    }}
                    style={{ flex: '1' }}
                />
                <Button variant="contained" onClick={handleSendMessage} style={{ marginLeft: '1em' }}>
                    Send
                </Button>
            </Box>
        </form>
    );
};

export default MessageForm;