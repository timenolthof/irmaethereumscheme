import axios from 'axios';
import parseXML from '../util/parse-xml';

import { FETCH_SCHEME, RECEIVE_SCHEME } from '../reducer/schemes';

const dataService = store => dispatch => action => {

  dispatch(action); // Note: Pass all actions through by default because this is middleware
  switch (action.type) {
    case FETCH_SCHEME:
      const { scheme } = action;
      fetchScheme(scheme, dispatch);
      break;

    default:
      break
  }

};

function fetchScheme(scheme, dispatch) {
  if (scheme.type === 'github') {
    return axios
      .get(`https://raw.githubusercontent.com/${scheme.owner}/${scheme.repository}/master/description.xml`)
      .then(response => response.data)
      .then(parseXML)
      .then(json => {
        dispatch({
          type: RECEIVE_SCHEME,
          metadata: {
            id: json.SchemeManager.Id[0],
            version: json.SchemeManager.$.version,
            contact: json.SchemeManager.Contact[0],
            name: json.SchemeManager.Name[0],
            description: json.SchemeManager.Description[0],
            url: json.SchemeManager.Url[0],
            keyshareAttribute: json.SchemeManager.KeyshareAttribute[0],
            keyshareServer: json.SchemeManager.KeyshareServer[0],
            keyshareWebsite: json.SchemeManager.KeyshareWebsite[0],
          },
          receivedAt: Date.now(),
        })
      });
  }
}

export default dataService;
