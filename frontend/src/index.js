import React from 'react';
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import { ConnectedRouter } from 'react-router-redux'

import './index.css';
import '../node_modules/grommet-css'

import store, { history } from './store'
import App from './containers/App'
import registerServiceWorker from './registerServiceWorker';

const target = document.getElementById('root');

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <div>
        <App />
      </div>
    </ConnectedRouter>
  </Provider>,
  target
)
registerServiceWorker();
