import {
  FETCH_SCHEME,
  RECEIVE_SCHEME,
  DEPLOY_NEW_SCHEME,
  ADD_SCHEME,
} from '../reducer/schemes';

import { deployNewScheme, fetchEthereumScheme } from './ethereum-scheme-service';
import { fetchGithubScheme } from './github-scheme-service';

const dataService = store => dispatch => action => {

  dispatch(action); // Note: Pass all actions through by default because this is middleware
  switch (action.type) {
    case FETCH_SCHEME:
      const { scheme } = action;
      if (scheme.type === 'github') {
        fetchGithubScheme(scheme.owner, scheme.repository)
          .then((scheme) => {
            dispatch({
              type: RECEIVE_SCHEME,
              metadata: scheme.metadata,
              receivedAt: Date.now(),
            });
          });
      } else {
        fetchEthereumScheme(scheme.address);
      }
      break;

    case DEPLOY_NEW_SCHEME:
      const { schemeId, metadata } = action;
      deployNewScheme(schemeId, metadata)
        .then(deployedContract => {
          dispatch({
            type: ADD_SCHEME,
            scheme: {
              id: schemeId,
              type: 'ethereum',
              address: deployedContract.address,
            },
          });
        });
      break;

    default:
      break
  }

};

export default dataService;
