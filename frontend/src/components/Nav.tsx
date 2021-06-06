import axios from 'axios';
import { FC } from 'react';
import { Link } from 'react-router-dom';

const Nav: FC<{ user: any; setLoggedIn: () => void }> = ({
  user,
  setLoggedIn,
}) => {
  let logoutLink;
  if (user) {
    const logout = async (e: React.MouseEvent<HTMLAnchorElement>) => {
      await axios.post('logout');
      setLoggedIn();
    };
    logoutLink = (
      <ul className="navbar-nav my-2 my-lg-0">
        <li className="nav-item">
          <Link onClick={logout} to="login" className="nav-link">
            Logout
          </Link>
        </li>
      </ul>
    );
  } else {
    logoutLink = (
      <ul className="navbar-nav my-2 my-lg-0">
        <li className="nav-item">
          <Link to="login" className="nav-link">
            Login
          </Link>
        </li>
        <li className="nav-item">
          <Link to="register" className="nav-link">
            Register
          </Link>
        </li>
      </ul>
    );
  }
  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark">
      <ul className="navbar-nav mr-auto">
        <li className="nav-item">
          <Link to="/" className="nav-link">
            Home
          </Link>
        </li>
      </ul>
      {logoutLink}
    </nav>
  );
};

export default Nav;
