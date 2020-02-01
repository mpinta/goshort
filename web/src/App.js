import React from 'react';
import './App.css';
import DataInput from './Components/DataInput';
import Redirect from './Components/Redirect';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path='/' exact component={ DataInput } />
        <Route path='/:url' exact component={ Redirect } />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
