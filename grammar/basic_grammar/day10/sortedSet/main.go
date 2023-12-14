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

	//1、ZAdd添加一个或者多个元素到集合，如果元素已经存在则更新分数
	// 添加一个集合元素到集合中， 这个元素的分数是2.5，元素名是zhangsan
	// err := rdb.ZAdd(ctx, "key", &redis.Z{Score: 2.5, Member: "zhangsan"}).Err()
	// if err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------
	//2、ZCard返回集合元素个数
	// size, err := rdb.ZCard(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(size)

	//--------------------------------------------------------------------------
	//3、ZCount统计某个分数范围内的元素个数
	// 返回： 1<=分数<=5 的元素个数, 注意："1", "5"两个参数是字符串
	// size, err := rdb.ZCount(ctx, "key", "1", "5").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(size)

	// // 返回： 1<分数<=5 的元素个数
	// // 说明：默认第二，第三个参数是大于等于和小于等于的关系。
	// // 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
	// size, err := rdb.ZCount(ctx, "key", "(1", "5").Result()

	//--------------------------------------------------------------------------
	//4、ZIncrBy增加元素的分数
	// 给元素zhangsan，加上2分
	// rdb.ZIncrBy(ctx, "key", 2, "zhangsan")

	//--------------------------------------------------------------------------
	//5、ZRange，ZRevRange返回集合中某个索引范围的元素，根据分数从小到大排序
	// 返回从0到-1位置的集合元素， 元素按分数从小到大排序
	// 0到-1代表则返回全部数据
	// vals, err := rdb.ZRange(ctx, "key", 0, -1).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, val := range vals {
	// 	fmt.Println(val)
	// }

	//--------------------------------------------------------------------------
	//6、ZRangeByScore根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页
	// 初始化查询条件， Offset和Count用于分页
	// op := redis.ZRangeBy{
	// 	Min:    "2",  // 最小分数
	// 	Max:    "10", // 最大分数
	// 	Offset: 0,    // 类似sql的limit, 表示开始偏移量
	// 	Count:  5,    // 一次返回多少数据
	// }

	// vals, err := rdb.ZRangeByScore(ctx, "key", &op).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, val := range vals {
	// 	fmt.Println(val)
	// }

	//--------------------------------------------------------------------------
	//8、ZRangeByScoreWithScores用法跟上一个一样，区别是除了返回集合元素，同时也返回元素对应的分数
	// 初始化查询条件， Offset和Count用于分页
	// op := redis.ZRangeBy{
	// 	Min:    "2",  // 最小分数
	// 	Max:    "10", // 最大分数
	// 	Offset: 0,    // 类似sql的limit, 表示开始偏移量
	// 	Count:  5,    // 一次返回多少数据
	// }

	// vals, err := rdb.ZRangeByScoreWithScores(ctx, "key", &op).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, val := range vals {
	// 	fmt.Println(val.Member) // 集合元素
	// 	fmt.Println(val.Score)  // 分数
	// }

	//--------------------------------------------------------------------------
	//9、ZRem删除集合元素
	// 删除集合中的元素zhangsan
	// rdb.ZRem(ctx, "key", "zhangsan")

	// // 删除集合中的元素zhangsan和zhangsan1
	// // 支持一次删除多个元素
	// rdb.ZRem(ctx, "key", "zhangsan", "zhangsan1")

	//--------------------------------------------------------------------------
	//10、ZRemRangeByRank根据索引范围删除元素
	// 集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
	// 这里相当于删除最低分的几个元素
	// rdb.ZRemRangeByRank(ctx, "key", 0, 5)

	// // 位置参数写成负数，代表从高分开始删除。
	// // 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
	// rdb.ZRemRangeByRank(ctx, "key", -1, -2)

	//--------------------------------------------------------------------------
	//11、ZRemRangeByScore根据分数范围删除元素
	// 删除范围： 2<=分数<=5 的元素
	rdb.ZRemRangeByScore(ctx, "key", "2", "5")

	// 删除范围： 2<=分数<5 的元素
	rdb.ZRemRangeByScore(ctx, "key", "2", "(5")

	//--------------------------------------------------------------------------

}
