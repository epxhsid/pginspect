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
		Addr:     os.Getenv("DATABASE_URL"),
		MaxRows:  100,
		ReadOnly: true,
})
```
2. Mount endpoint example
```go
mux := http.NewServeMux()
httpui.Mount(mux, "/__db", eng)
```

3. Access it via your defined endpoint, in this case:
```markdown
http://ADDRESS:PORT/__db
```

## Note
> Note: `engine.NewConn` creates its own connection pool internally.  
> If you already have a `*pgxpool.Pool`, you can still use it by passing its connection string; this library will open a separate pool without affecting your existing connections.
