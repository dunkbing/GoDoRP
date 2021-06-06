import axios from 'axios';
import React, { SyntheticEvent, useState } from 'react';

const Forgot: React.FC = props => {
  const [email, setEmail] = useState('');
  const [notify, setNotify] = useState({
    show: false,
    error: false,
    message: '',
  });

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    try {
      await axios.post('forgot', { email });

      setNotify({
        show: true,
        error: false,
        message: 'Email was sent',
      });
    } catch {
      setNotify({
        show: true,
        error: true,
        message: 'Email does not exist',
      });
    }
  };

  let info;
  if (notify.show) {
    info = (
      <div
        className={notify.error ? 'alert alert-danger' : 'alert alert-success'}
        role="alert"
      >
        {notify.message}
      </div>
    );
  }

  return (
    <form className="form-signin" onSubmit={handleSubmit}>
      {info}
      <h1 className="h3 mb-3 font-weight-normal">Enter your email</h1>
      <input
        id="email"
        type="email"
        className="form-control mb-3"
        placeholder="Email"
        required
        onChange={e => setEmail(e.target.value)}
      />
      <button className="btn btn-lg btn-primary btn-block" type="submit">
        Send Email
      </button>
    </form>
  );
};

export default Forgot;
