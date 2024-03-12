# Old Result Benchmark

## Backend Framework

- Expressjs
- Fastify
- Fastapi

Tools benchmark : `npm loadtest`

## Server

```markdown
Server: Google Cloud Shell
OS: Debian 11
CPU: 4
CPU Model: AMD Epyc 7B12
```

## Hello World

```markdown
Runner:
Node: v18.18.2 (nvm 0.35.2)
Python: 3.10.8
Uvicorn: 0.23.2
Gunicorn: 21.2.0 (worker 4, uvicorn worker, async)
Go: 1.21.3
```

| Framework | Runner    | user   | requests  | time taken (s) | request per second | mean latency (ms) |
| --------- | --------- | ------ | --------- | -------------- | ------------------ | ----------------- |
| Express   | Node      | 25     | 100       | 0.097          | 1031               | 34.1              |
|           |           |        | 1_000     | 0.526          | 1901               | 24.2              |
|           |           |        | 10_000    | 3.55           | 2817               | 17.1              |
|           |           |        | 100_000   | 14.7           | 3283               | 14.7              |
|           |           |        | 1_000_000 | 300.245        | 3331               | 14.5              |
|           |           |        | 2_000_000 | 604.137        | 3311               | 14.6              |
|           |           | 100    | 100       | 0.136          | 735                | 82.9              |
|           |           |        | 1_000     | 0.708          | 1412               | 126.4             |
|           |           |        | 10_000    | 4.838          | 2067               | 95.1              |
|           |           |        | 100_000   | 40.34          | 2479               | 80                |
|           |           |        | 1_000_000 | 409.096        | 2444               | 81.3              |
|           |           |        | 2_000_000 | 793.213        | 2521               | 78.8              |
|           |           | 1000   | 1000      | 0.858          | 1166               | 533.1             |
|           |           |        | 10_000    | 8.567          | 1167               | 1040.8            |
|           |           |        | 100_000   | 52.215         | 1915               | 894               |
|           |           |        | 1_000_000 | 439.606        | 2275               | 875.8             |
|           |           | 10_000 | 10_000    | 20.068         | 498                | 4655.9            |
|           |           |        | 100_000   | 203.161        | 492                | 37821.1           |
| Fastify   |           | 25     | 100       | 0.083          | 1205               | 29.6              |
|           |           |        | 1_000     | 0.363          | 2755               | 16.7              |
|           |           |        | 10_000    | 2.374          | 4212               | 11.1              |
|           |           |        | 100_000   | 17.915         | 5582               | 8.4               |
|           |           |        | 1_000_000 | 172.149        | 5809               | 8.1               |
|           |           |        | 2_000_000 | 347.652        | 5753               | 8.2               |
|           |           | 500    | 500       | 0.195          | 2564               | 104.5             |
|           |           |        | 50_000    | 9.696          | 5157               | 284.7             |
| FastApi   | Uvicorn   |        | 100       | 0.14           | 714                | 51.5              |
|           |           |        | 1_000     | 1.104          | 906                | 52.7              |
|           |           |        | 10_000    | 10.365         | 965                | 51.1              |
|           |           |        | 100_000   | 100.788        | 992                | 49.9              |
|           | Gunicorn  |        | 100       | 0.065          | 1538               | 21.9              |
|           |           |        | 1_000     | 0.403          | 2481               | 18.4              |
|           |           |        | 10_000    | 2.745          | 3643               | 13.1              |
|           |           |        | 100_000   | 22.651         | 4415               | 10.8              |
|           |           |        | 1_000_000 | 224.235        | 4460               | 10.7              |
|           |           |        | 2_000_000 | 445.594        | 4488               | 10.6              |
|           |           | 100    | 100       | 0.145          | 690                | 113.3             |
|           |           |        | 1_000     | 0.622          | 1608               | 112.7             |
|           |           |        | 10_000    | 5.016          | 1994               | 98                |
|           |           |        | 100_000   | 42.88          | 2332               | 84.9              |
|           |           |        | 1_000_000 | 411.378        | 2431               | 81.6              |
|           |           |        | 2_000_000 | 831.197        | 2406               | 82.6              |
|           |           | 1000   | 1000      | 0.611          | 1637               | 423.1             |
|           |           |        | 10_000    | 5.025          | 1990               | 896.3             |
|           |           |        | 100_000   | 43.629         | 2292               | 856.7             |
|           |           |        | 1_000_000 | 430.567        | 2323               | 859.7             |
|           |           | 10_000 | 10_000    | 3.416          | 2927               | 2051.5            |
|           |           |        | 100_000   | 86.201         | 1160               | 16734.9           |
| net/http  | Go        | 25     | 100       | 0.084          | 1190               | 19                |
|           |           |        | 1_000     | 0.313          | 3195               | 13.8              |
|           |           |        | 10_000    | 2.125          | 4706               | 9.9               |
|           |           |        | 100_000   | 16.886         | 5922               | 7.9               |
|           |           |        | 1_000_000 | 162.902        | 6139               | 7.6               |
|           |           |        | 2_000_000 | 329.544        | 6069               | 7.7               |
| Gin       | Go        |        | 100       | 0.056          | 1786               | 19.8              |
|           |           |        | 1_000     | 0.331          | 3021               | 15                |
|           |           |        | 10_000    | 2.65           | 3774               | 12.6              |
|           |           |        | 100_000   | 23.449         | 4265               | 11.2              |
|           |           |        | 1_000_000 | 255.255        | 3918               | 12.2              |
|           |           |        | 2_000_000 | 404.888        | 4940               | 9.5               |
|           |           | 500    | 500       | 0.199          | 2513               | 146.4             |
|           |           |        | 50_000    | 5.111          | 9783               | 195.6             |
| Express   | PM2       | 25     | 100       | 0.167          | 599                | 56                |
|           |           |        | 1_000     | 0.584          | 1712               | 25                |
|           |           |        | 10_000    | 2.929          | 3414               | 13.9              |
|           |           |        | 100_000   | 23.956         | 4174               | 11.5              |
|           |           |        | 1_000_000 | 225.632        | 4432               | 10.8              |
| Fastify   | PM2       |        | 100       | 0.12           | 833                | 41.9              |
|           |           |        | 1_000     | 0.336          | 2976               | 14.7              |
|           |           |        | 10_000    | 1.369          | 7305               | 6.2               |
|           |           |        | 100_000   | 8.265          | 12099              | 3.6               |
|           |           |        | 1_000_000 | 77.718         | 12867              | 3.4               |
| Fastify   | PM2 (Bun) |        | 100       | 0.124          | 806                | 40.8              |
|           |           |        | 1_000     | 0.284          | 3521               | 12.5              |
|           |           |        | 10_000    | 1.249          | 8006               | 5.6               |
|           |           |        | 100_000   | 7.571          | 13208              | 3.3               |
|           |           |        | 1_000_000 | 70.895         | 14105              | 3.1               |
|           |           |        | 2_000_000 | 144.955        | 13797              | 3.1               |

