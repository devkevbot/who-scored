# who-scored
A CLI application to output the scores for NHL games on a specific date.

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

### Print scores for an arbitrary date in the past

```sh
who-scored on 2023-03-16
```

#### Example output

```
+---------------------+------------------------------+-------+-------------------------------+-------+--------+
| START TIME          | AWAY TEAM (W-L)              | SCORE | HOME TEAM (W-L)               | SCORE | STATUS |
+---------------------+------------------------------+-------+-------------------------------+-------+--------+
| 13 Apr 23 16:00 PDT | Boston Bruins (65-12)        |     5 | Montréal Canadiens (31-45)    |     4 | Final  |
| 13 Apr 23 16:00 PDT | Carolina Hurricanes (52-21)  |     6 | Florida Panthers (42-32)      |     4 | Final  |
| 13 Apr 23 16:00 PDT | Detroit Red Wings (35-37)    |     0 | Tampa Bay Lightning (46-30)   |     5 | Final  |
| 13 Apr 23 16:00 PDT | New Jersey Devils (52-22)    |     5 | Washington Capitals (35-37)   |     4 | Final  |
| 13 Apr 23 16:00 PDT | Ottawa Senators (39-35)      |     3 | Buffalo Sabres (41-33)        |     4 | Final  |
| 13 Apr 23 16:00 PDT | Pittsburgh Penguins (40-31)  |     2 | Columbus Blue Jackets (25-47) |     3 | Final  |
| 13 Apr 23 16:00 PDT | Toronto Maple Leafs (50-21)  |     3 | New York Rangers (47-22)      |     2 | Final  |
| 13 Apr 23 17:00 PDT | Winnipeg Jets (46-33)        |     2 | Colorado Avalanche (50-24)    |     4 | Final  |
| 13 Apr 23 17:00 PDT | Minnesota Wild (46-25)       |     3 | Nashville Predators (42-31)   |     4 | Final  |
| 13 Apr 23 17:00 PDT | St. Louis Blues (37-38)      |     0 | Dallas Stars (47-21)          |     1 | Final  |
| 13 Apr 23 17:30 PDT | Philadelphia Flyers (31-38)  |     5 | Chicago Blackhawks (26-49)    |     4 | Final  |
| 13 Apr 23 18:00 PDT | San Jose Sharks (22-44)      |     2 | Edmonton Oilers (50-23)       |     5 | Final  |
| 13 Apr 23 19:00 PDT | Los Angeles Kings (47-25)    |     5 | Anaheim Ducks (23-47)         |     3 | Final  |
| 13 Apr 23 19:00 PDT | Vancouver Canucks (38-37)    |     5 | Arizona Coyotes (28-40)       |     4 | Final  |
| 13 Apr 23 19:30 PDT | Vegas Golden Knights (51-22) |     3 | Seattle Kraken (46-28)        |     1 | Final  |
+---------------------+------------------------------+-------+-------------------------------+-------+--------+
```