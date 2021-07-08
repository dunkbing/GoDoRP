import axios from 'axios';
import React, { FormEvent, useState } from 'react';
import { ChangeEvent } from 'react';
import { Link, Redirect } from 'react-router-dom';

interface LoginAccount {
  email: string;
  password: string;
}

const Login: React.FC<{ setLoggedIn: () => void }> = ({ setLoggedIn }) => {
  const [account, setAccount] = useState<LoginAccount>({
    email: '',
    password: '',
  });
  const [redirect, setRedirect] = useState(false);

  if (redirect) {
    return <Redirect to="/" />;
  }

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setAccount({
      ...account,
      [e.target.id]: e.target.value,
    });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await axios.post<LoginAccount>('login', {
      email: account.email,
      password: account.password,
    });
    if (response.status === 200) {
      setRedirect(true);
      setLoggedIn();
    }
  };

  return (
    <form className="form-signin" onSubmit={handleSubmit}>
      <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input
        id="email"
        type="email"
        className="form-control"
        placeholder="Email"
        required={true}
        onChange={handleChange}
      />
      <input
        id="password"
        type="password"
        className="form-control"
        placeholder="Password"
        required={true}
        onChange={handleChange}
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
