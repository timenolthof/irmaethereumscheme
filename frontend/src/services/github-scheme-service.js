import axios from 'axios';
import parseXML from '../util/parse-xml';

export const fetchGithubScheme = (schemeOwner, schemeRepository) => {
  return axios
    .get(`https://raw.githubusercontent.com/${schemeOwner}/${schemeRepository}/master/description.xml`)
    .then(response => response.data)
    .then(parseXML)
    .then(json => {
      const metadata = {
        id: json.SchemeManager.Id[0],
        version: json.SchemeManager.$.version,
        contact: json.SchemeManager.Contact[0],
        name: json.SchemeManager.Name[0],
        description: json.SchemeManager.Description[0],
        url: json.SchemeManager.Url[0],
        keyshareattribute: json.SchemeManager.KeyshareAttribute[0],
        keyshareserver: json.SchemeManager.KeyshareServer[0],
        keysharewebsite: json.SchemeManager.KeyshareWebsite[0],
      };
      Object.keys(metadata.name).forEach(function(key, index) {
         metadata.name[key] = metadata.name[key][0];
      });
      Object.keys(metadata.description).forEach(function(key, index) {
         metadata.description[key] = metadata.description[key][0];
      });
      console.log("SSS", metadata);
      return {
        metadata,
      };
    }
  );
}
