package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"myblog/model"
	"strconv"
)

type Info struct {
	Msg     string
	IsStar  int64
	StarNum int64
}

// 用户点赞
func SetStarArticle(p *model.ParamArticleSetStart) (i Info, err error) {
	key := getRedisKey(KeyArticleStartSetPF + strconv.Itoa(int(p.ArticleId)))
	exist := rdb.SIsMember(key, p.UserID).Val()
	if exist { //如果已经点赞，取消点赞
		fmt.Println("---取消点赞---")
		if err = rdb.SRem(key, p.UserID).Err(); err != nil {
			fmt.Printf("取消点赞错误---%#v\n", err)
		}
		i.Msg = "取消点赞"
		i.IsStar = 0
		i.StarNum = rdb.SCard(key).Val()

	} else {
		fmt.Println("---点赞---")
		if err = rdb.SAdd(key, p.UserID).Err(); err != nil {
			fmt.Printf("点赞错误---%#v\n", err)
		}
		i.Msg = "点赞"
		i.IsStar = 1
		i.StarNum = rdb.SCard(key).Val()
	}
	return
}

// 文章列表查询用户是否对文章进行了点赞操作
func CheckIsStar(id int64, list []*model.Article) (starList []int64) {
	pipeline := rdb.Pipeline()
	for _, article := range list {
		pipeline.SIsMember(getRedisKey(KeyArticleStartSetPF+strconv.Itoa(int(article.ID))), id)
	}
	cmders, _ := pipeline.Exec()
	for _, cmder := range cmders {
		var isStar int64
		num := cmder.(*redis.BoolCmd).Val()
		if num == false {
			isStar = 0
		} else {
			isStar = 1
		}
		starList = append(starList, isStar)
	}
	return
}
func CheckArticleDetailIsStar(id, articleID int64) bool {
	return rdb.SIsMember(getRedisKey(KeyArticleStartSetPF+strconv.Itoa(int(articleID))), id).Val()
}
func GetArticleListStarNum(ids []int64) (starNums []int64) {
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		pipeline.SCard(getRedisKey(KeyArticleStartSetPF + strconv.Itoa(int(id))))
	}
	cmders, _ := pipeline.Exec()
	for _, cmder := range cmders {
		num := cmder.(*redis.IntCmd).Val()

		starNums = append(starNums, num)
	}
	return
}
func GetArticleDetailStarNum(articleId int64) (starNums int64) {
	return rdb.SCard(getRedisKey(KeyArticleStartSetPF + strconv.Itoa(int(articleId)))).Val()

}
