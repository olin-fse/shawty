import request from 'superagent';
import getConfig from '../../config';

const config = getConfig(process.env.NODE_ENV);

export const generateCode = (url, cb) => {
  request.post(`${config.apiEndpoint}/generate`)
    .withCredentials()
    .send({Url: url})
    .end((err, res) => {
      if (err) return alert(err);

      const payload = JSON.parse(res.text);
      cb(payload.Code);
    });
};
