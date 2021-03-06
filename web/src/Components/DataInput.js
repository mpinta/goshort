import React from 'react';
import './DataInput.css';
import Page from './Page';
import { Form, Button, InputGroup, Alert } from 'react-bootstrap';

const Shorten = '/shorten';
const Limit = '/shorten/limit';

class DataInput extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      alertShow: false,
      alertValue: '',
      alertVariant: '',
      buttonValue: 'goshort',
      buttonSwitch: false,
      checkboxDisabled: false,
      minutesDisabled: false
    };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleFocus = this.handleFocus.bind(this);
    this.handleBlur = this.handleBlur.bind(this);
  }

  handleSubmit(e) {
    e.preventDefault();

    if(this.state.buttonSwitch) {
      this.handleCopy();
      return
    }

    let reqPath;
    let reqBody;
    let urlInput = document.getElementById('urlInput');
    let checkbox = document.getElementById('checkbox');
    let minutesInput = document.getElementById('minutesInput');
    
    if(checkbox.checked) {
      reqPath = Shorten;
      reqBody = JSON.stringify({
        full_url: urlInput.value
      })
    } else {
      reqPath = Limit;
      reqBody = JSON.stringify({
        full_url: urlInput.value,
        minutes_valid: parseInt(minutesInput.value)
      })
    }

    fetch(process.env.REACT_APP_API_URL + reqPath, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      mode: 'cors',
      cache: 'no-cache',
      credentials: 'same-origin',
      body: reqBody
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
        } else if(res.status !== 201) {
          this.setState({ 
            alertShow: true,
            alertVariant: 'danger',
            alertValue: data.exception
          })
          return
        }
        
        let value;
        if(checkbox.checked) {
          value = 'Your short URL is valid with no time limit!';
        } else {
          value = 'Your short URL is valid until ' + new Date(data.valid_until).toUTCString();
        }

        this.setState({ 
          alertShow: true,
          alertVariant: 'success',
          alertValue: value,
          buttonValue: 'Copy URL',
          buttonSwitch: true
        })

        urlInput.value = window.location.href + data.short_url;
      })
    )
    .catch(
      this.setState({
        alertShow: true,
        alertVariant: 'info',
        alertValue: 'Network error occurred!'
      })
    )
  }

  handleChange() {
    if(this.state.minutesDisabled) {
      this.setState({
        minutesDisabled: false
      })
      return
    }

    this.setState({
      minutesDisabled: true
    })
  }

  handleFocus() {
    this.setState({
      checkboxDisabled: true
    })
  }

  handleBlur() {
    let minutesInput = document.getElementById('minutesInput');
    if(minutesInput.value.length === 0) {
      this.setState({
        checkboxDisabled: false
      })
    }
  }

  handleCopy() {
    let urlInput = document.getElementById('urlInput'); 
    let shortenButton = document.getElementById('shortenButton');
    
    urlInput.select();
    urlInput.setSelectionRange(0, 99999);
    document.execCommand('copy');

    urlInput.classList.add('urlInputCopied');
    shortenButton.classList.add('shortenButtonCopied');
    this.setState({ 
      buttonValue: 'Copied!'
    })

    setTimeout(function(){
      shortenButton.classList.remove('shortenButtonCopied');
        this.setState({ 
        buttonValue: 'Copy URL'
      })
    }.bind(this), 1000);
  }

  render() {
    return (
      <Page>
        <div id='dataInput'>
          <Form onSubmit={(e) => {this.handleSubmit(e)}}>
            <Form.Group>
              <Form.Control id='urlInput' type='url' placeholder='URL you want to shorten' required />
            </Form.Group>
            <Form.Group>
                <Form.Check>
                  <Form.Check.Input id='checkbox' type='checkbox' onChange={this.handleChange} disabled={this.state.checkboxDisabled} required />
                  <Form.Check.Label htmlFor='checkbox'>No time limit</Form.Check.Label>
                </Form.Check>
            </Form.Group>
            <Form.Group>
              <InputGroup>
                <Form.Control id='minutesInput' type='number' placeholder='Minutes (1-60)' min='1' max='60' onFocus={this.handleFocus} onBlur={this.handleBlur} disabled={this.state.minutesDisabled} required />
                <Button id='shortenButton' type='submit' variant='primary'>{this.state.buttonValue}</Button>
              </InputGroup>
            </Form.Group>
          </Form>
          <Alert id='alert' show={this.state.alertShow} variant={this.state.alertVariant}>{this.state.alertValue}</Alert>
        </div>
      </Page>
    );
  }
}

export default DataInput;
