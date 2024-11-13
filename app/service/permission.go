package service

import (
	"isme-go/app/model"
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type Permission struct{}

// 根据角色id获取资源列表
func (*Permission) GetListByIds(ids []int, enable bool) []response.Permission {

	permissions := make([]response.Permission, 0)

	query := dal.Gorm.Model(&model.Permission{})

	if len(ids) > 0 {
		query = query.Where("id in ?", ids)
	}

	if enable {
		query = query.Where("enable = ?", 1)
	}

	query.Order("`order`").Scan(&permissions)

	return permissions
}

// 获取资源列表
func (*Permission) GetList(enable bool) []response.Permission {

	permissions := make([]response.Permission, 0)

	query := dal.Gorm.Model(&model.Permission{})

	if enable {
		query = query.Where("enable = ?", 1)
	}

	query.Order("`order`").Scan(&permissions)

	return permissions
}

// 资源列表转资源树
func (p *Permission) ListToTree(list []response.Permission, parentId int) []response.PermissionTree {

	tree := make([]response.PermissionTree, 0)

	if len(list) <= 0 {
		return tree
	}

	for _, item := range list {
		if item.ParentId == parentId {
			tree = append(tree, response.PermissionTree{
				Permission: item,
				Children:   p.ListToTree(list, item.Id),
			})
		}
	}

	return tree
}

// 根据路由地址和请求方法获取资源id
func (*Permission) GetDetailByPathAndMethod(path string, method string) response.Permission {

	var permission response.Permission

	query := dal.Gorm.Model(&model.Permission{}).Where("path = ?", path)

	if method != "" {
		query = query.Where("method = ?", method)
	}

	query.Take(&permission)

	return permission
}

// 删除资源
func (*Permission) Delete(id int) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.Permission{}).Where("id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	if err := query.Model(&model.RolePermissionsPermission{}).Where("permission_id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	return query.Commit().Error
}

// 修改资源
func (*Permission) Update(param request.PermissionUpdate) error {
	return dal.Gorm.Model(&model.Permission{}).Where("id = ?", param.Id).Updates(param).Error
}

// 添加资源
func (*Permission) Add(param request.PermissionAdd) error {
	return dal.Gorm.Model(&model.Permission{}).Select("enable", "order", "show", "keep_alive").Create(&model.Permission{
		Name:        param.Name,
		Code:        param.Code,
		Type:        param.Type,
		ParentId:    param.ParentId,
		Path:        param.Path,
		Redirect:    param.Redirect,
		Icon:        param.Icon,
		Component:   param.Component,
		Layout:      param.Layout,
		KeepAlive:   param.KeepAlive,
		Method:      param.Method,
		Description: param.Description,
		Show:        param.Show,
		Enable:      param.Enable,
		Order:       param.Order,
	}).Error
}

// 获取权限按钮
func (*Permission) GetButtons(parentId int) []response.Permission {

	permissions := make([]response.Permission, 0)

	dal.Gorm.Model(&model.Permission{}).Where("parent_id = ?", parentId).Where("type = ?", "BUTTON").Scan(&permissions)

	return permissions
}
