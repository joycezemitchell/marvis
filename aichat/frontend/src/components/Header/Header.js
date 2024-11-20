import './Header.css';
import ChatIcon from '@mui/icons-material/Chat';

const Header = () => {
    return (
        <div className="header">
            <ChatIcon className="logo" />
            <h1 className="title">Ally Chat</h1>
        </div>
    );
};

export default Header;