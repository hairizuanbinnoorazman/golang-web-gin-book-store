import React, { Component } from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import "./App.css";

import Main from "./Main";
import Home from "./Home";
import Callback from "./Callback";

class App extends Component {
  onClick() {
    window.location = "http://localhost:8000/login";
  }

  render() {
    return (
      <Router>
        <div>
          <Route exact path="/" component={Main} />
          <Route exact path="/callback" component={Callback} />
          <Route exact path="/home" component={Home} />
        </div>
      </Router>
    );
  }
}

export default App;
