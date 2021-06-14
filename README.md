## Blockchain node export
Скрипт выгрузки транзакций с ноды. В данный момент доступно только BTC-нод.

### [Install](./docs/INSTALL.md)

### Mode
1. [export to file (sql, csv)](./docs/EXPORT_TO_FILE.md)
1. [export to database](./docs/EXPORT_TO_DATABASE.md)
1. [worker](./docs/SERVICE.md) - systemd, supervisor (поддерживает первые два режима)

### Options
* ```--mode=simple``` - режим работы, доступны *simple* и *worker* (default - simple)
* ```--worker-interval=600``` - интервал запуска в секундах (default - 600)
* ```--blockchain=BTC``` - вид блокчейна (default - BTC)
* ```--day=2021-06-08``` - с какой даты выгружать (default - дата запуска). После запуска с помощью [сервиса](https://chain.api.btc.com/v3/block/date/2021-06-08) выбирается блок, с которого стартует выгрузка. В режиме воркера после все последующие вызовы будут вызываться от last_block полученного с ноды при последнем удачном запуске.
* ```--output-type=file``` - режим вывода, доступны *file* и *database* (default - file). 
* ```--output-file-extension=csv``` - расширение выгружаемого файла, доступны *csv* и *sql* (default - csv).
* ```--output-directory=./output``` - директория в которую будет сохранена выгрузка (default - директория запуска + /output)
* ```--blockchain-rpc-dsn="http://user:password@localhost:1234/wallet/wallet.dat"``` - rpc подключение к ноде
* ```--blockchain-rpc-timeout=120``` - таймаут соединения в секундах (default - 120)
* ```--blockchain-database-dsn="postgresql://user:password@localhost:1235/dbName/tableName/?charset=utf8"``` - подключение к базе данных
* ```--blockchain-rpc-use-mock=true``` - режим использования моков (default - false)
* ```--blockchain-rpc-mock-directory=".../mocks""``` - путь до моков, более подробно в репозитории [bitcoin-rpc](https://github.com/evgeny-klyopov/bitcoin-rpc)