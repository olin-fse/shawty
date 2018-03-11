const getConfig = function(env) {
  if (env !== 'production') {
    return {
      apiEndpoint: 'http://localhost:8080',
      staticEndpoint: 'http://localhost:9001',
    };
  } else {
    return {
      apiEndpoint: 'http://shwt.cf',
      staticEndpoint: 'http://shawty.cf',
    }
  }

};

module.exports = getConfig;
