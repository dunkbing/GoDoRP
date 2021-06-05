import React from 'react';

const Register: React.FC = () => {
  return (
    <form className="form-signin">
      <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
      <input
        type="text"
        className="form-control"
        placeholder="First name"
        required={true}
      />
      <input
        type="text"
        className="form-control"
        placeholder="Last name"
        required={true}
      />
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
      <input
        type="password"
        className="form-control"
        placeholder="Confirm Password"
        required={true}
      />
      <button className="btn btn-lg btn-primary btn-block" type="submit">
        Register
      </button>
    </form>
  );
};

export default Register;
