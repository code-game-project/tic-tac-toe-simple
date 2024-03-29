# Tic-Tac-Toe Simple
![CodeGame Version](https://img.shields.io/badge/CodeGame-v0.7-orange)
![CGE Version](https://img.shields.io/badge/CGE-v0.4-green)

A simple implementation of [tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) for [CodeGame](https://code-game.org).

This game is ideal for familiarizing yourself with CodeGame.

## Known instances

- `games.code-game.org/tic-tac-toe-simple`

## Usage

```sh
# Run on default port 8080
tic-tac-toe-simple

# Specify a custom port
tic-tac-toe-simple --port=5000

## Specify a custom port through an environment variable
CG_PORT=5000 tic-tac-toe-simple
```

### Running with Docker

Prerequisites:
- [Docker](https://docker.com/)

```sh
# Download image
docker pull codegameproject/tic-tac-toe-simple:0.3

# Run container
docker run -d --restart on-failure -p <port-on-host-machine>:8080 --name tic-tac-toe-simple codegameproject/tic-tac-toe-simple:0.3
```

## Event Flow

1. You receive a `start` event when a second player joins, which includes your sign ('x' or 'o').
2. You regularly receive a `board` event, which includes the current state of the board.
3. You receive a `turn` event, which includes the next sign to be placed.
4. When it is your turn you can send a `mark` event with the row and the columnn (zero based), which should be marked with your sign.
5. When the game is complete you will receive a `game_over` event, which includes which player wins and which fields form the winning row. Otherwise go to *3.*
6. You will receive an `invalid_action` event when you try to do something that's not allowed like marking an already marked field.

## Building

### Prerequisites

- [Go](https://go.dev) 1.18+

```sh
git clone https://github.com/code-game-project/tic-tac-toe-simple.git
cd tic-tac-toe-simple
codegame build
```

## License

Copyright (C) 2022-2023 Julian Hofmann

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
