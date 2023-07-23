import React from "react";
import { Routes, Route, Link } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

import AddRecord from "./components/createVaclog";
import Record from "./components/vaclog";
import Records from "./components/vaclogList";

const App: React.FC = () => {
  return (
    <div>
      <nav className="navbar navbar-expand navbar-dark bg-dark">
        <a href="/records" className="navbar-brand">
          Home
        </a>
        <div className="navbar-nav mr-auto">
          <li className="nav-item">
            <Link to={"/records"} className="nav-link">
              Records
            </Link>
          </li>
          <li className="nav-item">
            <Link to={"/add"} className="nav-link">
              Add
            </Link>
          </li>
        </div>
      </nav>

      <div className="container mt-3">
        <Routes>
          <Route path="/" element={<Records/>} />
          <Route path="/records" element={<Records/>} />
          <Route path="/add" element={<AddRecord/>} />
          <Route path="/records/:id" element={<Record/>} />
        </Routes>
      </div>
    </div>
  );
}

export default App;