const express = require("express");

const app = express();

app.get("/", async (req, res) => {
  res.send({ hello: "world" });
});

const start = async () => {
    try {
      await app.listen(3001, () => {
        console.log(`server listen on port 3001`)
      })
    } catch (err) {
      process.exit(1)
    }
  }
  start()