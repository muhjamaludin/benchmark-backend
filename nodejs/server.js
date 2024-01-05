import http from 'http'

http.createServer(function (req, res) {
  res.write("Hello World")
  res.end()
}).listen(2000)
