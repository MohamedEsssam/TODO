import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Home from "./views/home/home";
import Register from "./views/register/register";
import "./App.css";

function App() {
  return (
    <React.Fragment>
      <ToastContainer />
      <BrowserRouter>
        <Switch>
          <Route path="/register" component={Register}></Route>
          <Route path="/" component={Home}></Route>
        </Switch>
      </BrowserRouter>
    </React.Fragment>
  );
}

export default App;
