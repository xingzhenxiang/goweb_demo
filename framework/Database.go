package framework

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库的配置
const (
	username   = "root"
	password   = "123456"
	ip         = "192.168.0.170"
	port       = "3306"
	dbName     = "blog"
	driverName = "mysql"
)

//DB数据库连接池
var DB *sql.DB

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("database connect error")
	}
	fmt.Println("database connect success")
}

func CreateTable() {
	userTable := "CREATE TABLE IF NOT EXISTS `user`(" +
		"`id` INT UNSIGNED AUTO_INCREMENT," +
		"`username` VARCHAR(20) NOT NULL," +
		"`password` VARCHAR(40) NOT NULL," +
		"`create_time` DATETIME," +
		"PRIMARY KEY ( `id` )," +
		"UNIQUE KEY `idx_username` (`username`)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	topicTable := "CREATE TABLE IF NOT EXISTS  `topics` (                                                                 " +
		"  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,                                                 " +
		"  `title` varchar(255) NOT NULL,                                                                 " +
		"  `topics_preview` varchar(255) NOT NULL,                                                        " +
		"  `topics_info` text,                                                                            " +
		"  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,                                    " +
		"  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,        " +
		"  PRIMARY KEY (`id`)                                                                             " +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	_, err := DB.Exec(userTable)
	if err != nil {
		fmt.Println("CreateUserTable error")
	}
	fmt.Println("CreateUserTable success")
	_, err = DB.Exec(topicTable)
	if err != nil {
		fmt.Println("CreateTopicTable error")
		fmt.Println(err)
	}
	fmt.Println("CreateTopicTable success")
}
