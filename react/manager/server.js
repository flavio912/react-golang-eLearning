const express = require("express");
// const cors = require('cors');
const path = require("path");

const app = express();
const port = 4000;

app.use(express.json());

app.get("/*", function (req, res) {
  res.sendFile(path.join(__dirname, "build", "index.html"));
});

app.listen(port, () => console.log(`Listening on port: ${port}`));
