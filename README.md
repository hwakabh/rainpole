# rainpole
Dummy Company APIs with Go

## Local Setup
For build application with generating single binary of rainpole API:

```shell
% git clone git@github.com:hwakabh/rainpole.git
% make db
% make vuild

# Or, alternatively you can run `go build` manually
% scripts/initdb.sh
% GOOS=linux GOARCH=amd64 go build -o ./cmd/rainpole
```

## Deployment
TBA

## Good to know
TBA

## License
TBA
