# binproto
A binary protocol in go

# usage
```bash
# run server
$ go run ./cmd/server
# run interactive client
$ go run ./cmd/client
> ..
```

# test
```bash
$ go test
```

# todos
- [x] proto
- [x] server and interactive client
- [x] fixed size packet buf
- [x] tests
- [ ] mask
- [x] continuation frame bit
- [ ] include array len in payload so we may use one call to binary.Read|Write
- [ ] debug mode
- [ ] try passing a slice to binary.Read
- [ ] add packets.Add
- [ ] make field Fin private and use Fin()
- [ ] post repo to [mr Vladimirs blog](https://medium.com/learning-the-go-programming-language/encoding-data-with-the-go-binary-package-42c7c0eb3e73)

# license
MIT
