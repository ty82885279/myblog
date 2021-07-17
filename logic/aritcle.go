package logic

import (
	"myblog/dao/mysql"
	"myblog/dao/redis"
	"myblog/model"
)

// 获取某个分类下的所有文章
func GetArticlesByCategoryId(p model.ParamArticleList) (list []*model.Article, total int64) {
	articleIds := make([]int64, 0)

	list, total = mysql.GetArticlesByCategoryId(p)
	// 查询文章点赞数
	for _, article := range list {
		articleIds = append(articleIds, article.ID)
	}
	starS := redis.GetArticleListStarNum(articleIds)
	for i, article := range list {
		article.PraiseNums = starS[i]
	}
	// 如果登录，查询用户是否点赞
	if p.UserId != 0 { //需要查询用户的点赞情况
		stars := redis.CheckIsStar(p.UserId, list)
		for i, article := range list {
			article.IsStar = stars[i]
		}
	}

	return
}

// 查询文章是否存
func CheckArticleExist(articleID int64) bool {

	return mysql.CheckArticleExist(articleID)
}

// 文章详情
func GetArticleDetail(p model.ParamArticleDetail) (article *model.Article, err error) {

	article, err = mysql.GetArticleDetail(p)
	if err != nil {
		return nil, err
	}
	if p.UserID != 0 { // 需要查询登录用户是否点赞该文章
		isStar := redis.CheckArticleDetailIsStar(p.UserID, p.ArticleID)
		if isStar == true {
			article.IsStar = 1
		}
	}
	starNum := redis.GetArticleDetailStarNum(article.ID)
	article.PraiseNums = starNum
	return
}

// 添加文章
func CreateArticle(p *model.ParamAddArticle) error {
	return mysql.CreateArticle(p)
}

// 编辑文章
func EditArticle(article *model.Article) error {
	return mysql.EditArticle(article)
}

// 删除文章
func DeleteArticle(id int64) error {
	return mysql.DeleteArticle(id)
}

func SetStarArticle(p *model.ParamArticleSetStart) (redis.Info, error) {
	return redis.SetStarArticle(p)
}
