# Notifier

Simple app that can notify you about directory updates on Telegram

## Configuration

To configure the app you need to create `config.yaml` nearby your executable

### Config structure

```
telegram:
  chatId: 123 # open https://t.me/myidbot and call /getid
  botToken: 456:ABC # get it from https://t.me/BotFather
directories:
  - ./data # a list of watched dirs
filter: ^(.*).mkv$ # regex filename filter
```

## Development

### Running from source

`go run .`

### Building a binary

`go build -o /app/notifier`