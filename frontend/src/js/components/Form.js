import React, { Component } from 'react';
import request from 'superagent';

class Form extends Component {
    shortenUrl = (e) => {
      e.preventDefault();

      request.post('/generate')
        .send({ Url: this.refs.url.value })
        .end((err, res) => {
          if (err) return alert(err);

          const payload = JSON.parse(res.text);
          this.props.setUrl(payload.Url);
        });
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
