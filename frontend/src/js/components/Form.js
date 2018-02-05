import React, {Component} from 'react';
import * as api from '../api';

class Form extends Component {
  shortenUrl = (e) => {
    e.preventDefault();
    const { url, singleUse } = this.refs;
    api.generateCode(url.value, singleUse.checked, this.props.setUrl);
  };

  render() {
    return (
      <form onSubmit={this.shortenUrl} className="App-form">
        <input type="url" ref="url"/>
        <br/>
        <input type="checkbox" ref="singleUse" name="single-use"/>
        <label htmlFor="single-use">Single-Use Only</label>
        <br/>
        <input type="submit" value="Shorten"/>
      </form>
    );
  }
}

export default Form;
