# DTN's Simplified QA Framework

The goal was to take the carbon framework (https://github.com/pantheon-systems/carbon-webservice-testing), and update it with a handful of changes to make it a more "robust" as well as better "scalable" framework.

# Running It

This code requires you to have docker installed and running, as well as starting the offical go-vcs-service (tests run against the DB and API)

## The Service

Clone, Install, and start the go-vcs-service (https://github.com/pantheon-systems/qa-vcs-service)

```
$ task local-clean
$ task task local-db
$ task run:local-service
```
NOTE: depending on configs, you might need to run `$ task local-pubsubs` as well

## Running the QA Framework

Clone, Install, and start the qa-vcs-service-dtnp (https://github.com/dtnp/qa-vcs-service-dtnp)

### Generate the swagger based tests

```
$ rm -rf test/swagger
$ go run main/main.go --swaggerFile=swagger3.json
```
Output can be found specifically in the `test/swagger/` folder.  

### Individual Test

The following will run an individual test

```
$ go test -count=1 -v -run TestGETRequestingsourceId200 ./test/swagger/requesting_source/...
```

### Things to ignore

Ignore what is in the `test/` folder that is NOT specifically in the SWAGGER folder.

### Cleanup

DELETE the `test/swagger` folder between runs if you want to edit the swagger generating code.

