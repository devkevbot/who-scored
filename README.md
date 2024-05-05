# who-scored

`who-scored` is a command-line application that finds and displays scores for NHL games.

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
+-----------+--------------+----------------+---------------------+------------+---------+-------------+
| GAME TYPE | PLAYOFF GAME | PLAYOFF SERIES | START TIME          | TEAMS      | SCORE   | STATUS      |
+-----------+--------------+----------------+---------------------+------------+---------+-------------+
| Playoffs  | R2, GM 1     | TIED 0-0       | 05 May 24 13:00 PDT | CAR at NYR | 2-1 NYR | In-progress |
| Playoffs  | R1, GM 7     | TIED 3-3       | 05 May 24 16:30 PDT | VGK at DAL | 0-0     | Not started |
+-----------+--------------+----------------+---------------------+------------+---------+-------------+
```

---

### Print scores for yesterday's NHL games

#### Example input

```sh
who-scored yesterday
```

#### Example output

```
+-----------+--------------+----------------+---------------------+------------+---------+--------+----------------------------------------------------+
| GAME TYPE | PLAYOFF GAME | PLAYOFF SERIES | START TIME          | TEAMS      | SCORE   | STATUS | GAME-WINNING GOAL                                  |
+-----------+--------------+----------------+---------------------+------------+---------+--------+----------------------------------------------------+
| Playoffs  | R1, GM 6     | VAN 4-2        | 03 May 24 16:00 PDT | VAN at NSH | 1-0 VAN | Final  | P. Suter (2) from B. Boeser (2), E. Pettersson (3) |
| Playoffs  | R1, GM 6     | TIED 3-3       | 03 May 24 19:00 PDT | DAL at VGK | 2-0 VGK | Final  | N. Hanifin (2)                                     |
+-----------+--------------+----------------+---------------------+------------+---------+--------+----------------------------------------------------+
```

### Print scores for games played on a single day

#### Example input

```sh
# Date must be in YYYY-MM-DD format
who-scored on 2023-03-16
```

#### Example output

```
+----------------+---------------------+------------+---------+------------+-------------------------------------------------------------+
| GAME TYPE      | START TIME          | TEAMS      | SCORE   | STATUS     | GAME-WINNING GOAL                                           |
+----------------+---------------------+------------+---------+------------+-------------------------------------------------------------+
| Regular Season | 16 Mar 23 16:00 PDT | COL at OTT | 5-4 COL | Final      | L. Eller (8)                                                |
| Regular Season | 16 Mar 23 16:00 PDT | MTL at FLA | 9-5 FLA | Final      | S. Reinhart (24) from J. Mahura (11)                        |
| Regular Season | 16 Mar 23 16:00 PDT | PIT at NYR | 4-2 NYR | Final      | C. Kreider (29) from V. Trocheck (34), A. Fox (51)          |
| Regular Season | 16 Mar 23 16:00 PDT | TBL at NJD | 4-3 TBL | Final (SO) | A. Killorn                                                  |
| Regular Season | 16 Mar 23 17:00 PDT | BOS at WPG | 3-0 BOS | Final      | T. Frederic (15) from T. Bertuzzi (13), C. Coyle (25)       |
| Regular Season | 16 Mar 23 17:00 PDT | CHI at NSH | 2-1 CHI | Final      | J. Anderson (5) from B. Katchouk (7)                        |
| Regular Season | 16 Mar 23 18:00 PDT | DAL at EDM | 4-1 EDM | Final      | M. Janmark (7) SHG from C. McDavid (74), V. Desharnais (5)  |
| Regular Season | 16 Mar 23 19:00 PDT | CGY at VGK | 7-2 CGY | Final      | B. Coleman (15) from E. Lindholm (38), M. Weegar (20)       |
| Regular Season | 16 Mar 23 19:00 PDT | VAN at ARI | 3-2 ARI | Final      | L. Crouse (22) PPG from V. Soderstrom (6), M. Maccelli (32) |
| Regular Season | 16 Mar 23 19:30 PDT | CBJ at LAK | 4-1 LAK | Final      | A. Kopitar (26) from Q. Byfield (16), M. Anderson (13)      |
| Regular Season | 16 Mar 23 19:30 PDT | SEA at SJS | 2-1 SEA | Final (OT) | V. Dunn (13) from O. Bjorkstrand (22), J. McCann (20)       |
+----------------+---------------------+------------+---------+------------+-------------------------------------------------------------+
```

## Install

The program can either be installed from a release artifact or built locally.

### Install from a release artifact

1. Go to the [Releases page](https://github.com/devkevbot/who-scored/releases), download, and extract the latest version for your system.
2. Add the executable file to `$PATH`.

### Building from source

1. Clone the repository:

   ```sh
    https://github.com/devkevbot/who-scored.git
   ```

2. From the repository directory, build and install the project: `go install .`
