# go-life

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-life)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-life)

[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life).

## Installation

```
$ go get github.com/svetlana-rezvaya/go-life
```

## Usage

```
$ go-life -h | -help | --help
$ go-life [options]
```

Stdin: grid in the [plaintext](https://www.conwaylife.com/wiki/Plaintext) format.

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-outDelay DURATION` &mdash; delay between frames (e.g. `72h3m0.5s`; default: `100ms`).

## Output Example

```
........................O.............
......................O.O.............
............OO......OO............OO..
...........O...O....OO............OO..
OO........O.....O...OO................
OO........O...O.OO....O.O.............
..........O.....O.......O.............
...........O...O......................
............OO........................
.......................O..............
........................OO............
.......................OO.............
......................................
......................................
......................................
......................................
......................................
..............................O.O.....
...............................OO.....
...............................O......
```

## License

The MIT License (MIT)

Copyright &copy; 2020 svetlana-rezvaya
