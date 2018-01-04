import React, {Component} from 'react';
import * as api from '../api';

class Form extends Component {
  shortenUrl = (e) => {
    e.preventDefault();
    api.generateCode(this.refs.url.value, this.props.setUrl);
  };

  render() {
    return (
      <form onSubmit={this.shortenUrl} className="App-form">
        <input type="url" ref="url"/>
        <br/>
        <input type="submit" value="Shorten"/>
      </form>
    );
  }
}

export default Form;
