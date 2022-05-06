# anorm-cli

An anorm Code generator written in Go.

## usage

```
anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome"

anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome" -o "/to/path"

anorm-cli gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -m "awesome" -au "bigboss" -o "/to/path"
```

