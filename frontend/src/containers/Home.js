import 'babel-polyfill'
import React, { Component } from 'react';

import GrommetApp from 'grommet/components/App';
import Split from 'grommet/components/Split';

import AppHeader from '../components/AppHeader';
import Sidemenu from '../components/Sidemenu';
import SchemeDetail from '../components/SchemeDetail';

class App extends Component {
  render() {
    return (
      <GrommetApp>
        <AppHeader />
        <Split flex='right' fixed={true}>
          <Sidemenu />
          <SchemeDetail />
        </Split>
      </GrommetApp>
    );
  }
}

export default App;
