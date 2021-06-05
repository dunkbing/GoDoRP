import React from 'react';
import { Link } from 'react-router-dom';

const Login: React.FC = () => {
  return (
    <form className="form-signin">
      <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input
        type="email"
        className="form-control"
        placeholder="Email"
        required={true}
      />
      <input
        type="password"
        className="form-control"
        placeholder="Password"
        required={true}
      />
      <div className="mb-3">
        <Link to="/forgot">Forgot Password?</Link>
      </div>
      <button className="btn btn-lg btn-primary btn-block" type="submit">
        Sign in
      </button>
    </form>
  );
};

export default Login;
