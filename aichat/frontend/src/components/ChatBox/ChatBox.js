import {useState, useEffect, useRef, useContext} from 'react';
import { Box, List, CircularProgress} from '@mui/material';
import Message from "../Message/Message";
import TypingMessage from "../TypingMessage/TypingMessage";
import MessageForm from "../MessageForm/MessageForm";
import './ChatBox.css'
import ApiContext from "../../ApiContext";

const ChatBox = ({ username, token }) => {
    const [messages, setMessages] = useState([]);
    const [typing, setTyping] = useState(false);
    const [typingMessage, setTypingMessage] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const endOfMessagesRef = useRef(null);
    const apiUrl = useContext(ApiContext);

    const scrollToBottom = () => {
        endOfMessagesRef.current.scrollIntoView({ behavior: 'smooth' });
    };

    useEffect(scrollToBottom, [messages]);

    const handleSendMessage = (message) => {
        if (!typing) {
            setMessages((oldMessages) => [...oldMessages, { user: 'User', text: message }]);
            getAIResponse(message);
        }
    };

    const getAIResponse = async (message) => {
        try {
            setIsLoading(true); // Set isLoading to true while waiting for the fetch request
            const response = await fetch(`${apiUrl}chat`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}` // Include the token in the Authorization header
                },
                body: JSON.stringify({ message }),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            setIsLoading(false);
            const responseData = await response.json();
            simulateTyping(responseData.message);
        } catch (error) {
            console.error('Error:', error);
        }
    };

    const simulateTyping = (replyText) => {
        setTyping(true);
        let currentReplyText = '';

        const addCharacter = (char, delay) => {
            return new Promise((resolve) => {
                setTimeout(() => {
                    currentReplyText += char;
                    setTypingMessage(currentReplyText);
                    resolve();
                }, delay);
            });
        };

        const typeMessage = async () => {
            for (let char of replyText) {
                await addCharacter(char, 50);
            }

            let newReply = { user: 'AI', text: currentReplyText };
            setMessages((oldMessages) => [...oldMessages, newReply]);
            setTypingMessage('');
            setTyping(false);
        };

        typeMessage();
    };

    return (
        <div className="content" >
            <Box sx={{ height: '100vh', display: 'flex', flexDirection: 'column', py: 2, px: 2, maxWidth: '100%', margin: 'auto', boxSizing: 'border-box'}}>
                <List sx={{ overflow: 'auto', flexGrow: 1 }}>
                    {messages.map((msg, index) => (
                        <Box key={index}>
                            <Message user={msg.user} text={msg.text} />
                        </Box>
                    ))}
                    {isLoading && (
                        <Box sx={{ display: 'flex', alignItems: 'flex-start', pl: 2, marginTop: '10px' }}>
                            <CircularProgress variant="indeterminate" size={20} color="secondary" sx={{marginRight: '8px' }} /> thinking ...
                        </Box>
                    )}
                    {typing && <TypingMessage typingMessage={typingMessage} />}
                    <div ref={endOfMessagesRef} />
                </List>
                <MessageForm onSendMessage={handleSendMessage} />
            </Box>
        </div>
    );
};

export default ChatBox;
