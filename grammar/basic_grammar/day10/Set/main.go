package main

import (
	"fmt"

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

	//1、SAdd添加集合元素
	// 添加100到集合中
	// err := rdb.SAdd(ctx, "key", 100).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// // 将100,200,300添加到集合中
	// rdb.SAdd(ctx, "key", 100, 200, 300)

	//--------------------------------------------------------------------------
	//2、SCard获取集合元素个数
	// size, err := rdb.SCard(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(size)

	//--------------------------------------------------------------------------
	//3、SIsMember判断元素是否在集合中
	// ok, _ := rdb.SIsMember(ctx, "key", 100).Result()
	// if ok {
	// 	fmt.Println("集合包含指定元素")
	// }

	//--------------------------------------------------------------------------
	//4、SMemebers获取集合中所有的元素
	// es, _ := rdb.SMembers(ctx, "key").Result()
	// // 返回的es是string数组
	// fmt.Println(es)

	//--------------------------------------------------------------------------
	//5、SRem删除集合元素
	// 删除集合中的元素100
	rdb.SRem(ctx, "key", 100)

	// 删除集合中的元素200和300
	rdb.SRem(ctx, "key", 200, 300)

	//--------------------------------------------------------------------------
	//6、SPop、SPopN
	// 随机返回集合中的一个元素，并且删除这个元素
	val, _ := rdb.SPop(ctx, "key").Result()
	fmt.Println(val)

	// 随机返回集合中的5个元素，并且删除这些元素
	vals, _ := rdb.SPopN(ctx, "key", 5).Result()
	fmt.Println(vals)
}
