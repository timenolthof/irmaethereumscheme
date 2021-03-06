import 'babel-polyfill'
import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import GrommetApp from 'grommet/components/App';
import {
  Form, Header, FormFields, Footer, Button, Heading, TextInput
} from 'grommet';

import { deployNewScheme } from '../reducer/schemes';

class NewScheme extends Component {

  constructor(props) {
    super(props);
    this.state = {
      formSchemeId: 'new-ethereum-scheme',
      formSchemeName: 'Awesome Scheme on Ethereum',
    };
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value,
    });
  }

  handleSubmit(event) {
    const metadata = {
      version: 7,
      id: this.state.formSchemeId,
      name: [{
        lang: 'en',
        name: this.state.formSchemeName
      }]
    };
    this.props.deployNewScheme(this.state.formSchemeId, metadata);
    this.props.history.push('/');
  }

  render() {
    return (
      <GrommetApp>
        Coinbase: { this.state.coinbase }
        Balance: { this.state.balance }
        <Form pad="medium" >
          <Header>
            <Heading>
              New Scheme
            </Heading>
          </Header>
          <FormFields>
          Scheme Id:<br/>
          <TextInput id='schemeId'
            placeHolder="scheme id"
            style={{ width: '100%' }}
            value={this.state.formSchemeId}
            onDOMChange={this.handleInputChange.bind(this)}
            name='formSchemeId'/><br/>

          Scheme name:<br/>
          <TextInput id='schemeName'
            placeHolder="scheme name"
            style={{ width: '100%' }}
            value={this.state.formSchemeName}
            onDOMChange={this.handleInputChange.bind(this)}
            name='formSchemeName'/><br/>

          </FormFields>
          <Footer pad={{"vertical": "medium"}}>
            <Button label='Submit'
              primary={true}
              onClick={() => this.handleSubmit()} />
          </Footer>
        </Form>
      </GrommetApp>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => bindActionCreators({
  deployNewScheme,
}, dispatch);

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewScheme);
