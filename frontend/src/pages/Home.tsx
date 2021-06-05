import React from 'react';

const Home: React.FC<{ user: any }> = ({ user }) => {
  let message;
  if (user) {
    message = `Hi ${user.first_name} ${user.last_name}`;
  } else {
    message = 'you are not logged in';
  }

  return <div className="container">{message}</div>;
};

export default Home;
