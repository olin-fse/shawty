import request from 'superagent';

export const generateCode = (url, cb) => {
  request.post('/generate')
    .send({Url: url})
    .end((err, res) => {
      if (err) return alert(err);

      const payload = JSON.parse(res.text);
      cb(payload.Url);
    });
};
