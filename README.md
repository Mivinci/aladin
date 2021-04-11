# aladin
Aladin provides elegant interfaces and conventions for building configuration manager which helps you build one that reads config sources from local file, environment variables, remote configuration service or those stored in Redis.

This project is a conclusion I made after studying the source code of [go-micro](github.com/micro/go-micro) which's a rocking framework for building micro-services.

Aladin, I guess, is capable enough of building a full functioning configuration manager and two implementations are already written by me, you can use it to read configs from environment variables or a local file.

## Usage

```bash
go get -u github.com/mivinci/aladin
```
