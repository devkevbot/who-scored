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
+-----------+--------------+----------------+---------------------+-----------------------+---------+-------------+-------------------+
| GAME TYPE | PLAYOFF GAME | PLAYOFF SERIES | START TIME          | TEAMS                 | SCORE   | STATUS      | GAME-WINNING GOAL |
+-----------+--------------+----------------+---------------------+-----------------------+---------+-------------+-------------------+
| Playoffs  | R1, GM 7     | TIED 3-3       | 04 May 24 17:00 PDT | Maple Leafs at Bruins | 0-0 TIE | Not started | (No Score)        |
+-----------+--------------+----------------+---------------------+-----------------------+---------+-------------+-------------------+
```

---

### Print scores for yesterday's NHL games

#### Example input

```sh
who-scored yesterday
```

#### Example output

```
+-----------+--------------+----------------+---------------------+-------------------------+---------+--------+-------------------------------------------------------+
| GAME TYPE | PLAYOFF GAME | PLAYOFF SERIES | START TIME          | TEAMS                   | SCORE   | STATUS | GAME-WINNING GOAL                                     |
+-----------+--------------+----------------+---------------------+-------------------------+---------+--------+-------------------------------------------------------+
| Playoffs  | R1, GM 6     | VAN 4-2        | 03 May 24 16:00 PDT | Canucks at Predators    | 1-0 VAN | Final  | P. Suter (2) from B. Boeser (2), E. Pettersson (3)    |
| Playoffs  | R1, GM 6     | TIED 3-3       | 03 May 24 19:00 PDT | Stars at Golden Knights | 2-0 VGK | Final  | M. Stone (3) from W. Karlsson (1), A. Pietrangelo (1) |
+-----------+--------------+----------------+---------------------+-------------------------+---------+--------+-------------------------------------------------------+
```

### Print scores for games played on a single day

#### Example input

```sh
# Date must be in YYYY-MM-DD format
who-scored on 2023-03-16
```

#### Example output

```
+----------------+---------------------+--------------------------+---------+------------+--------------------------------------------------------------+
| GAME TYPE      | START TIME          | TEAMS                    | SCORE   | STATUS     | GAME-WINNING GOAL                                            |
+----------------+---------------------+--------------------------+---------+------------+--------------------------------------------------------------+
| Regular Season | 16 Mar 23 16:00 PDT | Avalanche at Senators    | 5-4 COL | Final      | B. Tkachuk (28) PPG from T. Stützle (42), J. Sanderson (24)  |
| Regular Season | 16 Mar 23 16:00 PDT | Canadiens at Panthers    | 9-5 FLA | Final      | R. Pitlick (5) PPG from M. Matheson (16)                     |
| Regular Season | 16 Mar 23 16:00 PDT | Penguins at Rangers      | 4-2 NYR | Final      | C. Kreider (30) from V. Trocheck (35), M. Zibanejad (37)     |
| Regular Season | 16 Mar 23 16:00 PDT | Lightning at Devils      | 4-3 TBL | Final (SO) | A. Killorn (22)                                              |
| Regular Season | 16 Mar 23 17:00 PDT | Bruins at Jets           | 3-0 BOS | Final      | T. Nosek (6) from D. Pastrnak (44)                           |
| Regular Season | 16 Mar 23 17:00 PDT | Blackhawks at Predators  | 2-1 CHI | Final      | R. Josi (18) from T. Novak (16), P. Tomasino (7)             |
| Regular Season | 16 Mar 23 18:00 PDT | Stars at Oilers          | 4-1 EDM | Final      | M. Janmark (8)                                               |
| Regular Season | 16 Mar 23 19:00 PDT | Flames at Golden Knights | 7-2 CGY | Final      | D. Dube (18) from N. Zadorov (6), J. Huberdeau (32)          |
| Regular Season | 16 Mar 23 19:00 PDT | Canucks at Coyotes       | 3-2 ARI | Final      | E. Pettersson (31) from A. Beauvillier (20), C. Garland (26) |
| Regular Season | 16 Mar 23 19:30 PDT | Blue Jackets at Kings    | 4-1 LAK | Final      | K. Marchenko (17) from G. Bayreuther (9)                     |
| Regular Season | 16 Mar 23 19:30 PDT | Kraken at Sharks         | 2-1 SEA | Final (OT) | V. Dunn (13) from O. Bjorkstrand (22), J. McCann (20)        |
+----------------+---------------------+--------------------------+---------+------------+--------------------------------------------------------------+
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
