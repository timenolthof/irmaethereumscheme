export const SELECT_SCHEME = 'schemes/SELECT_SCHEME';
export const ADD_SCHEME = 'schemes/ADD_SCHEME';

const initialState = {
  schemes: [
    {
      id: 'test-ethereum1',
      type: 'ethereum',
      address: '0x0000',
    },{
      id: 'test-ethereum2',
      type: 'ethereum',
      address: '0x0001',
    },
  ],
  selectedSchemeIndex: null,
}

export default (state = initialState, action) => {
  switch (action.type) {
    case SELECT_SCHEME:
      return {
        ...state,
        selectedSchemeIndex: action.index,
      }

    case ADD_SCHEME:
      return {
        ...state,
        schemes: state.schemes.push(action.scheme),
      }

    default:
      return state;
  }
}

export const selectScheme = (index) => {
  return dispatch => {
    dispatch({
      type: SELECT_SCHEME,
      index
    })
  }
}

export const addScheme = (scheme) => {
  return dispatch => {
    dispatch({
      type: ADD_SCHEME,
      scheme
    })
  }
}
