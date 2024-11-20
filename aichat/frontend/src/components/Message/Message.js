import { ListItem, ListItemText, Typography, Avatar } from '@mui/material';
import PersonIcon from '@mui/icons-material/Person';
import AndroidIcon from '@mui/icons-material/Android';

const Message = ({ user, text }) => {
    return (
        <ListItem
            alignItems="flex-start"
            sx={{
                backgroundColor: user === 'User' ? 'blue.100' : 'grey.100',
                borderRadius: 1,
                px: 1,
                my: 1,
                border: "1px solid",
                borderColor: user === 'User' ? 'blue.500' : 'grey.500'
            }}
        >
            <Avatar sx={{ marginRight: 1 }}>
                {user === 'User' ? <PersonIcon /> : <AndroidIcon />}
            </Avatar>
            <ListItemText
                primary={text}
                secondary={<Typography variant="body2">{user}</Typography>}
            />
        </ListItem>
    );
};

export default Message;