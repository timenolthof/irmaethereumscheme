import { parseString } from 'react-native-xml2js';

export default xml => {
  return new Promise(function(resolve, reject) {
    parseString(xml, function(err, result) {
      if(err) return reject(err);
      resolve(result);
    });
  });
};
