const express = require('express');
const path = require('path');
const port = process.env.PORT || 9001;
const app = express();
const getConfig = require('./src/config');

const config = getConfig(process.env.NODE_ENV);

// Serve all static assets normally
app.use(express.static(__dirname + '/public'));

// Serve react app on requests to root
app.get('/', function (request, response) {
  response.sendFile(path.resolve(__dirname, 'public', 'index.html'));
});

app.get('/:code', function (request, response) {
  response.redirect(`${config.apiEndpoint}/${request.params.code}`);
});

app.listen(port);

