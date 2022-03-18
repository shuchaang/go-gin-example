package models

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	gormDB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface {}) (count int64){
	gormDB.Model(&Tag{}).Where(maps).Count(&count)
	return
}


func ExistingTagByName(name string)bool{
	var tag Tag
	gormDB.Select("id").Where("name=?",name).First(&Tag{})
	return tag.ID>0
}

func AddTag(name string,state int,createBy string)bool{
	gormDB.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createBy,
	})
	return true
}

func ExistTagByID(id int)bool{
	var tag Tag
	gormDB.First(&tag,id)
	return tag.ID>0
}

func EditTag(id int,data interface{}){
	gormDB.Model(&Tag{}).Where("id=?",id).Updates(data)
}

func DeleteTag(id int){
	gormDB.Where("id=?",id).Delete(&Tag{})
}

//
//func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
//	t.CreatedOn = time.Now().Unix()
//	return
//}
//func (t *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
//	t.ModifiedOn = time.Now().Unix()
//	return
//}