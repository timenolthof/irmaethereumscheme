import axios from 'axios';
import parseXML from '../util/parse-xml';

export const fetchGithubScheme = (schemeOwner, schemeRepository) => {
  return axios
    .get(`https://raw.githubusercontent.com/${schemeOwner}/${schemeRepository}/master/description.xml`)
    .then(response => response.data)
    .then(parseXML)
    .then(json => ({
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
    }));
}
