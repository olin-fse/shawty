const getConfig = (env) => {
  if (env !== 'production') {
    return {
      apiEndpoint: 'localhost:8080'
    };
  }
};

export default getConfig;
