import React from 'react';
import logo from '../images/logo.svg';
import '../styles/App.css';
import {connect, sendMessage} from "../api/"

function ChatApp() {
  connect()
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          What a great chat app!
        </p>
        <button onClick={send}></button>
      </header>
    </div>
  );
}

function send() {
  sendMessage("Hello")
}

export default ChatApp;
