# Clicker

click-id generation service

## Installation

To install, use `go get`:

```shell
go get github.com/octoclick/clicker
```

Для установки и использования easyjson, выполните следующие шаги:

1. Сначала установите easyjson с помощью команды:

```bash
go install github.com/mailru/easyjson/...@latest
```

2. Добавьте путь к binary файлам Go в ваш PATH. Обычно это:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Можете добавить эту строку в ваш ~/.bashrc или ~/.zshrc файл, чтобы сделать изменение постоянным.

3. Или используйте полный путь к easyjson:

```bash
$(go env GOPATH)/bin/easyjson -all click.go
```
