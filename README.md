![Bbuild_and_test badge](https://github.com/andrewfitzy/2022-advent-of-code/actions/workflows/test.yml/badge.svg)

# 2022-advent-of-code

This repo contains the solutions for my path of [Advent of Code 2022](https://adventofcode.com/2022). I complete AoC to get familiar with a technology, its build tools and testing tools, it's kind of a mini-production type workflow I follow.

In this year I chose to use the following tools:
- [Kotlin v1.9.20](https://kotlinlang.org/docs/whatsnew1920.html). Language for this years AOC.
- [Gradle v8.1.1](https://docs.gradle.org/8.1.1/release-notes.html). Build tool for Kotlin projects.
- [Spotless v6.25.0](https://github.com/diffplug/spotless/releases/tag/gradle%2F6.25.0). Also using the Ktlint formatting configuration [v1.1.1](https://github.com/pinterest/ktlint/releases/tag/1.1.1).
- [Detekt v1.23.3](https://detekt.dev/docs/intro). Code analysis tool which is useful for finding bugs.
- [Kotlin JUnit](https://kotlinlang.org/api/latest/kotlin.test/). Unit test framework with a Kotlin flavour.

All development was completed using IntelliJ which is an awesome development environment.

## Setup
There is a small setup script that copies the projects pre-commit file to the `hooks` folder of the project's repo. This needs to be run after the project is cloned only. 
```bash
$ ./gradlew initProject  
```

## Testing
To run the tests, `cd` into this directory and then do:
```
$ go test -v ./...
```


## Committing
The pre-commit hook should kick-in, when it does it will run `spotless` and `detekt`
```bash
$ git add --all
$ git commit -a
```

## Progress
|                                                | Challenge                |                                                            Task 1                                                            |                                                            Task 2                                                            |
|:-----------------------------------------------|:-------------------------|:----------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------:|
| [Day 01](https://adventofcode.com/2022/day/1)  | Calorie Counting         |                                                                                                                              |                                                                                                                              |
| [Day 02](https://adventofcode.com/2022/day/2)  | Rock Paper Scissors      |                                                                                                                              |                                                                                                                              |
| [Day 03](https://adventofcode.com/2022/day/3)  | Rucksack Reorganization  |                                                                                                                              |                                                                                                                              |
| [Day 04](https://adventofcode.com/2022/day/4)  | Camp Cleanup             |                                                                                                                              |                                                                                                                              |
| [Day 05](https://adventofcode.com/2022/day/5)  | Supply Stacks            |                                                                                                                              |                                                                                                                              |
| [Day 06](https://adventofcode.com/2022/day/6)  | Tuning Trouble           |                                                                                                                              |                                                                                                                              |
| [Day 07](https://adventofcode.com/2022/day/7)  | No Space Left On Device  |                                                                                                                              |                                                                                                                              |
| [Day 08](https://adventofcode.com/2022/day/8)  | Treetop Tree House       |                                                                                                                              |                                                                                                                              |
| [Day 09](https://adventofcode.com/2022/day/9)  | Rope Bridge              |                                                                                                                              |                                                                                                                              |
| [Day 10](https://adventofcode.com/2022/day/10) | Cathode-Ray Tube         |                                                                                                                              |                                                                                                                              |
| [Day 11](https://adventofcode.com/2022/day/11) | Monkey in the Middle     |                                                                                                                              |                                                                                                                              |
| [Day 12](https://adventofcode.com/2022/day/12) | Hill Climbing Algorithm  |                                                                                                                              |                                                                                                                              |
| [Day 13](https://adventofcode.com/2022/day/13) | Distress Signal          |                                                                                                                              |                                                                                                                              |
| [Day 14](https://adventofcode.com/2022/day/14) | Regolith Reservoir       |                                                                                                                              |                                                                                                                              |
| [Day 15](https://adventofcode.com/2022/day/15) | Beacon Exclusion Zone    |                                                                                                                              |                                                                                                                              |
| [Day 16](https://adventofcode.com/2022/day/16) | Proboscidea Volcanium    |                                                                                                                              |                                                                                                                              |
| [Day 17](https://adventofcode.com/2022/day/17) | Pyroclastic Flow         |                                                                                                                              |                                                                                                                              |
| [Day 18](https://adventofcode.com/2022/day/18) | Boiling Boulders         |                                                                                                                              |                                                                                                                              |
| [Day 19](https://adventofcode.com/2022/day/19) | Not Enough Minerals      |                                                                                                                              |                                                                                                                              |
| [Day 20](https://adventofcode.com/2022/day/20) | Grove Positioning System |                                                                                                                              |                                                                                                                              |
| [Day 21](https://adventofcode.com/2022/day/21) | Monkey Math              |                                                                                                                              |                                                                                                                              |
| [Day 22](https://adventofcode.com/2022/day/22) | Monkey Map               |                                                                                                                              |                                                                                                                              |
| [Day 23](https://adventofcode.com/2022/day/23) | Unstable Diffusion       |                                                                                                                              |                                                                                                                              |
| [Day 24](https://adventofcode.com/2022/day/24) | Blizzard Basin           |                                                                                                                              |                                                                                                                              |
| [Day 25](https://adventofcode.com/2022/day/25) | Full of Hot Air          |                                                                                                                              |                                                                                                                              |