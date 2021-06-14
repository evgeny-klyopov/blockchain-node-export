# 2. cli выгрузка в БД

#### Перед запуском необходимо [создать таблицу](./../migrations/20210608001034_create_bne_table.up.sql)

Пример запуска:
```shell
bne --mode="simple" --blockchain-rpc-dsn="http://user:password@localhost:1234/wallet/wallet.dat" --output-type=database --blockchain-database-dsn="postgresql://user:password@localhost:1235/dbName/tableName/?charset=utf8"
```