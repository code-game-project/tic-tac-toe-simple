# Tic-Tac-Toe Simple
![CodeGame Version](https://img.shields.io/badge/CodeGame-v0.6-orange)
![CGE Version](https://img.shields.io/badge/CGE-v0.3-green)

A simple implementation of [tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) for [CodeGame](https://github.com/code-game-project).

This game is ideal for familiarizing yourself with CodeGame.

## Usage

```sh
# Run on default port 80
tic-tac-toe-simple

# Specify a custom port
tic-tac-toe-simple --port 8080

## Specify a custom port through an environment variable
CG_PORT=8080 pong
```

## Building

### Prerequisites

- [Go](https://go.dev) 1.18+

```sh
git clone https://github.com/code-game-project/tic-tac-toe-simple.git
cd tic-tac-toe-simple
go build .
```

## License

Copyright (C) 2022 Julian Hofmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
