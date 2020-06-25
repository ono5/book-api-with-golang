import React,{ Component } from 'react';
import './App.css';
import Books from './containers/Books';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';

class App extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <div className="App">
        <Books />
      </div>
    )
  }
}

export default App;
