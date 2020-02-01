import React from 'react';
import Page from './Page';
import { Alert } from 'react-bootstrap';

const Find = '/find/'

class Redirect extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      alertShow: false,
      alertValue: '',
      alertVariant: '',
    };
  }

  componentDidMount() {
    fetch(process.env.REACT_APP_API_URL + Find + this.props.match.params.url, {
      method: 'GET',
    })
    .then(res => res.json()
      .then(data => {
        if(res.status === 500) {
          this.setState({
            alertShow: true,
            alertVariant: 'info',
            alertValue: data.exception
          })
          return
        } else if(res.status !== 200) {
          this.setState({ 
            alertShow: true,
            alertVariant: 'danger',
            alertValue: data.exception
          })
          return
        }

        this.setState({ 
          alertShow: true,
          alertVariant: 'success',
          alertValue: 'Redirecting to ' + data.full_url,
        })

        window.location.href = data.full_url;
      })
    )
  }

  render() {
    return (
      <Page>
        <div id='redirect'>
          <Alert id='alert' show={this.state.alertShow} variant={this.state.alertVariant}>{this.state.alertValue}</Alert>
        </div>
      </Page>
    );
  }
}

export default Redirect;
