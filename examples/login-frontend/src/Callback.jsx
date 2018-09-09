import React, { Component } from "react";
import logo from "./logo.svg";
import axios from "axios";
import "./App.css";

class App extends Component {
  componentDidMount() {
    console.log(window.location.search);
    var lol = new URLSearchParams(window.location.search);
    axios
      .post("http://localhost:8000/oauth_callback", {
        code: lol.get("code")
      })
      .catch(err => {
        console.log("Error");
        console.log(err);
      })
      .then(response => {
        console.log(response);
        this.props.history.push("/home");
      });
  }

  onClick() {
    window.location = "http://localhost:8000/login";
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Callback</h1>
        </header>{" "}
        <p>Authentication in Progress. Please wait</p>
      </div>
    );
  }
}

export default App;
