import 'babel-polyfill'
import React, { Component } from 'react';
// import NetworkStatus from 'react-web3-network-status'

import Header from 'grommet/components/Header';
import Title from 'grommet/components/Title';
import Box from 'grommet/components/Box';
import Menu from 'grommet/components/Menu';
import Actions from 'grommet/components/icons/base/Actions';
import Anchor from 'grommet/components/Anchor';

class AppHeader extends Component {
  render() {
    return (
      <Header pad='small' colorIndex='accent-2-a'>
        <Title>
          IRMA Scheme Manager GUI
        </Title>
        <Box flex={true}
          justify='end'
          direction='row'
          responsive={false}>

          {/* <NetworkStatus />*/}
          <Menu icon={<Actions />}
            dropAlign={{"right": "right"}}>
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
      </Header>
    );
  }
}

export default AppHeader;
