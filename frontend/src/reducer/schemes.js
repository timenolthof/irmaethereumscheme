export const SELECT_SCHEME = 'schemes/SELECT_SCHEME';
export const ADD_SCHEME = 'schemes/ADD_SCHEME';
export const FETCH_SCHEME = 'schemes/FETCH_SCHEME';
export const RECEIVE_SCHEME = 'schemes/RECEIVE_SCHEME';
export const DEPLOY_NEW_SCHEME = 'schemes/DEPLOY_NEW_SCHEME';

const initialState = {
  schemes: [
    {
      id: 'pbdf',
      type: 'github',
      owner: 'privacybydesign',
      repository: 'pbdf-schememanager',
    },{
      id: 'ethereum-irma-scheme',
      type: 'ethereum',
      address: '0x71e13ea1d4bb434a3846c28a2a329f3ae83cf7c9',
    },
  ],
  selectedSchemeId: null,
}

export default (state = initialState, action) => {
  switch (action.type) {
    case SELECT_SCHEME:
      return {
        ...state,
        selectedSchemeId: action.schemeId,
      }

    case ADD_SCHEME:
      console.log(state);
      const newSchemes = state.schemes.slice();
      newSchemes.push(action.scheme);
      return {
        ...state,
        schemes: newSchemes,
      }

    case RECEIVE_SCHEME:
      return {
        ...state,
        schemes: state.schemes.map(scheme => {
          if (scheme.id === action.metadata.id) {
            scheme.metadata = action.metadata;
          }
          return scheme;
        }),
      }

    default:
      return state;
  }
}

export const selectScheme = (scheme) => {
  return dispatch => {
    dispatch({
      type: SELECT_SCHEME,
      schemeId: scheme.id,
    });
    dispatch({
      type: FETCH_SCHEME,
      scheme,
    });
  }
}

export const addScheme = (scheme) => {
  return dispatch => {
    dispatch({
      type: ADD_SCHEME,
      scheme
    });
  }
}

export const deployNewScheme = (schemeId, metadata) => {
  return dispatch => {
    dispatch({
      type: DEPLOY_NEW_SCHEME,
      schemeId,
      metadata,
    });
  }
}
