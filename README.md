# Gendiff

Утилита для сравнения конфигурационных файлов (JSON, YAML).

### Hexlet tests and linter status:
[![Actions Status](https://github.com/Maga-Gazdiev/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Maga-Gazdiev/go-project-244/actions)

## Установка

```bash
make install
```

## Использование

```bash
gendiff [options] <filepath1> <filepath2>
```

### Опции

- `-f, --format <format>` - формат вывода (stylish, plain, json). По умолчанию: stylish

## Примеры

### Формат stylish (по умолчанию)

```bash
$ gendiff file1.json file2.json
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

### Формат plain

```bash
$ gendiff --format plain file3.json file4.json
Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]
```

### Формат json

```bash
$ gendiff --format json file1.json file2.json
[
  {
    "key": "follow",
    "status": "remove",
    "oldValue": false
  },
  {
    "key": "host",
    "status": "unchanged",
    "oldValue": "hexlet.io"
  },
  {
    "key": "proxy",
    "status": "remove",
    "oldValue": "123.234.53.22"
  },
  {
    "key": "timeout",
    "status": "changed",
    "oldValue": 50,
    "newValue": 20
  },
  {
    "key": "verbose",
    "status": "added",
    "newValue": true
  }
]
```

## Поддерживаемые форматы

- JSON (.json)
- YAML (.yml, .yaml)

## Архитектура

Проект разделен на следующие слои:

- **parser** - парсинг файлов (JSON, YAML)
- **builder** - построение дерева различий
- **formatters** - форматирование вывода (stylish, plain, json)
- **model** - модели данных
