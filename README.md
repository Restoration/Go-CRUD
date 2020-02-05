# CRUD application with Go lang
CRUD application with Go lang


Execute bellow SQL.
```
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `city` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
```

Get package in your project folder.
```
$ go get -u github.com/go-sql-driver/mysql
```

## Reference
- [Example of Golang CRUD using MySQL from scratch](https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html)
