import request from 'superagent';
import getConfig from '../../config';

const config = getConfig(process.env.NODE_ENV);

console.log(config);

export const generateCode = (url, singleUse, cb) => {
  request.post(`${config.apiEndpoint}/generate`)
    .withCredentials()
    .send({Url: url, SingleUse: singleUse})
    .end((err, res) => {
      if (err) return alert(err);

      const payload = JSON.parse(res.text);
      cb(payload.code);
    });
};
