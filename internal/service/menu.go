package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问菜单数据
type Menu struct{}

// 创建菜单
func (t *Menu) CreateMenu(menu *model.Menu) error {

	err := db.GormClient.Model(&model.Menu{}).Create(&menu).Error

	return err
}

// 更新菜单
func (t *Menu) UpdateMenu(menu *model.Menu) error {

	err := db.GormClient.Model(&model.Menu{}).Where("id = ?", menu.Id).Updates(&menu).Error

	return err
}

// 删除菜单
func (t *Menu) DeleteMenu(id int) error {

	err := db.GormClient.Model(&model.Menu{}).Where("id = ?", id).Delete(&model.Menu{}).Error

	return err
}

// 获取菜单列表
func (t *Menu) GetMenuListByPage(page, size int) ([]*model.Menu, int) {

	var (
		list []*model.Menu
		count int64
	)

	query := db.GormClient.Model(&model.Menu{})

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取菜单详情
func (t *Menu) GetMenuById(id int) *model.Menu {

	var detail *model.Menu

	db.GormClient.Model(&model.Menu{}).Where("id = ?", id).Take(&detail)

	return detail
}