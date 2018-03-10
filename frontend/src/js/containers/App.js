import React, { Component } from 'react';

import Form from '../components/Form';
import '../../css/App.css';
import getConfig from '../../config';

const config = getConfig(process.env.NODE_ENV);

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      shortenedUrl: ''
    };
  }

  setUrl = (code) => {
    this.setState({ shortenedUrl: `${config.staticEndpoint}/${code}` });
  };

  render() {
    return (
      <div className="App">
        <h1 className="App-title">Welcome to Shawty</h1>
        <Form setUrl={this.setUrl}/>
        <div className="App-result">
          {this.state.shortenedUrl}
        </div>
      </div>
    );
  }
}

export default App;
