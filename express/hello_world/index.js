const express = require("express");

const app = express();

app.get("/", (req, res) => {
  res.send({ hello: "world" });
});

app.listen(3001, () => {
  console.log(`server listen on port 3001`);
});
