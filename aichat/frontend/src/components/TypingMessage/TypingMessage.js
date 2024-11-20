import { ListItem, ListItemText, Typography, Avatar } from '@mui/material';
import AndroidIcon from '@mui/icons-material/Android';

const TypingMessage = ({ typingMessage }) => {
    return (
        <ListItem
            alignItems="flex-start"
            sx={{
                backgroundColor: 'grey.100',
                borderRadius: 1,
                px: 1,
                my: 1,
                border: "1px solid",
                borderColor: 'grey.500'
            }}
        >
            <Avatar sx={{ marginRight: 1 }}>
                <AndroidIcon />
            </Avatar>
            <ListItemText primary={<span>{typingMessage}<span className="blinking-cursor">|</span></span>} secondary={<Typography variant="body2">AI</Typography>} />
        </ListItem>
    );
};

export default TypingMessage;