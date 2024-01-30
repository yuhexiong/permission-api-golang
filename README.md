# Permission API
About how to set permissions in a YAML file and read them using Viper, review permissions and update database.  
Automatically create a system user (username: admin, password: password) with all permissions each time the project is launched.

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


### ENV
copy .env.default and rename as .env
```
MONGO_URL=
DB_NAME=
API_PORT=
JWTKey=
```

## API

### auth 權限（不需要登入）
- POST /auth/login: 登入

### user 使用者
- POST /user/logout: 登出
- POST /user: 建立使用者
- POST /user/find: 搜尋使用者

### mapUserPermission 使用者與權限關聯
- POST /mapUserPermission: 建立使用者與權限關聯

### permission 權限
- POST /permission／find: 搜尋所有權限


## Customized Error Code
- InvalidTokenError: 000002
- UserNotFoundError: 000003
- PermissionDeniedError: 000107
- InvalidParameterError: 000108