import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";

class App extends Component {
  onClick() {
    window.location = "http://localhost:8000/login";
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Home</h1>
        </header>{" "}
        <h1>Login complete</h1>
      </div>
    );
  }
}

export default App;
