package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"something-init/log-ini"
	"something-init/mongo-init"
	"time"
)

var DB *gorm.DB

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName("config-dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	//或者直接解析文件
	//env := "dev"
	//viper.SetConfigFile("./config/config-" + env + ".yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
}

func main() {
	//初始化日志（zerolog）
	log_ini.SetLogLevel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//初始化mongo
	mongoClient, err := mongo.InitMongo(ctx)
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(ctx)
	//初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.db"),
	})
	r := rdb.Ping(ctx)
	if r.Err() != nil {
		log.Fatal().Str("addr", viper.GetString("redis.addr")).
			Err(err).Msg("established redis failed.")
		panic(err)
	}
	////初始化mysql
	//db, err := sql.Open("mysql", viper.GetString("mysql.endpoint"))
	//if err != nil {
	//	log.Fatal().Str("endpoint", viper.GetString("mysql.endpoint")).
	//		Err(err).Msg("established mysql failed.")
	//	panic(err)
	//}
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//db.SetConnMaxLifetime(viper.GetDuration("mysql.max_life_time"))
	//db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conn"))
	//db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))
	//defer db.Close()

	//初始化gorm
	mysqldb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       viper.GetString("mysql.endpoint"),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// 使用CreateBatchSize 选项初始化 GORM 时，所有的创建& 关联 INSERT 都将遵循该选项
		CreateBatchSize: 1000,
		// 全局模式：执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		PrepareStmt: true,
		// 表名加前缀和禁用复数表名
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})

	sqlDB, _ := mysqldb.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = mysqldb
	if err != nil {
		panic(err)
	}

	//cols := mongoClient.Database("abc").Collection("test")
	//query := bson.M{
	//	"user": 5,
	//}
	//var user User
	//err = cols.FindOne(ctx, query).Decode(&user)
	//if err != nil {
	//	log.Fatal().Err(err)
	//}
	//fmt.Println(user)

	//sqlc
	//queries := tutorial.New(sqlDB)
	//err = queries.DeleteAuthor(ctx, 5)
	//if err != nil {
	//	panic(err)
	//}
}

type User struct {
	//Id   primitive.ObjectID `json:"id" bson:"_id"`
	User int64
}
