# 3. воркер

#### Если воркер будет запущен в режиме экспорта в БД - перед запуском необходимо [создать таблицу](./../migrations/20210608001034_create_bne_table.up.sql)

[Пример сервиса](./../service/blockchain_node_export.service)

```shell
cp blockchain_node_export.service /etc/systemd/system
systemctl enable blockchain_node_export
systemctl start blockchain_node_export
systemctl status blockchain_node_export
```

#### Переменных окружения
```shell
$BNE_BLOCKCHAIN_RPC_DSN
$BNE_BLOCKCHAIN_DATABASE_DSN
```