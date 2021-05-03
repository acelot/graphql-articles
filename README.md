# graphql-articles
Articles GraphQL Service

Requierements:
- libvips

## How to build

```shell
make build
```

## How to apply DB migrations

```shell
migration update postgres://user:pass@localhost/articles\?sslmode=disable
```

## How to run app

```shell
app postgres://user:pass@localhost/articles\?sslmode=disable
```