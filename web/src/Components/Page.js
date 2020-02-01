import React from 'react';
import './Page.css';
import Banner from './Banner';
import { Container, Row, Col } from 'react-bootstrap';

class Page extends React.Component {
  render() {
    return (
      <Container>
        <Row>
          <Col xs={0} sm={0} md={1} large={2} xl={2} />
          <Col xs={12} sm={10} md={10} large={8} xl={8}>
            <Banner />
            { this.props.children }
          </Col>
          <Col xs={0} sm={0} md={1} large={2} xl={2} />
        </Row>
      </Container>
    );
  }
}

export default Page;
