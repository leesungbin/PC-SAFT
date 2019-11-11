import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './routes/Home';
import Document from './routes/Documents';
import Database from './routes/Database';

type State = {
  width: number,
  height: number,
};
class App extends React.Component<{}, State> {
  state: State = {
    width: 0,
    height: 0,
  }
  componentDidMount = () => {
    this.updateWindowDimensions();
    window.addEventListener('resize', this.updateWindowDimensions);
  }

  componentWillUnmount = () => {
    window.removeEventListener('resize', this.updateWindowDimensions);
  }

  updateWindowDimensions = () => {
    this.setState({ width: window.innerWidth, height: window.innerHeight });
  }

  render() {
    return (
      <Router>
        <Header width={this.state.width} />
        <Switch>
          <Route path="/docs" ><Document /></Route>
          <Route path="/db"><Database /></Route>
          <Route exact path="/"><Home width={this.state.width} height={this.state.height}/></Route>
        </Switch>
      </Router>
    );
  }
}

export default App;