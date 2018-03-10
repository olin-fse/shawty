const getConfig = function(env) {
  if (env !== 'production') {
    return {
      apiEndpoint: 'http://localhost:8080',
      staticEndpoint: 'http://localhost:9001',
    };
  } else {
    return {
      apiEndpoint: 'http://104.197.107.231:8080',
      staticEndpoint: 'http://35.226.201.59:9001',
    }
  }

};

module.exports = getConfig;
