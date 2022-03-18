package models

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
//func (t *Article) BeforeCreate(tx *gorm.DB) (err error) {
//	t.CreatedOn = time.Now().Unix()
//	return
//}
//func (t *Article) BeforeUpdate(tx *gorm.DB) (err error) {
//	t.ModifiedOn = time.Now().Unix()
//	return
//}


func ExistArticleByID(id int) bool {
	var article Article
	gormDB.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface {}) (count int64){
	gormDB.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
	gormDB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	gormDB.Where("id = ?", id).First(&article)
	gormDB.Model(&article).Association("tag_id").Find(&Tag{})

	return
}

func EditArticle(id int, data interface {}) bool {
	gormDB.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface {}) bool {
	gormDB.Create(&Article {
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	gormDB.Where("id = ?", id).Delete(Article{})

	return true
}