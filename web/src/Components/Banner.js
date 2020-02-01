import React from 'react';
import './Banner.css';
import { Jumbotron, Button } from 'react-bootstrap';

class Banner extends React.Component {
  render() {
    return (
      <Jumbotron id='jumbotron'>
        <h1>goshort</h1>
        <p>Simple URL shortener written in React and Golang, using SQLite.</p>
        <Button id='projectButton' href='https://github.com/mpinta/goshort'>Project on Github</Button>
      </Jumbotron>
    );
  }
}

export default Banner;
