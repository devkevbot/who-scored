# who-scored
CLI to pretty-print the scores for NHL games in a given date range.

## Install

### Releases

Go to the [Releases page](https://github.com/devkevbot/who-scored/releases), download, and extract the latest version for your system.

### Building from source

1. Build the project: `go build`
2. Run the binary: `./who-scored[.exe]`

---

## Usage

This guide assumes you have added the binary to `$PATH`.

### Help

```sh
who-scored --help
```

### Print scores for today's NHL games

```sh
who-scored today
```

#### Example output

```
+---------------------+--------------------------+-------+---------------------------+-------+-------------+
| START TIME          | AWAY TEAM (W-L)          | SCORE | HOME TEAM (W-L)           | SCORE | STATUS      |
+---------------------+--------------------------+-------+---------------------------+-------+-------------+
| 19 Apr 23 18:30 PDT | Minnesota Wild (1-0)     |     3 | Dallas Stars (0-1)        |     6 | In Progress |
| 19 Apr 23 19:00 PDT | Los Angeles Kings (1-0)  |     0 | Edmonton Oilers (0-1)     |     2 | In Progress |
| 19 Apr 23 16:00 PDT | New York Islanders (0-2) |     3 | Carolina Hurricanes (2-0) |     4 | Final       |
| 19 Apr 23 16:30 PDT | Florida Panthers (1-1)   |     6 | Boston Bruins (1-1)       |     3 | Final       |
+---------------------+--------------------------+-------+---------------------------+-------+-------------+
```

---

### Print scores for yesterday's NHL games

```sh
who-scored yesterday
```

#### Example output

```
+---------------------+---------------------------+-------+----------------------------+-------+--------+
| START TIME          | AWAY TEAM (W-L)           | SCORE | HOME TEAM (W-L)            | SCORE | STATUS |
+---------------------+---------------------------+-------+----------------------------+-------+--------+
| 18 Apr 23 16:00 PDT | New York Rangers (1-0)    |     5 | New Jersey Devils (0-1)    |     1 | Final  |
| 18 Apr 23 16:30 PDT | Tampa Bay Lightning (1-0) |     7 | Toronto Maple Leafs (0-1)  |     3 | Final  |
| 18 Apr 23 18:30 PDT | Winnipeg Jets (1-0)       |     5 | Vegas Golden Knights (0-1) |     1 | Final  |
| 18 Apr 23 19:00 PDT | Seattle Kraken (1-0)      |     3 | Colorado Avalanche (0-1)   |     1 | Final  |
+---------------------+---------------------------+-------+----------------------------+-------+--------+
```
