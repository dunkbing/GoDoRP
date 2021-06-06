import axios from 'axios';
import React, { ChangeEvent, SyntheticEvent, useState } from 'react';
import { Redirect, useParams } from 'react-router';

interface ResetPassword {
  password: string;
  confirmPass: string;
}

const Reset: React.FC<{ match: any }> = ({ match }) => {
  const [resetPassword, setResetPassword] = useState<ResetPassword>({
    password: '',
    confirmPass: '',
  });
  const [redirect, setRedirect] = useState(false);
  const { token } = useParams<any>();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setResetPassword({
      ...resetPassword,
      [e.target.id]: e.target.value,
    });
  };

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    await axios.post('reset', {
      token,
      password: resetPassword.password,
      confirm_pass: resetPassword.confirmPass,
    });

    setRedirect(true);
  };

  if (redirect) {
    return <Redirect to="/login" />;
  }

  return (
    <form className="form-signin" onSubmit={handleSubmit}>
      <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input
        id="password"
        type="password"
        className="form-control mb-3"
        placeholder="Password"
        required
        onChange={handleChange}
      />
      <input
        id="confirmPass"
        type="password"
        className="form-control mb-3"
        placeholder="Confirm password"
        required
        onChange={handleChange}
      />
      <button className="btn btn-lg btn-primary btn-block" type="submit">
        Send Email
      </button>
    </form>
  );
};

export default Reset;
