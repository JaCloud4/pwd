import React from 'react';

const message="Welcome!!";
const welcome= <h1>I said {{message}}</h1>;
ReactDOM.render(
welcome,
  document.getElementById('root')
);
