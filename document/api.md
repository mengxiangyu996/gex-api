---
title: 初始接口
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.19"

---

# 初始接口

> v1.0.0

Base URLs:

* <a href="http://127.0.0.1:3000">开发环境: http://127.0.0.1:3000</a>

# Authentication

# 后台

## POST 管理员登录

POST /api/admin/admin/login

> Body Parameters

```json
{
  "username": "string",
  "password": "string"
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|object| no ||none|
|» username|body|string| yes | 用户名|none|
|» password|body|string| yes | 密码|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "token": "eyJJZCI6MSwiRXhwaXJlIjoiMjAyNC0wMS0wNFQwODo0MDoxMi42NDMwODYxKzA4OjAwIn1lOTdhYmQ4NzFlYTI3OTNmZmE5OGY0MzBlZDI2OGMzNQ=="
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 修改管理员密码

POST /api/admin/admin/changePassword

> Body Parameters

```json
{
  "password": "string"
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» password|body|string| yes | 密码|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 创建管理员

POST /api/admin/admin/create

> Body Parameters

```json
{
  "username": "string",
  "nickname": "string",
  "gender": 0,
  "email": "string",
  "phone": "string",
  "avatar": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» username|body|string| yes | 用户名|none|
|» nickname|body|string| no | 昵称|none|
|» gender|body|integer| no | 性别|1-男；2-女|
|» email|body|string| no | 邮箱|none|
|» phone|body|string| no | 手机号|none|
|» avatar|body|string| no | 头像|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 更新管理员

POST /api/admin/admin/update

> Body Parameters

```json
{
  "id": 0,
  "nickname": "string",
  "gender": 0,
  "email": "string",
  "phone": "string",
  "avatar": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 用户id|none|
|» nickname|body|string| no | 昵称|none|
|» gender|body|integer| no | 性别|1-男；2-女|
|» email|body|string| no | 邮箱|none|
|» phone|body|string| no | 手机号|none|
|» avatar|body|string| no | 头像|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 删除管理员

POST /api/admin/admin/delete

> Body Parameters

```json
{
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 用户id|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 管理员列表

GET /api/admin/admin/page

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|page|query|integer| no ||页码|
|size|query|integer| no ||数量|
|username|query|string| no ||用户名|
|nickname|query|string| no ||昵称|
|email|query|string| no ||邮箱|
|phone|query|string| no ||手机号|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "count": 1,
    "list": [
      {
        "id": 1,
        "createTime": "2023-12-28T08:38:05+08:00",
        "updateTime": "2023-12-28T08:38:05+08:00",
        "deleteTime": null,
        "username": "admin",
        "nickname": "超级管理员",
        "gender": 1,
        "email": "",
        "phone": "",
        "password": "",
        "avatar": "",
        "status": 1
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 管理员详情

GET /api/admin/admin/detail

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|query|integer| no ||用户id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "admin": {
      "id": 2,
      "createTime": "2023-12-28T08:40:35+08:00",
      "updateTime": "2023-12-28T08:44:31+08:00",
      "username": "test",
      "nickname": "测试",
      "gender": 1,
      "email": "1008611@163.com",
      "phone": "1008611",
      "avatar": "http://dummyimage.com/100x100",
      "status": 1,
      "roles": [
        {
          "id": 1,
          "name": "管理员"
        },
        {
          "id": 2,
          "name": "测试人员"
        }
      ]
    }
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 角色绑定菜单

POST /api/admin/role/bindMenu

> Body Parameters

```json
{
  "roleId": 0,
  "menuIds": [
    0
  ]
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» roleId|body|integer| yes | 角色id|none|
|» menuIds|body|[integer]| no | 菜单ids|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 角色菜单列表

GET /api/admin/role/menus

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|roleId|query|integer| yes ||角色id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "bindTree": [
      {
        "id": 1,
        "parentId": 0,
        "name": "主页",
        "type": 2,
        "sort": 0,
        "path": "/dashboard",
        "component": "Dashboard",
        "icon": "icon-home",
        "redirect": "nulla aliquip tempor commodo",
        "status": "1",
        "children": null
      },
      {
        "id": 2,
        "parentId": 0,
        "name": "系统管理",
        "type": 1,
        "sort": 0,
        "path": "/system",
        "component": "",
        "icon": "icon-system",
        "redirect": "",
        "status": "1",
        "children": [
          {
            "id": 4,
            "parentId": 2,
            "name": "角色管理",
            "type": 2,
            "sort": 0,
            "path": "/role",
            "component": "role",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          }
        ]
      }
    ],
    "tree": [
      {
        "id": 1,
        "parentId": 0,
        "name": "主页",
        "type": 2,
        "sort": 0,
        "path": "/dashboard",
        "component": "Dashboard",
        "icon": "icon-home",
        "redirect": "nulla aliquip tempor commodo",
        "status": "1",
        "children": null
      },
      {
        "id": 2,
        "parentId": 0,
        "name": "系统管理",
        "type": 1,
        "sort": 0,
        "path": "/system",
        "component": "",
        "icon": "icon-system",
        "redirect": "",
        "status": "1",
        "children": [
          {
            "id": 3,
            "parentId": 2,
            "name": "菜单管理",
            "type": 2,
            "sort": 0,
            "path": "/menu",
            "component": "menu",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 4,
            "parentId": 2,
            "name": "角色管理",
            "type": 2,
            "sort": 0,
            "path": "/role",
            "component": "role",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 5,
            "parentId": 2,
            "name": "权限管理",
            "type": 2,
            "sort": 0,
            "path": "/permission",
            "component": "permission",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 6,
            "parentId": 2,
            "name": "用户管理",
            "type": 2,
            "sort": 0,
            "path": "/admin",
            "component": "admin",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          }
        ]
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 角色权限列表

GET /api/admin/role/permissions

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|roleId|query|integer| yes ||角色id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "bindTree": [
      {
        "id": 1,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "创建权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/create",
        "method": "POST",
        "status": 1
      },
      {
        "id": 2,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "更新权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/update",
        "method": "POST",
        "status": 1
      }
    ],
    "tree": [
      {
        "id": 1,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "创建权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/create",
        "method": "POST",
        "status": 1
      },
      {
        "id": 2,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "更新权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/update",
        "method": "POST",
        "status": 1
      },
      {
        "id": 3,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "删除权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/delete",
        "method": "POST",
        "status": 1
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 角色绑定权限

POST /api/admin/role/bindPermission

> Body Parameters

```json
{
  "roleId": 0,
  "permissionIds": [
    0
  ]
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» roleId|body|integer| yes | 角色id|none|
|» permissionIds|body|[integer]| no | 权限ids|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 创建权限

POST /api/admin/permission/create

> Body Parameters

```json
{
  "name": "string",
  "groupName": "string",
  "path": "string",
  "method": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» name|body|string| no | 权限名称|none|
|» groupName|body|string| no | 组名|none|
|» path|body|string| yes | 路径|none|
|» method|body|string| yes | 请求方式|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 更新权限

POST /api/admin/permission/update

> Body Parameters

```json
{
  "id": 0,
  "name": "string",
  "groupName": "string",
  "path": "string",
  "method": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 权限id|none|
|» name|body|string| no | 权限名称|none|
|» groupName|body|string| no | 组名|none|
|» path|body|string| yes | 路径|none|
|» method|body|string| yes | 请求方式|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 删除权限

POST /api/admin/permission/delete

> Body Parameters

```json
{
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 权限id|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 权限列表

GET /api/admin/permission/page

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|page|query|integer| no ||页码|
|size|query|integer| no ||数量|
|name|query|string| no ||权限名称|
|groupName|query|string| no ||组名|
|path|query|string| no ||路径|
|method|query|string| no ||请求方式|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "count": 1,
    "list": [
      {
        "id": 1,
        "createTime": "2023-12-28T09:27:50+08:00",
        "updateTime": "2023-12-28T09:29:14+08:00",
        "deleteTime": null,
        "name": "创建权限",
        "groupName": "权限组",
        "path": "/api/admin/permission/create",
        "method": "POST",
        "status": 2
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 权限详情

GET /api/admin/permission/detail

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|query|integer| yes ||权限id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "permission": {
      "id": 1,
      "createTime": "2023-12-28T09:27:50+08:00",
      "updateTime": "2023-12-28T09:29:14+08:00",
      "deleteTime": null,
      "name": "创建权限",
      "groupName": "权限组",
      "path": "/api/admin/permission/create",
      "method": "POST",
      "status": 2
    }
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 创建菜单

POST /api/admin/menu/create

> Body Parameters

```json
{
  "parentId": 0,
  "name": "string",
  "type": 0,
  "sort": 0,
  "path": "string",
  "component": "string",
  "icon": "string",
  "redirect": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» parentId|body|integer| no | 父级id|none|
|» name|body|string| yes | 菜单名称|none|
|» type|body|integer| yes | 类型|1-目录；2-菜单；3-按钮|
|» sort|body|integer| no | 排序|none|
|» path|body|string| yes | 路径|none|
|» component|body|string| no | 组件|none|
|» icon|body|string| no | 图标|none|
|» redirect|body|string| no | 重定向地址|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 创建角色

POST /api/admin/role/create

> Body Parameters

```json
{
  "name": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» name|body|string| yes | 角色名称|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 更新角色

POST /api/admin/role/update

> Body Parameters

```json
{
  "id": 0,
  "name": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 角色id|none|
|» name|body|string| yes | 角色名称|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 删除角色

POST /api/admin/role/delete

> Body Parameters

```json
{
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 角色id|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 角色列表

GET /api/admin/role/page

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|page|query|integer| yes ||页码|
|size|query|integer| yes ||数量|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "count": 1,
    "list": [
      {
        "id": 1,
        "createTime": "2023-12-29T11:03:51+08:00",
        "updateTime": "2023-12-29T11:06:35+08:00",
        "deleteTime": null,
        "name": "管理员",
        "status": 1
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 角色详情

GET /api/admin/role/detail

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|query|integer| yes ||角色id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "role": {
      "id": 1,
      "createTime": "2023-12-29T11:03:51+08:00",
      "updateTime": "2023-12-29T11:06:35+08:00",
      "deleteTime": null,
      "name": "管理员",
      "status": 1
    }
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 更新菜单

POST /api/admin/menu/update

> Body Parameters

```json
{
  "id": 0,
  "parentId": 0,
  "name": "string",
  "type": 0,
  "sort": 0,
  "path": "string",
  "component": "string",
  "icon": "string",
  "redirect": "string",
  "status": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 菜单id|none|
|» parentId|body|integer| no | 父级id|none|
|» name|body|string| yes | 菜单名称|none|
|» type|body|integer| yes | 类型|1-目录；2-菜单；3-按钮|
|» sort|body|integer| no | 排序|none|
|» path|body|string| yes | 路径|none|
|» component|body|string| no | 组件|none|
|» icon|body|string| no | 图标|none|
|» redirect|body|string| no | 重定向地址|none|
|» status|body|integer| no | 状态|1-启用；2-禁用|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 删除菜单

POST /api/admin/menu/delete

> Body Parameters

```json
{
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» id|body|integer| yes | 菜单id|none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 菜单列表

GET /api/admin/menu/tree

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "tree": [
      {
        "id": 1,
        "parentId": 0,
        "name": "主页",
        "type": 2,
        "sort": 0,
        "path": "/dashboard",
        "component": "Dashboard",
        "icon": "icon-home",
        "redirect": "nulla aliquip tempor commodo",
        "status": "1",
        "children": null
      },
      {
        "id": 2,
        "parentId": 0,
        "name": "系统管理",
        "type": 1,
        "sort": 0,
        "path": "/system",
        "component": "",
        "icon": "icon-system",
        "redirect": "",
        "status": "1",
        "children": [
          {
            "id": 3,
            "parentId": 2,
            "name": "菜单管理",
            "type": 2,
            "sort": 0,
            "path": "/menu",
            "component": "menu",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 4,
            "parentId": 2,
            "name": "角色管理",
            "type": 2,
            "sort": 0,
            "path": "/role",
            "component": "role",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 5,
            "parentId": 2,
            "name": "权限管理",
            "type": 2,
            "sort": 0,
            "path": "/permission",
            "component": "permission",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          },
          {
            "id": 6,
            "parentId": 2,
            "name": "用户管理",
            "type": 2,
            "sort": 0,
            "path": "/admin",
            "component": "admin",
            "icon": "icon-system",
            "redirect": "",
            "status": "1",
            "children": null
          }
        ]
      }
    ]
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 菜单详情

GET /api/admin/menu/detail

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|query|integer| yes ||菜单id|
|Authorization|header|string| yes ||none|

> Response Examples

> 成功

```json
{
  "code": 10200,
  "message": "成功",
  "data": {
    "menu": {
      "id": 1,
      "createTime": "2023-12-29T10:23:46+08:00",
      "updateTime": "2023-12-29T10:27:06+08:00",
      "deleteTime": null,
      "parentId": 0,
      "name": "主页",
      "type": 2,
      "sort": 0,
      "path": "/dashboard",
      "component": "Dashboard",
      "icon": "icon-home",
      "redirect": "nulla aliquip tempor commodo",
      "status": "1"
    }
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 管理员绑定角色

POST /api/admin/admin/bindRole

> Body Parameters

```json
{
  "adminId": 0,
  "roleIds": [
    0
  ]
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|Authorization|header|string| yes ||none|
|body|body|object| no ||none|
|» adminId|body|integer| yes | 管理员id|none|
|» roleIds|body|[integer]| no | 角色ids|none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

# Data Schema