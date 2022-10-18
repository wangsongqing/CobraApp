package links

import "CobraApp/pkg/database"

// GetLink []Links 查询多个
func GetLink() (LinksModel []Links) {
	database.DB.Find(&LinksModel)
	return
}
