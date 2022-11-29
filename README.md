# Tutorial for Hexagonal Architecture that utilizes gRPC

A foray into the world of [gRPC](https://grpc.io/docs/) and [hexagonal architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>) following this [Youtube Tutorial](https://www.youtube.com/watch?v=MpFog2kZsHk&list=PLDs4q279bMw8PENgLN8RAVZbsHcBFASnw&index=2)

## Requirements

- [Golang](https://go.dev/doc/install)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)
- [Docker](https://www.docker.com/get-started/)
- Docker Compose
- [Task](https://taskfile.dev/)

Optionally, if using a Mac and Homebrew:

```
brew install go protobuf go-task/tap/go-task
brew install --cask docker
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Running Tests

```
task gen-grpc
task test
```
