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

#### Example input

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

### Print scores for games played on a single day

#### Example input

```sh
# Date must be in YYYY-MM-DD format
who-scored on 2023-03-16
```

#### Example output

```
+----------------+---------------------+---------------------------------+-------+--------------------------------+-------+--------+
| GAME TYPE      | START TIME          | AWAY TEAM (RECORD)              | SCORE | HOME TEAM (RECORD)             | SCORE | STATUS |
+----------------+---------------------+---------------------------------+-------+--------------------------------+-------+--------+
| Regular Season | 16 Mar 23 16:00 PDT | Colorado Avalanche (39-22-6)    | 5     | Ottawa Senators (33-31-4)      | 4     | Final  |
| Regular Season | 16 Mar 23 16:00 PDT | Montréal Canadiens (27-36-6)    | 5     | Florida Panthers (34-27-7)     | 9     | Final  |
| Regular Season | 16 Mar 23 16:00 PDT | Pittsburgh Penguins (34-24-10)  | 2     | New York Rangers (39-19-10)    | 4     | Final  |
| Regular Season | 16 Mar 23 16:00 PDT | Tampa Bay Lightning (41-22-6)   | 4     | New Jersey Devils (44-17-7)    | 3     | Final  |
| Regular Season | 16 Mar 23 17:00 PDT | Boston Bruins (51-11-5)         | 3     | Winnipeg Jets (38-28-3)        | 0     | Final  |
| Regular Season | 16 Mar 23 17:00 PDT | Chicago Blackhawks (24-38-6)    | 2     | Nashville Predators (34-25-7)  | 1     | Final  |
| Regular Season | 16 Mar 23 18:00 PDT | Dallas Stars (37-19-13)         | 1     | Edmonton Oilers (38-23-8)      | 4     | Final  |
| Regular Season | 16 Mar 23 19:00 PDT | Calgary Flames (31-24-14)       | 7     | Vegas Golden Knights (42-21-6) | 2     | Final  |
| Regular Season | 16 Mar 23 19:00 PDT | Vancouver Canucks (29-33-5)     | 2     | Arizona Coyotes (26-32-11)     | 3     | Final  |
| Regular Season | 16 Mar 23 19:30 PDT | Columbus Blue Jackets (21-39-7) | 1     | Los Angeles Kings (40-20-9)    | 4     | Final  |
| Regular Season | 16 Mar 23 19:30 PDT | Seattle Kraken (38-23-7)        | 2     | San Jose Sharks (19-36-14)     | 1     | Final  |
+----------------+---------------------+---------------------------------+-------+--------------------------------+-------+--------+
```

### Print scores for games played between a start and end date (inclusive)

#### Example input

```sh
# Both the start and end date must be in YYYY-MM-DD format
who-scored during 2023-04-14 2023-04-18
```

#### Example output

```
+----------------+---------------------+------------------------------+-------+---------------------------------+-------+--------+
| GAME TYPE      | START TIME          | AWAY TEAM (RECORD)           | SCORE | HOME TEAM (RECORD)              | SCORE | STATUS |
+----------------+---------------------+------------------------------+-------+---------------------------------+-------+--------+
| Regular Season | 14 Apr 23 16:30 PDT | Buffalo Sabres (42-33-7)     | 5     | Columbus Blue Jackets (25-48-9) | 2     | Final  |
| Regular Season | 14 Apr 23 17:00 PDT | Colorado Avalanche (51-24-7) | 4     | Nashville Predators (42-32-8)   | 3     | Final  |
+----------------+---------------------+------------------------------+-------+---------------------------------+-------+--------+
| Playoffs       | 17 Apr 23 16:00 PDT | New York Islanders (0-1)     | 1     | Carolina Hurricanes (1-0)       | 2     | Final  |
| Playoffs       | 17 Apr 23 16:30 PDT | Florida Panthers (0-1)       | 1     | Boston Bruins (1-0)             | 3     | Final  |
| Playoffs       | 17 Apr 23 18:30 PDT | Minnesota Wild (1-0)         | 3     | Dallas Stars (0-1)              | 2     | Final  |
| Playoffs       | 17 Apr 23 19:00 PDT | Los Angeles Kings (1-0)      | 4     | Edmonton Oilers (0-1)           | 3     | Final  |
+----------------+---------------------+------------------------------+-------+---------------------------------+-------+--------+
| Playoffs       | 18 Apr 23 16:00 PDT | New York Rangers (1-0)       | 5     | New Jersey Devils (0-1)         | 1     | Final  |
| Playoffs       | 18 Apr 23 16:30 PDT | Tampa Bay Lightning (1-0)    | 7     | Toronto Maple Leafs (0-1)       | 3     | Final  |
| Playoffs       | 18 Apr 23 18:30 PDT | Winnipeg Jets (1-0)          | 5     | Vegas Golden Knights (0-1)      | 1     | Final  |
| Playoffs       | 18 Apr 23 19:00 PDT | Seattle Kraken (1-0)         | 3     | Colorado Avalanche (0-1)        | 1     | Final  |
+----------------+---------------------+------------------------------+-------+---------------------------------+-------+--------+
```
---

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
