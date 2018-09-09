import React, { Component } from "react";
import logo from "./logo.svg";
import Button from "@material-ui/core/Button";
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
          <h1 className="App-title">Main</h1>
        </header>{" "}
        <p className="App-intro">
          To get started, edit <code> src / App.js </code> and save to reload.{" "}
        </p>{" "}
        <Button onClick={() => this.onClick()} color="primary">
          Login{" "}
        </Button>{" "}
      </div>
    );
  }
}

export default App;
