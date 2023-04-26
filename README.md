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
+-----------+---------------------+--------------------------+-------+---------------------------+-------+-------------+
| GAME TYPE | START TIME          | AWAY TEAM (RECORD)       | SCORE | HOME TEAM (RECORD)        | SCORE | STATUS      |
+-----------+---------------------+--------------------------+-------+---------------------------+-------+-------------+
| Playoffs  | 25 Apr 23 16:00 PDT | New York Islanders (1-3) | 2     | Carolina Hurricanes (3-1) | 1     | In Progress |
| Playoffs  | 25 Apr 23 17:00 PDT | Minnesota Wild (2-2)     | 0     | Dallas Stars (2-2)        | 2     | In Progress |
| Playoffs  | 25 Apr 23 18:30 PDT | Los Angeles Kings (2-2)  | 0     | Edmonton Oilers (2-2)     | 0     | Scheduled   |
+-----------+---------------------+--------------------------+-------+---------------------------+-------+-------------+
```

---

### Print scores for yesterday's NHL games

```sh
who-scored yesterday
```

#### Example output

```
+-----------+---------------------+----------------------------+-------+---------------------------+-------+--------+
| GAME TYPE | START TIME          | AWAY TEAM (RECORD)         | SCORE | HOME TEAM (RECORD)        | SCORE | STATUS |
+-----------+---------------------+----------------------------+-------+---------------------------+-------+--------+
| Playoffs  | 24 Apr 23 16:00 PDT | New Jersey Devils (2-2)    | 3     | New York Rangers (2-2)    | 1     | Final  |
| Playoffs  | 24 Apr 23 16:30 PDT | Toronto Maple Leafs (3-1)  | 5     | Tampa Bay Lightning (1-3) | 4     | Final  |
| Playoffs  | 24 Apr 23 18:30 PDT | Vegas Golden Knights (3-1) | 4     | Winnipeg Jets (1-3)       | 2     | Final  |
| Playoffs  | 24 Apr 23 19:00 PDT | Colorado Avalanche (2-2)   | 2     | Seattle Kraken (2-2)      | 3     | Final  |
+-----------+---------------------+----------------------------+-------+---------------------------+-------+--------+
```

### Print scores for an arbitrary date in the past

```sh
who-scored on 2023-03-16
```

#### Example output

```
+----------------+---------------------+--------+---------------------------------+-------+--------------------------------+-------+
| GAME TYPE      | START TIME          | STATUS | AWAY TEAM (RECORD)              | SCORE | HOME TEAM (RECORD)             | SCORE |
+----------------+---------------------+--------+---------------------------------+-------+--------------------------------+-------+
| Regular Season | 16 Mar 23 16:00 PDT | Final  | Colorado Avalanche (39-22-6)    | 5     | Ottawa Senators (33-31-4)      | 4     |
| Regular Season | 16 Mar 23 16:00 PDT | Final  | Montréal Canadiens (27-36-6)    | 5     | Florida Panthers (34-27-7)     | 9     |
| Regular Season | 16 Mar 23 16:00 PDT | Final  | Pittsburgh Penguins (34-24-10)  | 2     | New York Rangers (39-19-10)    | 4     |
| Regular Season | 16 Mar 23 16:00 PDT | Final  | Tampa Bay Lightning (41-22-6)   | 4     | New Jersey Devils (44-17-7)    | 3     |
| Regular Season | 16 Mar 23 17:00 PDT | Final  | Boston Bruins (51-11-5)         | 3     | Winnipeg Jets (38-28-3)        | 0     |
| Regular Season | 16 Mar 23 17:00 PDT | Final  | Chicago Blackhawks (24-38-6)    | 2     | Nashville Predators (34-25-7)  | 1     |
| Regular Season | 16 Mar 23 18:00 PDT | Final  | Dallas Stars (37-19-13)         | 1     | Edmonton Oilers (38-23-8)      | 4     |
| Regular Season | 16 Mar 23 19:00 PDT | Final  | Calgary Flames (31-24-14)       | 7     | Vegas Golden Knights (42-21-6) | 2     |
| Regular Season | 16 Mar 23 19:00 PDT | Final  | Vancouver Canucks (29-33-5)     | 2     | Arizona Coyotes (26-32-11)     | 3     |
| Regular Season | 16 Mar 23 19:30 PDT | Final  | Columbus Blue Jackets (21-39-7) | 1     | Los Angeles Kings (40-20-9)    | 4     |
| Regular Season | 16 Mar 23 19:30 PDT | Final  | Seattle Kraken (38-23-7)        | 2     | San Jose Sharks (19-36-14)     | 1     |
+----------------+---------------------+--------+---------------------------------+-------+--------------------------------+-------+
```