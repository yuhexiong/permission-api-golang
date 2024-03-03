# Permission API
About how to set permissions in a YAML file and read them using Viper.  
Automatically Review permissions and update database.  
And create a system user (username: admin, password: in .env) with all permissions each time the project is launched.

## Overview

- Language: Go v1.21.1
- Web FrameWork: Gin v1.9.1
- MongoDB v7.0.2


### 更新 Modules
```
go get -u && go mod tidy -v
```


### 執行程式
```
go run main.go
```

## Permission
copy etc/apiPermission.yaml.default and rename as etc/apiPermission.yaml to restrict api access rights.   
```
PermissionDefs:
  ChangePassword:
    Category: "PASSWORD"
    Code: "changePassword"

APIToPermission:
  - Url: "/user/:userId/password"
    Methods: "PATCH"
    PermissionName: "ChangePassword"
```


### ENV
copy .env.default and rename as .env
```
MONGO_URL=
DB_NAME=
API_PORT=
ADMIN_PASSWORD=
JWTKey=
```

## API

### auth 權限
- POST /auth/login: 登入（不需要帶token）
- POST /auth/logout: 登出

### user 使用者
- PATCH /user/myPassword: 修改自己的密碼
- PATCH /user/{userId}/password: 修改別人的密碼（需有權限）
- POST /user: 建立使用者（需有權限）
- POST /user/find: 搜尋使用者

### mapUserPermission 使用者與權限關聯
- POST /mapUserPermission: 建立使用者與權限關聯（需有權限）
- POST /mapUserPermission/find: 搜尋使用者與權限關聯（需有權限）
- DELETE /mapUserPermission: 刪除使用者與權限關聯（需有權限）

### permission 權限
- POST /permission/find: 搜尋所有權限（需有權限）

### task 任務
- POST /task: 分派任務（需有權限）
- POST /task/find: 搜尋所有任務
- PATCH /task/{id}/{checked}: 驗收/驗收失敗任務（需有權限/需視原先指派者）
- DELETE /task/{id}: 刪除任務（需有權限/需視原先指派者）

## Customized Error Code
- InternalServerError: 000001
- InvalidTokenError: 000002
- WrongPasswordError: 000003
- UserNotFoundError: 000004
- PermissionDeniedError: 000107
- InvalidParameterError: 000108