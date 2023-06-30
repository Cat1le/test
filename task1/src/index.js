import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.css'

/** @type {any} */
const el = document.getElementById('root');
const root = ReactDOM.createRoot(el);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
