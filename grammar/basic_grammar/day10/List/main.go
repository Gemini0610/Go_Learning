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

	//1、LPush从列表左边插入数据
	//插入一个数据
	//func (c cmdable) LPush(ctx context.Context, key string, values ...interface{})
	// rdb.LPush("key", "data1")

	//LPush支持一次插入任意个数据
	// err := rdb.LPush("key", 1, 2, 3, 4, 5, 6).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//2、LPushX跟LPush的区别是仅当列表存在的时候才插入数据

	//--------------------------------------------------------------------------
	//3、RPop从列表的右边删除第一个数据，并返回删除的数据
	// val, err := rdb.RPop("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("值:", val)

	//--------------------------------------------------------------------------
	//4、RPush、RPushX从右边插入相关值
	//插入一个值
	// rdb.RPush("key", "data1")

	// //支持一次插入任意个数据
	// err := rdb.RPush("key", 1, 2, 3, 4, 5, 6).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//5、LPop从列表的左边删除第一个数据，并返回删除的数据
	// val, err := rdb.LPop("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("值:", val)

	//--------------------------------------------------------------------------
	//6、LLen返回列表大小
	// 	val, err := rdb.LLen("key").Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("值:", val)

	//--------------------------------------------------------------------------
	//7、LRange返回列表的一个范围内的数据，也可以返回全部数据
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err := rdb.LRange("key", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

	//--------------------------------------------------------------------------
	//8、LRem删除列表中的数据
	// 从列表左边开始，删除100， 如果出现重复元素，仅删除1次，也就是删除第一个
	// dels, err := rdb.LRem("key", 1, 100).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// // 如果存在多个100，则从列表左边开始删除2个100
	// rdb.LRem("key", 2, 100)

	// // 如果存在多个100，则从列表右边开始删除2个100
	// // 第二个参数负数表示从右边开始删除几个等于100的元素
	// rdb.LRem("key", -2, 100)

	// // 如果存在多个100，第二个参数为0，表示删除所有元素等于100的数据
	// rdb.LRem("key", 0, 100)

	//--------------------------------------------------------------------------
	//9、LIndex根据索引坐标，查询列表中的数据
	// 列表索引从0开始计算，这里返回第6个元素
	// val, err := rdb.LIndex("key", 5).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(val)

	//--------------------------------------------------------------------------
	//9、LInsert在指定位置插入数据
	// err := rdb.LInsert("key", "before", 5, 4).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// // 在列表中 zhangsan 元素的前面插入 欢迎你
	// rdb.LInsert("key", "before", "zhangsan", "欢迎你")

	// // 在列表中 zhangsan 元素的后面插入 2022
	// rdb.LInsert("key", "after", "zhangsan", "2022")
}
