# vmmgr Imacon

imacon is Image Controller.

### MariaDB(備忘録)
#### root以外でも操作可能にする方法
```
$ sudo mysql -u root

mysql> USE mysql;
mysql> UPDATE user SET plugin='mysql_native_password' WHERE User='root';
mysql> FLUSH PRIVILEGES;
mysql> exit;
```