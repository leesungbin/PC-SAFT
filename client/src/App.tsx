import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './routes/Home';
import Document from './routes/Document';
import Database from './routes/Database';

class App extends React.Component {
  render() {
    return (
      <Router>
        <Header />
        <Switch>
          <Route path="/docs"><Document /></Route>
          <Route path="/db"><Database /></Route>
          <Route path="/"><Home /></Route>
        </Switch>
      </Router>
    );
  }
}

export default App;