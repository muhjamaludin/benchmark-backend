# Benchmark Backend Framework

## Prerequisite

- Go
- Python3
- Nodejs
- Npm
- Pnpm
- PM2

## Backend Framework

- Go Gin (Golang)
- Fastify (Nodejs)
- Fastapi (Python)

### How to run framework

```bash
# Go gin
$ cd golang/gin_hello_world
$ go run main.go

# Fastify
$ cd fastify/hello_world
$ pnpm install
$ pm2 start hello_world/index.js -i max

# Fastapi
$ cd fastapi
$ python3 -m venv .venv
$ source .venv/bin/activate
$ pip install -r requirements.txt
$ granian --interface asgi --workers <max cpu> --threads <max thread> --loop uvloop hello_world/main:app
```

## Tools benchmark

- [bombardier](https://github.com/codesenberg/bombardier)
- [autocannon](https://github.com/mcollina/autocannon)

## Result 1

```markdown
Server: Localhost
OS: Linux mint 21.1
CPU: 8
Threads: 16
CPU Model: Intel Core i7 8650u
Python: 3.10
Go: 1.21
Node: 20
```

### Autocannon

| Framework | Concurrent | Pipeline | Requests (k) | Time (s) | Size (MB) |
| --------- | ---------- | -------- | ------------ | -------- | --------- |
| Fastify   | 500        | 500      | 1198         | 12.69    | 178       |
|           |            | 1000     | 1842         | 21       | 252       |
| Go Gin    |            | 500      | 622          | 10.43    | 54        |
|           |            | 1000     | 870          | 10.61    | 53.6      |
| Granian   |            | 500      | 593          | 10.72    | 48.7      |
|           |            | 1000     | 699          | 10.47    | 28.2      |

### Bombardier

| Framework | Concurrent | requests | Avg (k) | Avg (ms) | Max (k) | Max (ms) |
| --------- | ---------- | -------- | ------- | -------- | ------- | -------- |
| Fastify   | 500        | 500      | 4076    | 105.17   | 7865    | 155.18   |
|           |            | 1000     | 7563    | 54.40    | 13189   | 137.60   |
| Go Gin    |            | 500      | 6603    | 52.53    | 14409   | 92.92    |
|           |            | 1000     | 17078   | 35.51    | 27681   | 99.50    |
| Granian   |            | 500      | 5449    | 81.46    | 10037   | 142.45   |
|           |            | 1000     | 5994    | 92.16    | 15764   | 227.62   |

## Result 2

```markdown
Server: Localhost
OS: WSL, Debian 12
CPU: 8
Threads: 16
CPU Model: Intel Core i7 8650u
Python: 3.11
Go: 1.21
Node: 20
```

### Autocannon 2

```bash
autocannon -c <concurrent> -p <max request> <url>
```

| Framework | Concurrent | Pipeline | Requests (k) | Time (s) | Size (MB) |
| --------- | ---------- | -------- | ------------ | -------- | --------- |
| Fastify   | 500        | 500      | 1017         | 20.71    | 144       |
|           |            | 1000     | 1260         | 22.2     | 143       |
| Go Gin    |            | 500      | 275          | 11.16    | 3.68      |
|           |            | 1000     | 523          | 11.13    | 3.27      |
| Granian   |            | 500      | 371          | 11.64    | 17.1      |
|           |            | 1000     | 550          | 12.51    | 7.09      |

### Bombardier 2

```bash
bombardier -c <concurrent> -n <max request> <url>
```

| Framework | Concurrent | requests | Avg (k) | Avg (ms) | Max (k) | Max (ms) |
| --------- | ---------- | -------- | ------- | -------- | ------- | -------- |
| Fastify   | 500        | 500      | 4702    | 69.13    | 13073   | 151.87   |
|           |            | 1000     | 4967    | 80.04    | 14929   | 264.46   |
| Go Gin    |            | 500      | 3647    | 138.93   | 12046   | 220.11   |
|           |            | 1000     | 9161    | 38.83    | 19411   | 141.28   |
| Granian   |            | 500      | 7849    | 72.57    | 19314   | 101.84   |
|           |            | 1000     | 13308   | 47.20    | 17696   | 122.63   |
