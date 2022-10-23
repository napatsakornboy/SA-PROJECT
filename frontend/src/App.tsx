import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import BASKETCreate from "./components/BASKETCreate";

export default function App() {

return (
  <Router>
   <div>
   <Navbar />
   <BASKETCreate/>
   <Routes>
   </Routes>
   </div>
  </Router>
);

}