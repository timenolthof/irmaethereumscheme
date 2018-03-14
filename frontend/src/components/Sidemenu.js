import 'babel-polyfill'
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Sidebar from 'grommet/components/Sidebar';
import Header from 'grommet/components/Header';
import Title from 'grommet/components/Title';
import Box from 'grommet/components/Box';
import Menu from 'grommet/components/Menu';
import Anchor from 'grommet/components/Anchor';
import Button from 'grommet/components/Button';
import User from 'grommet/components/icons/base/User';
import Footer from 'grommet/components/Footer';

import {
  selectScheme,
  addScheme,
} from '../reducer/schemes'

class Sidemenu extends Component {
  static propTypes = {
    schemes: PropTypes.array.isRequired,
    selectedSchemeIndex: PropTypes.number,
  };

  render() {
    return (
      <Sidebar
        colorIndex='neutral-1'>
        <Header pad='medium'
          justify='between'>
          <Title>
            Schemes
          </Title>
        </Header>
        <Box flex='grow'
          justify='start'>
          <Menu primary={true}>
            { this.props.schemes.map((scheme, index) =>
              <Anchor key={index} onClick={() => this.props.selectScheme(index)}
                className={index === this.props.selectedSchemeIndex ? 'active' : ''}>
                {scheme.id}
              </Anchor>
            )}
          </Menu>
        </Box>
        <Footer pad='medium'>
          <Button icon={<User />} />
        </Footer>
      </Sidebar>
    );
  }
}

const mapStateToProps = state => ({
  schemes: state.schemes.schemes,
  selectedSchemeIndex: state.schemes.selectedSchemeIndex,
})

const mapDispatchToProps = dispatch => bindActionCreators({
  selectScheme,
  addScheme,
}, dispatch)

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Sidemenu);
