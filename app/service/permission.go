package service

import (
	"gex-api/app/model"
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/pkg/dal"
)

// 权限
type Permission struct{}

// 创建权限
func (*Permission) Create(permission *request.CreatePermission) error {
	return dal.Gorm.Model(&model.Permission{}).Create(&model.Permission{
		Name:      permission.Name,
		GroupName: permission.GroupName,
		Path:      permission.Path,
		Method:    permission.Method,
		Status:    permission.Status,
	}).Error
}

// 更新权限
func (*Permission) Update(permission *request.UpdatePermission) error {
	return dal.Gorm.Model(&model.Permission{}).Where("id = ?", permission.Id).Updates(&model.Permission{
		Name:      permission.Name,
		GroupName: permission.GroupName,
		Path:      permission.Path,
		Method:    permission.Method,
		Status:    permission.Status,
	}).Error
}

// 删除权限
func (*Permission) DeleteById(id int) error {
	return dal.Gorm.Model(&model.Permission{}).Where("id = ?", id).Delete(nil).Error
}

// 权限列表
func (*Permission) GetList(param *request.QueryListPermission) ([]*response.PermissionList, int) {

	var count int64
	list := make([]*response.PermissionList, 0)

	query := dal.Gorm.Model(&model.Permission{})

	if param.Name != "" {
		query.Where("name like ?", "%"+param.Name+"%")
	}

	if param.GroupName != "" {
		query.Where("group_name like ?", "%"+param.GroupName+"%")
	}

	if param.Path != "" {
		query.Where("path like ?", "%"+param.Path+"%")
	}

	if param.Method != "" {
		query.Where("method = ?", param.Method)
	}

	if param.Status > 0 {
		query.Where("status = ?", param.Status)
	}

	if param.IsPaging {
		query.Count(&count).Limit(param.Size).Offset((param.Page - 1) * param.Size)
	}

	query.Scan(&list)

	return list, int(count)
}

// 权限详情
func (*Permission) GetDetailById(id int) *response.PermissionDetail {

	var detail *response.PermissionDetail

	dal.Gorm.Model(&model.Permission{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 权限列表
func (*Permission) GetListByIds(ids []int) []*response.PermissionList {

	list := make([]*response.PermissionList, 0)

	dal.Gorm.Model(&model.Permission{}).Where("id in ?", ids).Scan(&list)

	return list
}

// 权限是否存在
func (*Permission) GetDetailByPathAndMethod(path, method string) *response.PermissionDetail {

	var detail *response.PermissionDetail

	dal.Gorm.Model(&model.Permission{}).Where("path = ? and method = ?", path, method).Take(&detail)

	return detail
}
