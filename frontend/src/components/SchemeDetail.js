import 'babel-polyfill'
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import Article from 'grommet/components/Article';
import Section from 'grommet/components/Section';
import Headline from 'grommet/components/Headline';;

// import { selectScheme } from '../reducer/schemes';

class SchemeDetail extends Component {
  static propTypes = {
    schemes: PropTypes.array.isRequired,
    selectedSchemeIndex: PropTypes.number,
  };

  render() {
    const scheme = this.props.schemes[this.props.selectedSchemeIndex];
    return scheme ? (
      <Article>
        <Section pad='large'
          justify='center'
          align='center'>
          <Headline margin='none'>
            {scheme.id}
          </Headline>
        </Section>
      </Article>
    ) : (
      <Article>
        <Section pad='large'
          justify='center'
          align='center'>
          <Headline margin='none'>
            Select a scheme
          </Headline>
        </Section>
      </Article>
    );
  }
}

const mapStateToProps = state => ({
  schemes: state.schemes.schemes,
  selectedSchemeIndex: state.schemes.selectedSchemeIndex,
})

const mapDispatchToProps = dispatch => bindActionCreators({
  // selectScheme,
}, dispatch)

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SchemeDetail);
