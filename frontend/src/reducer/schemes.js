export const SELECT_SCHEME = 'schemes/SELECT_SCHEME';
export const ADD_SCHEME = 'schemes/ADD_SCHEME';
export const FETCH_SCHEME = 'schemes/FETCH_SCHEME';
export const RECEIVE_SCHEME = 'schemes/RECEIVE_SCHEME';

const initialState = {
  schemes: [
    {
      id: 'pbdf',
      type: 'github',
      owner: 'privacybydesign',
      repository: 'pbdf-schememanager',
    },{
      id: 'test-ethereum2',
      type: 'ethereum',
      address: '0x0001',
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
      return {
        ...state,
        schemes: state.schemes.push(action.scheme),
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
