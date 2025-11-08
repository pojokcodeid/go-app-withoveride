# go-app-withoveride

## jalankan overide

```
cd myapp-custom
go rum .

curl -X POST http://localhost:8080/users -d '{"name":"Alice"}'
curl http://localhost:8080/users
```

## jalankan main

```
cd myapp-main
go rum .

curl -X POST http://localhost:8080/users -d '{"name":"Alice"}'
curl http://localhost:8080/users
```
