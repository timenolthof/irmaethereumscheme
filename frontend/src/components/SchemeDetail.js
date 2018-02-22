import 'babel-polyfill'
import React, { Component } from 'react';

import Article from 'grommet/components/Article';
import Section from 'grommet/components/Section';
import Headline from 'grommet/components/Headline';;

class SchemeDetail extends Component {
  render() {
    return (
      <Article>
        <Section pad='large'
          justify='center'
          align='center'>
          <Headline margin='none'>
            Section 1
          </Headline>
        </Section>
      </Article>
    );
  }
}

export default SchemeDetail;
