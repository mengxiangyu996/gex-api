package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问菜单数据
type Menu struct{}

type MenuTree struct {
	Id       int         `json:"id"`
	ParentId int         `json:"parentId"`
	Name     string      `json:"name"`
	Children []*MenuTree `json:"children"`
}

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
		list  []*model.Menu
		count int64
	)

	query := db.GormClient.Model(&model.Menu{}).Order("id desc")

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取菜单列表
func (t *Menu) GetMenuTree() []*MenuTree {

	var list []*model.Menu

	db.GormClient.Model(&model.Permission{}).Order("id desc").Find(&list)

	return t.MenuListToTree(list, 0)
}

// 获取菜单详情
func (t *Menu) GetMenuById(id int) *model.Menu {

	var detail *model.Menu

	db.GormClient.Model(&model.Menu{}).Where("id = ?", id).Take(&detail)

	return detail
}

// list 转 tree
func (t *Menu) MenuListToTree(menuList []*model.Menu, parentId int) []*MenuTree {

	var menuTree []*MenuTree

	for _, menu := range menuList {
		if menu.ParentId == parentId {

			node := &MenuTree{
				Id:       menu.Id,
				ParentId: menu.ParentId,
				Name:     menu.Name,
			}

			node.Children = t.MenuListToTree(menuList, menu.Id)
			menuTree = append(menuTree, node)
		}
	}

	return menuTree
}
