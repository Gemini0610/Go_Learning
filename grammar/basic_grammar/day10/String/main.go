package main

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //默认DB
	})
}

func main() {
	//context上下文，一个跨 API 和进程用来传递截止日期、取消信号和请求范围的值的接口。
	//Background 返回一个非 nil 但空的上下文
	//TODO 返回一个非 nil 但空的上下文。
	//ctx := context.Background()

	//--------------------------------------------------------------------------
	//1、Sets
	//func (c cmdable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	//设定ctx,
	// err := rdb.Set("goredistestkey", "goredistestvalue", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//2、Get
	//func (c cmdable) Get(ctx context.Context, key string) *StringCmd
	//Result源码
	//func (cmd *StringCmd) Result() (string, error) {
	//return cmd.Val(), cmd.err}
	// val, err := rdb.Get("goredistestkey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("redis的值:", val)

	//--------------------------------------------------------------------------

	// result, err := rdb.Do(context.TODO(), "get", "goredistestkey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("redis22的值:", result.(string))

	//---------------------------------------------------------------------------
	//3、GetSet 设置一个key的值，并返回这个key的旧值
	//func (c cmdable) GetSet(ctx context.Context, key string, value interface{})
	// oldVal, err := rdb.GetSet("goredistestkey", "newvalue").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// //打印key的旧值
	// fmt.Println("key的旧值:", oldVal)

	//--------------------------------------------------------------------------
	//4、SetNX若key不存在，则设置这个key的值
	// err := rdb.SetNX("key2", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// getvalue, err := rdb.Get("key2").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("所设定的新值为:", getvalue)

	//--------------------------------------------------------------------------
	// //5、Mget 批量查询key的值
	// getMoreValue, err := rdb.MGet("key1", "key2").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("批量获取的值:", getMoreValue)

	//--------------------------------------------------------------------------
	//6、Mset批量设定key的值
	// err := rdb.MSet("key1", "value1", "key2", "value2").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// //批量获取
	// vals, err := rdb.MGet("key1", "key2").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("相关值:", vals)

	//--------------------------------------------------------------------------
	//7、Incr,IncrBy 针对一个key的数值进行递增操作

	//Incr函数每次加一
	// val, err := rdb.Incr("key11").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("最新值", val)

	//IncrBy函数可以指定每次递增多少
	// valBy, err := rdb.IncrBy("key11", 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("批次递增新值:", valBy)

	//IncrByFloat函数可以指定每次递增多少，可以递增浮点数
	// valBy, err := rdb.IncrByFloat("key11", 1.1).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("批次浮点递增:", valBy)

	//--------------------------------------------------------------------------
	//8、Decr,DecrBy递减，与7相同

	//--------------------------------------------------------------------------
	//9、Del删除key操作，支持批量删除
	err := rdb.Del("key1", "key2").Err()
	if err != nil {
		panic(err)
	}

	//--------------------------------------------------------------------------
	//10、Expire设定key的过期时间
	//rdb.Expire("key11", 3*time.Second)
}
