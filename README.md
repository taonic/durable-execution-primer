## A Primer on Durable Execution - Sample Code

### How to run it?

Run the Temporal dev server:

```
temporal server start-dev
```

Start a worker, e.g.
```
go run 1_setup/worker.go
```

Execute a workflow:
```
go run client.go
```
