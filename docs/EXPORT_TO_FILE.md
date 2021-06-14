# 1. cli выгрузка в файл (sql, csv)

Примеры запуска:
* выгрузка в csv
```shell
bne --mode="simple" --blockchain-rpc-dsn="http://user:password@localhost:1234/wallet/wallet.dat" --output-type=file --output-file-extension="csv"
```
* выгрузка в sql
```shell
bne --mode="simple" --blockchain-rpc-dsn="http://user:password@localhost:1234/wallet/wallet.dat" --output-type=file --blockchain-database-dsn="postgresql://user:password@localhost:1235/dbName/tableName/?charset=utf8" --output-file-extension="sql"
```