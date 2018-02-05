const getConfig = (env) => {
  if (env !== 'production') {
    return {
      apiEndpoint: 'http://localhost:8080'
    };
  }
};

export default getConfig;
