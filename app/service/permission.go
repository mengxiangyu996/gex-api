package service

import (
	"isme-go/app/model"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type Permission struct{}

// 根据角色id获取资源列表
func (*Permission) GetListByIds(ids []int, enable bool) []response.Permission {

	permissions := make([]response.Permission, 0)

	query := dal.Gorm.Model(&model.Permission{}).Where("id in ?", ids)

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
