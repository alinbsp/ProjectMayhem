# ProjectMayhem Tracker - Go implementation

## How to run
This guide assumes you will be using CockroachDB as your database.
The easiest way to have a running database is to use Docker.
If you haven't already done so, go ahead and install Docker.

Then:
```
$ docker pull cockroachdb/cockroach

$ docker run -d -p 8090:8080 -p 26257:26257 --name cockroachdb-mayhem cockroachdb/cockroach start --insecure
```

Now that we have a CockroachDB node up and running, we need to create an user and a database.
```
$ docker exec -it cockroachdb-mayhem bash

(cockroachdb-mayhem)$ ./cockroach user set mayhem --insecure

(cockroachdb-mayhem)$ ./cockroach sql --insecure -e 'CREATE DATABASE mayhem'

(cockroachdb-mayhem)$ ./cockroach sql --insecure -e 'GRANT ALL ON DATABASE mayhem TO mayhem'

(cockroachdb-mayhem)$ exit
```

Get and run the project. In the project directory ('tracker-go') run:
```
go run main.go
```
