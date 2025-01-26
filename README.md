# Permission API


**(also provided Traditional Chinese version document [README-CH.md](README-CH.md).)**


About how to set permissions in a YAML file and read them using Viper.  
Automatically review permissions and update the database.  
And create a system user (username: admin, password: in .env) with all permissions each time the project is launched.  
With logFile to record all logs.

## Overview

- Language: Go v1.21.1
- Web Framework: Gin v1.9.1
- Database: MongoDB v7.0.2

## Run

### Update Modules
```
go get -u && go mod tidy -v
```


### Run
```
go run main.go
```


## Permission
Copy `etc/apiPermission.yaml.default` and rename it to `etc/apiPermission.yaml` to restrict API access rights.  
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
Copy `.env.default` and rename it to `.env`
```
MONGO_URL=
DB_NAME=
API_PORT=
ADMIN_PASSWORD=
JWTKey=
```

## API

### auth
- POST /auth/login: Login (no token required)
- POST /auth/logout: Logout

### user
- PATCH /user/myPassword: Change own password
- PATCH /user/{userId}/password: Change another user's password (requires permission)
- POST /user: Create a user (requires permission)
- POST /user/find: Search for users

### mapUserPermission
- POST /mapUserPermission: Create user-permission mapping (requires permission)
- POST /mapUserPermission/find: Search user-permission mapping (requires permission)
- DELETE /mapUserPermission/{id}: Delete user-permission mapping (requires permission)

### permission
- POST /permission/find: Search all permissions (requires permission)

### setting
- GET /setting/{code}: Get setting by code
- PATCH /setting/{code}/{value}: Update setting value (requires permission)

### task
- POST /task: Assign a task (requires permission)
- POST /task/find: Search all tasks
- PATCH /task/{id}/progressType/{progressType}: Update task progress (requires original assigner or assignee/DONE requires acceptance before completion/DELETE requires original assigner); notifications are sent to the original assigner for testing acceptance
- PATCH /task/{id}/checked/{checked}: Accept or reject task completion (requires permission/original assigner); successful acceptance sends notifications to the assignee
- DELETE /task/{id}: Delete a task (requires permission/original assigner)

### notification
- GET /notification: Retrieve all notifications (sorted from recent to oldest)
- PATCH /notification/{id}/read: Mark a notification as read (must be the recipient)
- PATCH /notification/read/all: Mark all personal notifications as read

## Customized Error Code
- InternalServerError: 000001
- InvalidTokenError: 000002
- WrongPasswordError: 000003
- UserNotFoundError: 000004
- PermissionDeniedError: 000107
- InvalidParameterError: 000108
