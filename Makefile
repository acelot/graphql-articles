default: help

.PHONY: help
help: ## Show help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.PHONY: tools
tools: ## Install required tools
	go get -u github.com/99designs/gqlgen@v0.13.0
	go get -u github.com/vektah/dataloaden@v0.3.0
	go get -u golang.org/x/lint/golint

.PHONY: clean
clean: ## Run `go clean`
	go clean

.PHONY: fmt
fmt: ## Run `go fmt`
	go fmt ./...

.PHONY: lint
lint: ## Run `go lint`
	golint ./...

.PHONY: test
test: generate ## Run `go test`
	go test -v ./... -short

.PHONY: generate-gqlgen
generate-gqlgen: ## Generate GraphQL models and resolvers
	gqlgen generate

.PHONY: generate-dataloaders
generate-dataloaders: ## Generate dataloaders
	(cd internal/feature/article \
     && dataloaden LoaderByID LoaderByIDKey \*github.com/acelot/articles/pkg/feature/article.Article)
	(cd internal/feature/articleblock \
     && dataloaden LoaderByArticleID LoaderByArticleIDKey \[\]github.com/acelot/articles/pkg/feature/articleblock.ArticleBlock)
	(cd internal/feature/articletag \
     && dataloaden LoaderByArticleID LoaderByArticleIDKey \[\]github.com/acelot/articles/pkg/feature/articletag.ArticleTag)
	(cd internal/feature/image \
     && dataloaden LoaderByID LoaderByIDKey \*github.com/acelot/articles/pkg/feature/image.Image)
	(cd internal/feature/project \
     && dataloaden LoaderByID LoaderByIDKey \*github.com/acelot/articles/pkg/feature/project.Project)
	(cd internal/feature/tag \
     && dataloaden LoaderByID LoaderByIDKey \*github.com/acelot/articles/pkg/feature/tag.Tag)

.PHONY: generate
generate: generate-gqlgen generate-dataloaders ## Generate all
	go generate ./...

.PHONY: build-app
build-app: generate ## Build application
	CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o app cmd/app/main.go

.PHONY: build-migration
build-migration: ## Build migration command
	CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o migration cmd/migration/main.go

.PHONY: build
build: generate build-app build-migration ## Build all
