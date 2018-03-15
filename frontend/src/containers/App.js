import React from 'react';
import { Route } from 'react-router-dom';
import Home from './Home';
import NewScheme from './NewScheme';

const App = () => (
  <div>
    <Route exact path="/" component={Home} />
    <Route exact path="/new-scheme" component={NewScheme} />
  </div>
)

export default App;
