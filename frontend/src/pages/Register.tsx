import React, { ChangeEvent, FormEvent, useState } from 'react';
import axios from 'axios';
import { Redirect } from 'react-router';

interface RegisterAccount {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  confirmPass: string;
}

const Register: React.FC = () => {
  const [account, setAccount] = useState<RegisterAccount>({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    confirmPass: '',
  });
  const [redirect, setRedirect] = useState(false);

  if (redirect) {
    return <Redirect to="/login" />;
  }

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setAccount({
      ...account,
      [e.target.id]: e.target.value,
    });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await axios.post<RegisterAccount>('auth/register', {
      first_name: account.firstName,
      last_name: account.lastName,
      email: account.email,
      password: account.password,
      confirm_pass: account.confirmPass,
    });
    if (response.status === 200) {
      setRedirect(true);
    }
  };

  return (
    <form className="form-signin" onSubmit={handleSubmit}>
      <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input
        id="firstName"
        type="text"
        className="form-control"
        placeholder="First name"
        required={true}
        onChange={handleChange}
      />
      <input
        id="lastName"
        type="text"
        className="form-control"
        placeholder="Last name"
        required={true}
        onChange={handleChange}
      />
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
      <input
        id="confirmPass"
        type="password"
        className="form-control"
        placeholder="Confirm Password"
        required={true}
        onChange={handleChange}
      />
      <button className="btn btn-lg btn-primary btn-block" type="submit">
        Register
      </button>
    </form>
  );
};

export default Register;
