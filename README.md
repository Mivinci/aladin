# aladin
Aladin provides elegant interfaces and conventions for building configuration manager that reads config sources from local file, environment variables, remote configuration service or those stored in Redis.

This project is a conclusion I made after studying the source code of [go-micro](github.com/micro/go-micro) which's a rocking framework for building micro-services.

Aladin, I guess, is capable enough of building a full functioning configuration manager and two implementations are already written by me, you can use it to read configs from environment variables or a local file.

## Usage

### Install

```bash
go get -u github.com/mivinci/aladin
```

### Quick Start
Suppose we've got a YAML file named `config.yml` which has fields:
```yaml
name: xjj
age: 18
bio:
  favorites:
    - wfs
    - sing
```

To get a field value, just tell aladin what you want:

```go
aladin.Init()
aladin.Get("name").String("") // xjj
aladin.Get("bio.favorites").StringSlice(nil) // ["wfs", "sing"]
```

## Functionality

**Custom configuration sources**

To fetch source from environment variables, use an `EnvSource`:

```go
aladin.Init(
    aladin.WithSource(aladin.NewEnvSource()),
)
```

If you wanna get value of env variable `DATABASE_HOST`, just call:

```go
aladin.Get("database.host") 
```

Aladin now supports only `EnvSource` and `FileSource` which's set as default when calling `aladin.Init()`, you can send a PR to provide more useful sources like 

- Remote configuration services
- One stored in Redis 

etc.

**Custom file parsers**

If you're gonna use a `FileSource` for your configuration, you can pick up a parser - a YAML parser is set as default - that helps handle the content of your config file correctly.

So far, aladin has supported `YAML` and `JSON` formats of file. To use one of them, call:

```go
aladin.Init(
    aladin.WithParser(aladin.NewJSONParser()),
)
```

Then you can delightfully use aladin to get field values from your configuration file with the matching format.

**Enable hot-reload**

The `FileSource` supports - `EnvSource` **does not** - a `hot-reload` feature which means you don't have to shutdown a server when updating some configuration. Enable it on initialing the aladin.

```go
aladin.Init(aladin.WithHotReload())
```

After starting your server, try modifying your configuration file, you can see the `aladin.Get` method returns a different value :)

------

Have fun!



