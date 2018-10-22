# Twitter Frequenter

## Inspiration

<https://blog.twitter.com/official/en_us/topics/company/2018/enabling-further-research-of-information-operations-on-twitter.html>

<https://about.twitter.com/en_us/values/elections-integrity.html#data>

## Building

- Build the main command:

```shell
$ go build ./cmd/twitter-frequenter
<Empty output on build success>
```

- Run the concurrent tweet parser:

```shell
$ ./twitter-frequenter -f ~/Dropbox/Stuff/ira_tweets_csv_hashed.csv -t=tweet -c
2018/10/20 01:13:11 Starting twitter-frequenter @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting concurrent processing @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting 7 workers...
2018/10/20 01:13:11 Finished tweet waits @ 2018-10-20T01:13:11+02:00
```

- Run the concurrent user parser:

```shell
$ ./twitter-frequenter -f ~/Dropbox/Stuff/ira_users_csv_hashed.csv -t=user -c
2018/10/20 01:13:11 Starting twitter-frequenter @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting concurrent processing @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting 7 workers...
2018/10/20 01:13:11 Finished user waits @ 2018-10-20T01:13:11+02:00
```

- Run the non-concurrent tweet parser:

```shell
$ ./twitter-frequenter -f ~/Dropbox/Stuff/ira_tweets_csv_hashed.csv -t=tweet
2018/10/20 01:13:11 Starting twitter-frequenter @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting concurrent processing @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting 7 workers...
2018/10/20 01:13:11 Finished tweet waits @ 2018-10-20T01:13:11+02:00
```

- Run the non-concurrent user parser:

```shell
$ ./twitter-frequenter -f ~/Dropbox/Stuff/ira_users_csv_hashed.csv -t=user
2018/10/20 01:13:11 Starting twitter-frequenter @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting concurrent processing @ 2018-10-20T01:13:11+02:00
2018/10/20 01:13:11 Starting 7 workers...
2018/10/20 01:13:11 Finished user waits @ 2018-10-20T01:13:11+02:00
```

## Running Tests

- Run the testsuite with coverage enabled:

```shell
$ go test -coverpkg=all ./... -coverprofile=coverage.out
?       gitlab.com/TheDonDope/twitter-frequenter/cmd/twitter-frequenter     [no test files]
?       gitlab.com/TheDonDope/twitter-frequenter/pkg/api      [no test files]
ok      gitlab.com/TheDonDope/twitter-frequenter/pkg/test     0.136s  coverage: 18.7% of statements in all
?       gitlab.com/TheDonDope/twitter-frequenter/pkg/types    [no test files]
?       gitlab.com/TheDonDope/twitter-frequenter/pkg/util/errors      [no test files]
?       gitlab.com/TheDonDope/twitter-frequenter/pkg/version  [no test files]
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
