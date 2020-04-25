const express = require("express");
const path = require("path");

const app = express();
const port = 4000;

app.use(express.json());

app.use("/static", express.static(path.join(__dirname, "./build/static")));
app.use("/images", express.static(path.join(__dirname, "./build/images")));

app.get("/*", function (req, res) {
  res.sendFile(path.join(__dirname, "build", "index.html"));
});

app.listen(port, () => console.log(`Listening on port: ${port}`));
