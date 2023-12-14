package main

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func main() {

	//--------------------------------------------------------------------------
	//1、HSet根据key和field字段设置，field字段的值
	// err := rdb.HSet("user_1", "username", "zhangsan").Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//2、HGet根据key和field字段，查询field字段的值
	//user_1是hash key,username是字段名, zhangsan是字段值
	// username, err := rdb.HGet("user_1", "username").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("所设定的值为:", username)

	//--------------------------------------------------------------------------
	//3、HGetAll根据key查询所有字段和值
	// data, err := rdb.HGetAll("user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("所有的值为:", data)

	//--------------------------------------------------------------------------
	//4、HIncrBy根据key和field字段，累加字段的数值
	// count, err := rdb.HIncrBy("user_1", "count", 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(count)

	//--------------------------------------------------------------------------
	//5、HKeys根据key返回所有字段名
	// keys, err := rdb.HKeys("user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("所有字段名:", keys)

	//--------------------------------------------------------------------------
	//6、HLen根据key，查询hash的字段数量
	// size, err := rdb.HLen("user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("字段数量为:", size)

	//--------------------------------------------------------------------------
	//7、HMGet根据key和多个字段名，批量查询多个hash字段值
	// vals, err := rdb.HMGet("user_1", "count", "username").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("多个值:", vals)

	//--------------------------------------------------------------------------
	//8、HMSet根据key和多个字段名和字段值，批量设置hash字段值
	// data := make(map[string]interface{})
	// data["id"] = 1
	// data["username"] = "lisi"

	// //一次性保存多个hash字段值
	// err := rdb.HMSet("key", data).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//9、HSetNX若field字段不存在，则设置hash字段值
	// err := rdb.HSetNX("key", "id", 100).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//10、HDel 根据key和字段名，删除hash字段，支持批量删除hash字段
	//删除一个字段
	//rdb.HDel("key","id")
	//删除多个字段
	//rdb.HDel("key", "id", "username")

	//--------------------------------------------------------------------------
	//11、HExists检测hash字段名是否存在
	// err := rdb.HExists("key", "id").Err()
	// if err != nil {
	// 	panic(err)
	// }
}
