package redis

const (
	Prefix               = "myblog:"  // 项目前缀
	KeyArticleStartSetPF = "article:" //记录每个文章的点赞的id
)

func getRedisKey(key string) string {

	return Prefix + key
}
