![Bbuild_and_test badge](https://github.com/andrewfitzy/2022-advent-of-code/actions/workflows/test.yml/badge.svg)

# 2022-advent-of-code

This repo contains the solutions for my path of [Advent of Code 2022](https://adventofcode.com/2022). I complete AoC to get familiar with a technology, its build tools and testing tools, it's kind of a mini-production type workflow I follow.

In this year I chose to use the following tools:
- [Go v1.19](https://go.dev/doc/devel/release#go1.19) - Programming language for this years solutions


All development was completed using [Visual Studio Code](https://code.visualstudio.com) which is an ok text editor.

## Setup
After cloning the repo, perform the setup operations:
```bash
$ pre-commit install
```

## Testing
To run the tests, `cd` into this directory and then do:
```bash
$ go test -v ./...
```


## Committing
The pre-commit hook should kick-in, when it does it will run `go-mod-tidy` and `go-fmt`
```bash
$ git add --all
$ git commit -a
```

## Progress
|                                                | Challenge                |                                                            Task 1                                                            |                                                            Task 2                                                            |
|:-----------------------------------------------|:-------------------------|:----------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------:|
| [Day 01](https://adventofcode.com/2022/day/1)  | Calorie Counting         | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_01/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_01/task_02_main.go)|
| [Day 02](https://adventofcode.com/2022/day/2)  | Rock Paper Scissors      | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_02/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_02/task_02_main.go)|
| [Day 03](https://adventofcode.com/2022/day/3)  | Rucksack Reorganization  | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_03/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_03/task_02_main.go)|
| [Day 04](https://adventofcode.com/2022/day/4)  | Camp Cleanup             | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_04/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_04/task_02_main.go)|
| [Day 05](https://adventofcode.com/2022/day/5)  | Supply Stacks            | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_05/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_05/task_02_main.go)|
| [Day 06](https://adventofcode.com/2022/day/6)  | Tuning Trouble           | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_06/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_06/task_02_main.go)|
| [Day 07](https://adventofcode.com/2022/day/7)  | No Space Left On Device  | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_07/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_07/task_02_main.go)|
| [Day 08](https://adventofcode.com/2022/day/8)  | Treetop Tree House       | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_08/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_08/task_02_main.go)|
| [Day 09](https://adventofcode.com/2022/day/9)  | Rope Bridge              | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_09/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_09/task_02_main.go)|
| [Day 10](https://adventofcode.com/2022/day/10) | Cathode-Ray Tube         | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_10/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_10/task_02_main.go)|
| [Day 11](https://adventofcode.com/2022/day/11) | Monkey in the Middle     | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_11/task_01_main.go) |[🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_11/task_02_main.go)|
| [Day 12](https://adventofcode.com/2022/day/12) | Hill Climbing Algorithm  |                                                                                                                              |                                                                                                                              |
| [Day 13](https://adventofcode.com/2022/day/13) | Distress Signal          |                                                                                                                              |                                                                                                                              |
| [Day 14](https://adventofcode.com/2022/day/14) | Regolith Reservoir       |                                                                                                                              |                                                                                                                              |
| [Day 15](https://adventofcode.com/2022/day/15) | Beacon Exclusion Zone    |                                                                                                                              |                                                                                                                              |
| [Day 16](https://adventofcode.com/2022/day/16) | Proboscidea Volcanium    |                                                                                                                              |                                                                                                                              |
| [Day 17](https://adventofcode.com/2022/day/17) | Pyroclastic Flow         |                                                                                                                              |                                                                                                                              |
| [Day 18](https://adventofcode.com/2022/day/18) | Boiling Boulders         |                                                                                                                              |                                                                                                                              |
| [Day 19](https://adventofcode.com/2022/day/19) | Not Enough Minerals      |                                                                                                                              |                                                                                                                              |
| [Day 20](https://adventofcode.com/2022/day/20) | Grove Positioning System |                                                                                                                              |                                                                                                                              |
| [Day 21](https://adventofcode.com/2022/day/21) | Monkey Math              | [🌟](https://github.com/andrewfitzy/2022-advent-of-code/blob/main/day_21/task_01_main.go) |                                                                                                                              |
| [Day 22](https://adventofcode.com/2022/day/22) | Monkey Map               |                                                                                                                              |                                                                                                                              |
| [Day 23](https://adventofcode.com/2022/day/23) | Unstable Diffusion       |                                                                                                                              |                                                                                                                              |
| [Day 24](https://adventofcode.com/2022/day/24) | Blizzard Basin           |                                                                                                                              |                                                                                                                              |
| [Day 25](https://adventofcode.com/2022/day/25) | Full of Hot Air          |                                                                                                                              |                                                                                                                              |