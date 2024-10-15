package service

import (
	"gex-api/app/model"
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/pkg/dal"
)

// 菜单
type Menu struct{}

// 创建菜单
func (*Menu) Create(menu *request.CreateMenu) error {
	return dal.Gorm.Model(&model.Menu{}).Create(&model.Menu{
		ParentId:  menu.ParentId,
		Name:      menu.Name,
		Type:      menu.Type,
		Sort:      menu.Sort,
		Path:      menu.Path,
		Component: menu.Component,
		Icon:      menu.Icon,
		Hidden:    menu.Hidden,
		KeepAlive: menu.KeepAlive,
		Redirect:  menu.Redirect,
		Status:    menu.Status,
	}).Error
}

// 更新菜单
func (*Menu) Update(menu *request.UpdateMenu) error {
	return dal.Gorm.Model(&model.Menu{}).Where("id = ?", menu.Id).Updates(&model.Menu{
		ParentId:  menu.ParentId,
		Name:      menu.Name,
		Type:      menu.Type,
		Sort:      menu.Sort,
		Path:      menu.Path,
		Component: menu.Component,
		Icon:      menu.Icon,
		Hidden:    menu.Hidden,
		KeepAlive: menu.KeepAlive,
		Redirect:  menu.Redirect,
		Status:    menu.Status,
	}).Error
}

// 删除菜单
func (*Menu) DeleteById(id int) error {
	return dal.Gorm.Model(&model.Menu{}).Where("id = ?", id).Delete(nil).Error
}

// 菜单列表
func (*Menu) GetList(param *request.QueryListMenu) ([]*response.MenuList, int) {

	var count int64
	list := make([]*response.MenuList, 0)

	query := dal.Gorm.Model(&model.Menu{})

	if param.Name != "" {
		query.Where("name like ?", "%"+param.Name+"%")
	}

	if param.Path != "" {
		query.Where("path like ?", "%"+param.Path+"%")
	}

	if param.Component != "" {
		query.Where("component like ?", "%"+param.Component+"%")
	}

	if param.IsPaging {
		query.Count(&count).Limit(param.Size).Offset((param.Page - 1) * param.Size)
	}

	query.Scan(&list)

	return list, int(count)
}

// 菜单详情
func (*Menu) GetDetailById(id int) *response.MenuDetail {

	var detail *response.MenuDetail

	dal.Gorm.Model(&model.Menu{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 菜单列表
func (*Menu) GetListByIds(ids []int) []*response.MenuList {

	list := make([]*response.MenuList, 0)

	query := dal.Gorm.Model(&model.Menu{})

	if len(ids) > 0 {
		query.Where("id in ?", ids)
	}

	query.Scan(&list)

	return list
}

// 菜单列表转菜单树
func (t *Menu) ListToTree(list []*response.MenuList, parentId int) []*response.MenuTree {

	tree := make([]*response.MenuTree, 0)

	if len(list) <= 0 {
		return tree
	}

	for _, item := range list {
		if item.ParentId == parentId {
			tree = append(tree, &response.MenuTree{
				MenuList: item,
				Children: t.ListToTree(list, item.Id),
			})
		}
	}

	return tree
}

// 检查是否包含下级菜单
func (*Menu) IsExistChildren(id int) bool {

	var detail *response.MenuDetail

	dal.Gorm.Model(&model.Menu{}).Where("parent_id = ?", id).Take(&detail)

	return detail.Id > 0
}
