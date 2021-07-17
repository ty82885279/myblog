package mysql

import (
	"gorm.io/gorm"
	"myblog/model"
)

// 获取某个分类下的所有文章
func GetArticlesByCategoryId(p model.ParamArticleList) (list []*model.Article, total int64) {

	// 获取总数
	if err := db.Debug().Table("t_article").Select("id").Where("cid =?", p.CID).Find(&list).Count(&total).Error; err != nil {
		panic(err)
	}
	// 根据cid查询分类
	cate := new(model.Category)
	if err := db.Debug().Table("t_category").Select("cid", "cname").Where("cid = ?", p.CID).Find(cate).Error; err != nil {
		panic(err)
	}
	// 获取分类下的文章
	if err := db.Debug().Table("t_article").Omit("cid").Where("cid = ?", p.CID).Limit(p.Size).Offset((p.Page - 1) * p.Size).Find(&list).Error; err != nil {
		panic(err)
	}
	// 将分类赋值到文章里
	for _, article := range list {
		article.Category = *cate
	}
	return
}

// 查询文章是否存
func CheckArticleExist(articleID int64) bool {

	var count int64
	if err := db.Debug().Model(&model.Article{}).Where("id = ?", articleID).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		return false
	}
	return true
}

// 文章详情
func GetArticleDetail(p model.ParamArticleDetail) (article *model.Article, err error) {

	article = new(model.Article)
	exist := CheckArticleExist(p.ArticleID)
	if !exist {
		err = ErrorArtiCleIDNotExist
		return
	}
	if err = db.Debug().Model(&model.Article{}).Where("id = ?", p.ArticleID).Find(article).Error; err != nil {
		panic(err)
	}

	// 根据cid查询分类
	cate := new(model.Category)
	if err := db.Debug().Table("t_category").Select("cid", "cname").Where("cid = ?", article.CID).Find(cate).Error; err != nil {
		panic(err)
	}
	article.Category = *cate
	if err = IncreaseReadNum(p.ArticleID); err != nil {
		panic(err)
	}
	return
}

// 文章评论数+1
func IncreaseReadNum(article int64) (err error) {
	if err = db.Debug().Model(&model.Article{}).Where("id = ?", article).Update("read_nums", gorm.Expr("read_nums + ?", 1)).Error; err != nil {
		panic(err)
	}
	return
}

// 添加文章
func CreateArticle(p *model.ParamAddArticle) (err error) {
	article := model.Article{
		Title:   p.Title,
		Content: p.Content,
		CID:     p.Cid,
	}

	if err = db.Debug().Create(&article).Error; err != nil {
		return err
	}

	return
}

// 编辑文章
func EditArticle(article *model.Article) (err error) {

	if err = db.Debug().Updates(article).Error; err != nil {
		return err
	}

	return
}

// 删除文章
func DeleteArticle(id int64) (err error) {

	err = db.Debug().Table("t_article").Where("id = ?", id).Update("is_del", 1).Error
	return
}

// 更新文章评论数+1
func IncreaseCommentNum(articleID int64) (err error) {

	if err = db.Debug().Model(&model.Article{}).Where("id = ?", articleID).Update("comment_nums", gorm.Expr("comment_nums + ?", 1)).Error; err != nil {
		panic(err)
	}
	return
}
