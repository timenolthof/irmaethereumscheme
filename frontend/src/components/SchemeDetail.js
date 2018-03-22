import 'babel-polyfill'
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import Article from 'grommet/components/Article';
import Section from 'grommet/components/Section';
import Headline from 'grommet/components/Headline';

import List from 'grommet/components/List';
import ListItem from 'grommet/components/ListItem';

class SchemeDetail extends Component {
  static propTypes = {
    schemes: PropTypes.array.isRequired,
    selectedSchemeId: PropTypes.string,
  };

  render() {
    const scheme = this.props.schemes.find(s => s.id === this.props.selectedSchemeId);
    const lang = 'en';

    return scheme ? (
      <Article>
        <Section pad='large'
          justify='center'
          align='center'>
          <Headline margin='none'>
            {scheme.id}
          </Headline>
        </Section>
        {
          scheme.metadata ? (
            <div>
              <List>
                <ListItem justify='between'
                  separator='horizontal'>
                  <span>
                    Name
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.name[lang] }
                  </span>
                </ListItem>
                <ListItem justify='between'
                  separator='horizontal'>
                  <span>
                    Description
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.description[lang] }
                  </span>
                </ListItem>
                <ListItem justify='between'
                  separator='horizontal'>
                  <span>
                    Version
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.version }
                  </span>
                </ListItem>
                <ListItem justify='between'>
                  <span>
                    Contact
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.contact }
                  </span>
                </ListItem>
                <ListItem justify='between'>
                  <span>
                    Url
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.url }
                  </span>
                </ListItem>
                <ListItem justify='between'
                  separator='horizontal'>
                  <span>
                    KeyshareAttribute
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.keyshareattribute }
                  </span>
                </ListItem>
                <ListItem justify='between'>
                  <span>
                    KeyshareServer
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.keyshareserver }
                  </span>
                </ListItem>
                <ListItem justify='between'>
                  <span>
                    keyshareWebsite
                  </span>
                  <span className='secondary'>
                    { scheme.metadata.keysharewebsite }
                  </span>
                </ListItem>
              </List>
            </div>
          ) : (
            "Geen metadata"
          )
        }
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
  selectedSchemeId: state.schemes.selectedSchemeId,
})

const mapDispatchToProps = dispatch => bindActionCreators({
  // selectScheme,
}, dispatch)

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SchemeDetail);