## Hello World Async

```markdown
Runner:
Node: v18.18.2 (nvm 0.35.2)
Python: 3.10.8
Uvicorn: 0.23.2
Gunicorn: 21.2.0 (worker 4, uvicorn worker)
```

| Framework | Runner   | user   | requests  | time taken (s) | request per second | mean latency (ms) |
| --------- | -------- | ------ | --------- | -------------- | ------------------ | ----------------- |
| Express   | Node     | 25     | 100       | 0.107          | 935                | 36.8              |
|           |          |        | 1_000     | 0.687          | 1456               | 32.3              |
|           |          |        | 10_000    | 4.816          | 2076               | 23.4              |
|           |          |        | 100_000   | 41.906         | 2386               | 20.4              |
|           |          |        | 1_000_000 | 403.972        | 2475               | 19.7              |
|           |          |        | 2_000_000 | 790.485        | 2530               | 19.3              |
|           |          | 100    | 100       | 0.108          | 926                | 64.4              |
|           |          |        | 1_000     | 0.725          | 1379               | 131.8             |
|           |          |        | 10_000    | 4.925          | 2030               | 96.7              |
|           |          |        | 100_000   | 42.403         | 2358               | 84.2              |
|           |          |        | 1_000_000 | 407.285        | 2455               | 80.9              |
|           |          |        | 2_000_000 | 810.539        | 2467               | 80.5              |
|           |          | 1000   | 1_000     | 0.826          | 1211               | 500.5             |
|           |          |        | 10_000    | 8.476          | 1180               | 1042.7            |
|           |          |        | 100_000   | 50.633         | 1975               | 882.2             |
|           |          |        | 1_000_000 | 447.499        | 2235               | 859.9             |
|           |          | 10_000 | 10_000    | 17.293         | 578                | 4464.7            |
|           |          |        | 100_000   | 196.681        | 508                | 38515.2           |
| Fastify   |          | 25     | 100       | 0.093          | 1075               | 34.1              |
|           |          |        | 1_000     | 0.439          | 2278               | 19.6              |
|           |          |        | 10_000    | 2.846          | 3514               | 13.5              |
|           |          |        | 100_000   | 21.556         | 4639               | 10.2              |
|           |          |        | 1_000_000 | 205.323        | 4870               | 9.7               |
|           |          |        | 2_000_000 | 409.502        | 4884               | 9.7               |
|           |          | 100    | 100       | 0.118          | 847                | 71.2              |
|           |          |        | 1_000     | 0.558          | 1792               | 94.2              |
|           |          |        | 10_000    | 3.096          | 3230               | 60.4              |
|           |          |        | 100_000   | 23.241         | 4303               | 45.7              |
|           |          |        | 1_000_000 | 220.919        | 4527               | 43.6              |
|           |          |        | 2_000_000 | 448.308        | 4461               | 44.3              |
|           |          | 1_000  | 1_000     | 0.421          | 2375               | 267.9             |
|           |          |        | 10_000    | 4.163          | 2402               | 602.6             |
|           |          |        | 100_000   | 26.174         | 3821               | 509.4             |
|           |          |        | 1_000_000 | 246.046        | 4064               | 490.5             |
|           |          | 10_000 | 10_000    | 9.708          | 1030               | 3143.1            |
|           |          |        | 100_000   | 193.395        | 517                | 38086.1           |
| FastApi   | Uvicorn  |        | 100       | 0.084          | 1190               | 28.8              |
|           |          |        | 1_000     | 0.565          | 1770               | 26.1              |
|           |          |        | 10_000    | 3.503          | 2855               | 16.8              |
|           |          |        | 100_000   | 32.412         | 3085               | 15.7              |
|           |          |        | 1_000_000 | 318.094        | 3144               | 15.4              |
|           |          |        | 2_000_000 | 632.7          | 3161               | 15.3              |
|           | Gunicorn |        | 100       | 0.082          | 1220               | 25.1              |
|           |          |        | 1_000     | 0.452          | 2212               | 20                |
|           |          |        | 10_000    | 3.281          | 3048               | 15.7              |
|           |          |        | 100_000   | 26.903         | 3717               | 12.8              |
|           |          |        | 1_000_000 | 287.524        | 3478               | 13.8              |
|           |          |        | 2_000_000 | 584.986        | 3419               | 14.1              |
|           |          | 100    | 100       | 0.093          | 1075               | 57.6              |
|           |          |        | 1000      | 0.498          | 2008               | 89.5              |
|           |          |        | 10_000    | 3.61           | 2770               | 70.1              |
|           |          |        | 100_000   | 29.692         | 3368               | 58.3              |
|           |          |        | 1_000_000 | 286.755        | 3487               | 56.7              |
|           |          |        | 2_000_000 | 570.318        | 3507               | 56.5              |
|           |          | 1000   | 1_000     | 0.491          | 2037               | 331.5             |
|           |          |        | 10_000    | 3.833          | 2609               | 698.3             |
|           |          |        | 100_000   | 30.659         | 3262               | 607.2             |
|           |          |        | 1_000_000 | 306.196        | 3266               | 610.2             |
|           |          | 10_000 | 10_000    | 2.889          | 3461               | 1985.9            |
|           |          |        | 100_000   | 107.266        | 932                | 21059.2           |

