# who-scored

`who-scored` is a command line application that finds and displays scores for NHL games.

## Usage

This guide assumes you have already [installed](#Install) the application.

### Help

```sh
who-scored --help
```

### Print scores for today's NHL games

#### Example input

```sh
who-scored today
```

#### Example output

```
+-----------+---------------------+----------------+-------+-----------+-------+
| GAME TYPE | START TIME          | AWAY TEAM      | SCORE | HOME TEAM | SCORE |
+-----------+---------------------+----------------+-------+-----------+-------+
| Playoffs  | 01 May 24 16:30 PDT | Golden Knights | 0     | Stars     | 0     |
| Playoffs  | 01 May 24 19:00 PDT | Kings          | 0     | Oilers    | 0     |
+-----------+---------------------+----------------+-------+-----------+-------+
```

---

### Print scores for yesterday's NHL games

#### Example input

```sh
who-scored yesterday
```

#### Example output

```
+-----------+---------------------+-------------+-------+------------+-------+
| GAME TYPE | START TIME          | AWAY TEAM   | SCORE | HOME TEAM  | SCORE |
+-----------+---------------------+-------------+-------+------------+-------+
| Playoffs  | 30 Apr 24 16:00 PDT | Maple Leafs | 2     | Bruins     | 1     |
| Playoffs  | 30 Apr 24 16:30 PDT | Islanders   | 3     | Hurricanes | 6     |
| Playoffs  | 30 Apr 24 18:30 PDT | Avalanche   | 6     | Jets       | 3     |
| Playoffs  | 30 Apr 24 19:00 PDT | Predators   | 2     | Canucks    | 1     |
+-----------+---------------------+-------------+-------+------------+-------+
```

### Print scores for games played on a single day

#### Example input

```sh
# Date must be in YYYY-MM-DD format
who-scored on 2023-03-16
```

#### Example output

```
+----------------+---------------------+--------------+-------+----------------+-------+
| GAME TYPE      | START TIME          | AWAY TEAM    | SCORE | HOME TEAM      | SCORE |
+----------------+---------------------+--------------+-------+----------------+-------+
| Regular Season | 16 Mar 23 16:00 PDT | Avalanche    | 5     | Senators       | 4     |
| Regular Season | 16 Mar 23 16:00 PDT | Canadiens    | 5     | Panthers       | 9     |
| Regular Season | 16 Mar 23 16:00 PDT | Penguins     | 2     | Rangers        | 4     |
| Regular Season | 16 Mar 23 16:00 PDT | Lightning    | 4     | Devils         | 3     |
| Regular Season | 16 Mar 23 17:00 PDT | Bruins       | 3     | Jets           | 0     |
| Regular Season | 16 Mar 23 17:00 PDT | Blackhawks   | 2     | Predators      | 1     |
| Regular Season | 16 Mar 23 18:00 PDT | Stars        | 1     | Oilers         | 4     |
| Regular Season | 16 Mar 23 19:00 PDT | Flames       | 7     | Golden Knights | 2     |
| Regular Season | 16 Mar 23 19:00 PDT | Canucks      | 2     | Coyotes        | 3     |
| Regular Season | 16 Mar 23 19:30 PDT | Blue Jackets | 1     | Kings          | 4     |
| Regular Season | 16 Mar 23 19:30 PDT | Kraken       | 2     | Sharks         | 1     |
+----------------+---------------------+--------------+-------+----------------+-------+
```

## Install

The program can either be installed from a release artifact or built locally.

### Install from release artifact

1. Go to the [Releases page](https://github.com/devkevbot/who-scored/releases), download, and extract the latest version for your system.
2. Add the executable file to `$PATH`.

### Building from source

1. Clone the repository:
   ```sh
    https://github.com/devkevbot/who-scored.git
    ```
2. From the repository directory, build and install the project: `go install .`
