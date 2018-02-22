import 'babel-polyfill'
import React, { Component } from 'react';

import Sidebar from 'grommet/components/Sidebar';
import Header from 'grommet/components/Header';
import Title from 'grommet/components/Title';
import Box from 'grommet/components/Box';
import Menu from 'grommet/components/Menu';
import Anchor from 'grommet/components/Anchor';
import Button from 'grommet/components/Button';
import User from 'grommet/components/icons/base/User';
import Footer from 'grommet/components/Footer';

class Sidemenu extends Component {
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
            <Anchor href='#'
              className='active'>
              First
            </Anchor>
            <Anchor href='#'>
              Second
            </Anchor>
            <Anchor href='#'>
              Third
            </Anchor>
          </Menu>
        </Box>
        <Footer pad='medium'>
          <Button icon={<User />} />
        </Footer>
      </Sidebar>
    );
  }
}

export default Sidemenu;
