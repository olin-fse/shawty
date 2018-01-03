import React, { Component } from 'react';

import Form from '../components/Form';
import '../../css/App.css';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      shortenedUrl: ''
    };
  }

  setUrl = (shortenedUrl) => {
    this.setState({ shortenedUrl });
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
