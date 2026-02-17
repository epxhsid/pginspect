# pginspect

![CI](https://github.com/epxhsid/pginspect/actions/workflows/go.yml/badge.svg)
![Go Version](https://img.shields.io/badge/go-1.25.4+-blue)
![License](https://img.shields.io/github/license/epxhsid/pginspect)

lightweight Go library for inspecting pgSQL database in a browser.  
where it provides a simple UI to view schemas, tables, and table data via REST endpoints

## Installation
```bash
go get github.com/epxhsid/pginspect
go mod tidy # needed to retrieve other dependencies
```

### How to use
1. Configuration example
```go
eng, err := engine.NewConn(context.Background(), &engine.Options{
		Addr:     os.Getenv("postgres://user:pass@host:port/dbname"),
		MaxRows:  100,
		ReadOnly: true,
})
```
2. Mount endpoint example
```go
endpoint := "/__db"
mux := http.NewServeMux()
httpui.Mount(mux, endpoint, eng)
```

3. Access it via your defined endpoint, in this case:
```shell
http://ADDRESS:PORT/__db # replace __db with your endpoint
```

## Note
> Note: `engine.NewConn` creates its own connection pool internally.  
> If you already have a `*pgxpool.Pool`, you can still use it by passing its connection string; this library will open a separate pool without affecting your existing connections.
