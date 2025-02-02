# Permission API
關於如何在 YAML 文件中設定權限並使用 Viper 讀取。  
自動檢查權限並更新資料庫。  
每次啟動專案時，建立一個系統使用者（使用者名稱: admin，密碼: 存在 .env 中），並賦予其所有權限。  
使用 logFile 記錄所有日誌。

## Overview

- 語言: Go v1.21.1
- Web 框架: Gin v1.9.1
- 資料庫: MongoDB v7.0.2

## Run

### 更新模組
```
go get -u && go mod tidy -v
```

### 執行
```
go run main.go
```

## Permission
複製 `etc/apiPermission.yaml.default` 並重新命名為 `etc/apiPermission.yaml` 以限制 API 存取權限。  
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

## ENV
複製 `.env.default` 並重新命名為 `.env`
```
MONGO_URL=
DB_NAME=
API_PORT=
ADMIN_PASSWORD=
JWTKey=
```

## API

### auth 身分認證（基本權限設定）
- POST /auth/login: 登入（不需要帶 token）
- POST /auth/logout: 登出

### user 使用者（基本權限設定）
- PATCH /user/myPassword: 修改自己的密碼
- PATCH /user/{userId}/password: 修改別人的密碼（需有權限）
- POST /user: 建立使用者（需有權限）
- POST /user/find: 搜尋使用者

### mapUserPermission 使用者與權限關聯（基本權限設定）
- POST /mapUserPermission: 建立使用者與權限關聯（需有權限）
- POST /mapUserPermission/find: 搜尋使用者與權限關聯（需有權限）
- DELETE /mapUserPermission/{id}: 刪除使用者與權限關聯（需有權限）

### permission 權限（基本權限設定）
- POST /permission/find: 搜尋所有權限（需有權限）

### setting 設定（權限相關應用）
- GET /setting/{code}: 搜尋設定
- PATCH /setting/{code}/{value}: 更新設定值（需有權限）

### task 任務（權限相關應用）
- POST /task: 分派任務（需有權限）
- POST /task/find: 搜尋所有任務
- PATCH /task/{id}/progressType/{progressType}: 更新任務進度（需為原先指派者或被指派者/DONE 前須驗收完畢/DELETE 需為原先指派者），移至測試會發送通知原先指派者驗收
- PATCH /task/{id}/checked/{checked}: 驗收/驗收失敗任務（需有權限/需為原先指派者），驗收成功會發送通知給被指派者
- DELETE /task/{id}: 刪除任務（需有權限/需為原先指派者）

### notification 通知（權限相關應用）
- GET /notification: 搜尋所有通知（時間從近排到遠）
- PATCH /notification/{id}/read: 已讀通知（需為被發送者）
- PATCH /notification/read/all: 已讀自己所有通知

## Customized Error Code
- InternalServerError: 000001
- InvalidTokenError: 000002
- WrongPasswordError: 000003
- UserNotFoundError: 000004
- PermissionDeniedError: 000107
- InvalidParameterError: 000108
