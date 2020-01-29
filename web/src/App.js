import React from 'react';
import './App.css';
import { Container, Row, Col } from 'react-bootstrap';
import DataInput from './DataInput/DataInput.js'
import Banner from './Banner/Banner';

function App() {
  return (
    <div id="app">
      <Container>
          <Row>
            <Col xs={0} sm={0} md={1} large={2} xl={2} />
            <Col xs={12} sm={10} md={10} large={8} xl={8}>
              <Banner />
              <DataInput />
            </Col>
            <Col xs={0} sm={0} md={1} large={2} xl={2} />
          </Row>
      </Container>
    </div>
  );
}

export default App;
