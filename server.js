const express = require('express');
// const cors = require('cors');
const path = require('path');

const app = express();
const port = 80;

app.use(express.json());
// app.use(cors({ origin: true, credentials: true }));
app.use('/static', express.static(path.join(__dirname, '/build/static')));
app.use('/images', express.static(path.join(__dirname, '/build/images')));

app.get('/*', function(req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.listen(port, () => console.log(`Listening on port: ${port}`));
