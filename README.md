## Description

Monorepo powered by [NX](https://nx.dev/)

[Golang Getting Started](https://github.com/nx-go/nx-go)

✨ **GO library** ✨


```
NAME=env && nx g @obiente-lab/nx-go:library $NAME --directory library/go/$NAME

To remove:
nx g rm <name>
```

✨ **GO application** ✨


```
NAME=api && nx g @obiente-lab/nx-go:application $NAME --directory apps/$NAME

To remove:
nx g rm <name>
```

✨ **Commands** ✨

**For Go library**

test
```bash
nx test <name>
```

lint
```bash
nx lint <name>
```

**For Go Application**
serve
```bash
nx serve <name>
```

lint
```bash
nx lint <name>
```

test
```bash
nx test <name>
```

build
```bash
nx build <name>
```

### Prerequisites

`protoc` versions

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```
wget [link](https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip)

`gow`

```
go install github.com/mitranim/gow@latest
```