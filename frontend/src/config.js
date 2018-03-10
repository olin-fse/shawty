const getConfig = (env) => {
  if (env !== 'production') {
    return {
      apiEndpoint: 'http://localhost:8080',
      staticEndpoint: 'http://localhost:9001',
    };
  }
};

module.exports = getConfig;
