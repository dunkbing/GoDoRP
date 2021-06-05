import React from 'react';
import { Link } from 'react-router-dom';

const Nav = () => {
  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark">
      <ul className="navbar-nav mr-auto">
        <li className="nav-item">
          <Link to="/" href="#" className="nav-link">
            Home
          </Link>
        </li>
      </ul>
      <ul className="navbar-nav my-2 my-lg-8">
        <li className="nav-item">
          <Link to="login" href="#" className="nav-link">
            Login
          </Link>
        </li>
        <li className="nav-item">
          <Link to="register" href="#" className="nav-link">
            Register
          </Link>
        </li>
      </ul>
    </nav>
  );
};

export default Nav;