### Hello

| Runner           | Concurrency | Max request | Total Errors | Total time | Mean latency | Effective rps |
| ---------------- | ----------- | ----------- | ------------ | ---------- | ------------ | ------------- |
| Nodejs (Native)  | 100         | 1_000       | 0            | 0.311      | 86.9         | 3215          |
|                  |             | 10_000      | 0            | 1.679      | 63.9         | 5956          |
|                  |             | 100_000     | 0            | 12.896     | 50.8         | 7754          |
|                  |             | 1_000_000   | 0            | 151.168    | 59.9         | 6615          |
|                  | 1_000       | 1_000       | 0            | 0.384      | 275.6        | 2604          |
|                  |             | 10_000      | 0            | 3.752      | 649.5        | 2665          |
|                  |             | 100_000     | 222          | 28.259     | 665.6        | 3539          |
| Nodejs (Express) | 100         | 1_000       | 0            | 0.453      | 134.8        | 2208          |
|                  |             | 10_000      | 0            | 2.93       | 113.1        | 3413          |
|                  |             | 100_000     | 0            | 23.854     | 94.5         | 4192          |
|                  |             | 1_000_000   | 0            | 284.65     | 113.3        | 3513          |
|                  | 1_000       | 1_000       | 0            | 1.335      | 517.3        | 749           |
|                  |             | 10_000      | 0            | 7.653      | 1082         | 1307          |
|                  |             | 100_000     | 468          | 55.987     | 1165.8       | 1786          |
| Bun (express.js) | 100         | 1_000       | 0            | 0.3        | 93.2         | 3333          |
|                  |             | 10_000      | 0            | 1.564      | 59.1         | 6394          |
|                  |             | 100_000     | 0            | 11.207     | 44.1         | 8923          |
|                  |             | 1_000_000   | 0            | 124.879    | 49.4         | 8008          |
|                  | 1_000       | 1_000       | 0            | 0.405      | 306          | 2469          |
|                  |             | 10_000      | 0            | 2.451      | 660.9        | 4080          |
|                  |             | 100_000     | 0            | 13.099     | 459.2        | 7634          |
|                  |             | 1_000_000   | 270          | 167.89     | 504.4        | 5956          |
| Bun (Native)     | 100         | 1_000       | 0            | 0.449      | 139          | 2227          |
|                  |             | 10_000      | 0            | 3.105      | 119.4        | 3221          |
|                  |             | 100_000     | 0            | 27.543     | 109.3        | 3631          |
|                  | 1_000       | 1_000       | 0            | 0.606      | 415.8        | 1650          |
|                  |             | 10_000      | 0            | 4.628      | 1111.7       | 2161          |
|                  |             | 100_000     | 0            | 33.621     | 1068.9       | 2974          |
|                  |             | 1_000_000   | 580          | 400.162    | 1378.3       | 2499          |
| Bun (express.ts) | 100         | 1_000       | 0            | 0.31       | 97.9         | 3226          |
|                  |             | 10_000      | 0            | 1.509      | 57.1         | 6627          |
|                  |             | 100_000     | 0            | 11.256     | 44.3         | 8884          |
|                  |             | 1_000_000   | 0            | 124.101    | 49.1         | 8058          |
|                  | 1000        | 1_000       | 0            | 0.399      | 290.6        | 2506          |
|                  |             | 10_000      | 0            | 2.581      | 727.1        | 3874          |
|                  |             | 100_000     | 0            | 12.428     | 442.7        | 8046          |
|                  |             | 1_000_000   | 0            | 151.034    | 509          | 6621          |
| Go (Native)      | 100         | 1_000       | 0            | 0.394      | 130.3        | 2538          |
|                  |             | 10_000      | 0            | 1.335      | 49.1         | 7491          |
|                  |             | 100_000     | 0            | 8.397      | 32.6         | 11909         |
|                  |             | 1_000_000   | 0            | 85.757     | 33.6         | 11661         |
|                  | 1_000       | 1_000       | 0            | 0.38       | 288.7        | 2632          |
|                  |             | 10_000      | 0            | 1.641      | 561.9        | 6094          |
|                  |             | 100_000     | 0            | 9.539      | 372.2        | 10483         |
|                  |             | 1_000_000   | 0            | 94.766     | 377.1        | 10552         |