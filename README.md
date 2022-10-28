# Solution

A simple command line program that summarizes a local log file, where csv is parsed into different maps based on main operation to help parsing summarization data easier.

# Example Commands

## Operations

```
log-analyzer --file=/home/aaron/Downloads/coding_challenge_go/server_log.csv operations summarize
```

```
log-analyzer --file=/home/aaron/Downloads/coding_challenge_go/server_log.csv operations summarize --size 50
```

## Users

```
log-analyzer --file=/home/aaron/Downloads/coding_challenge_go/server_log.csv users unique
```

```
log-analyzer --file=/home/aaron/Downloads/coding_challenge_go/server_log.csv users summarize
```

```
log-analyzer --file=/home/aaron/Downloads/coding_challenge_go/server_log.csv users summarize --name sarah94 --date 2020-04-13
```

# Improvement

* Make more general summary functions that can be reused more across different filtering criteria.
* Open up the potential to grab logs from remote locations.
* Investigate potential ways to summarize information while looping through general array vs specific maps.
* Make the output of the commands prettier.