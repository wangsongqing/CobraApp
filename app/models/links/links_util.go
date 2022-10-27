package links

import (
	"CobraApp/pkg/database"
)

// GetLink []Links 查询多个
func GetLink() (LinksModel []Links) {
	database.DB.Find(&LinksModel)
	return
}

func Get(ID int64) (link Links) {
	database.DB.Where("id", ID).First(&link)
	return
}
