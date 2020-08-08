# ENEXT-CHALLENGE

[![Test Status](https://github.com/bgildson/enext-challenge/workflows/Test%20and%20Send%20Coverage%20Report/badge.svg)](https://github.com/bgildson/enext-challenge/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/bgildson/enext-challenge/badge.svg?branch=master)](https://coveralls.io/github/bgildson/enext-challenge?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgildson/enext-challenge)](https://goreportcard.com/report/github.com/bgildson/enext-challenge)

This repository contains the source code to satisfy the [enext challenge](CHALLENGE.md).

## How to run the challenge solution

To facilitate the command executions, the commands bellow will use Docker as runner.
_In case of need to run pieces of the solution, you should use [Golang 1.14](https://golang.org/dl/) with modules and install the dependencies._

### Task 1

The first task was to create the **Quake 3 log parser**, run the command bellow to execute the parser.

```sh
docker run --rm -it -v $(pwd):/app -w /app golang:1.14 go run ./cmd/parser/main.go -log=./games.log -out=./games.json
```

The parser will generate a `.json` with the parsed games and should looks like the bellow representation.

```json
{
  ...
  "2": {
    "id": "2",
    "total_kills": 11,
    "players": [
      "Isgalamido",
      "Mocinha"
    ],
    "kills": {
      "Isgalamido": -7,
      "Mocinha": 0
    }
  },
  ...
}
```

### Task 2

The second task was to create the **Game Report**, the report has two output types, one for **players results ranking grouped by game** and other for **players general results ranking**. run the command bellow to execute the games report.

_obs: the commands bellow depends of the task 1 command execution, because it uses the `games.json` parser output._

Report **players results ranking grouped by game**
```sh
docker run --rm -it -v $(pwd):/app -w /app golang:1.14 go run ./cmd/report/main.go -games-json-path=./games.json -general=false
```

The first report should looks the like bellow representation.

```json
Game 1                              Total Kills: 0
Position | Player                         | Points

Game 2                             Total Kills: 11
Position | Player                         | Points
       1 | Mocinha                        | 0
       2 | Isgalamido                     | -7
...
```

Report **players general results ranking**
```sh
docker run --rm -it -v $(pwd):/app -w /app golang:1.14 go run ./cmd/report/main.go -games-json-path=./games.json -general=true
```

The second report should looks the like bellow representation.

```json
General Ranking                  Total Kills: 1069
Position | Player                         | Points
       1 | Isgalamido                     | 138
       2 | Zeh                            | 120
       3 | Oootsimo                       | 108
...
```

### Task 3

The third task was to create the **api for games results**, the api was created using a Clean Architecture minimum implementation and using the output from _the parser_ as data source. The api has two endpoints **/games** to list the games and the **/games/{id}** to find the game by id.

_obs: the commands bellow depends of the task 1 command execution, because it uses the `games.json` parser output._

```sh
docker run --rm -it -v $(pwd):/app -p 8080:8080 -w /app golang:1.14 go run ./cmd/api/main.go -games-json-path=./games.json -port=8080
```

The api will run on http://localhost:8080 and will provide one endpoint for **[/games](http://localhost:8080/games)** and other for **[/games/{id}](http://localhost:8080/games/2)**.

## How to run the solution tests

All the code is covered by tests and to execute the tests use the command bellow.

```sh
docker run --rm -it -v $(pwd):/app -w /app golang:1.14 go test -v ./...
```
