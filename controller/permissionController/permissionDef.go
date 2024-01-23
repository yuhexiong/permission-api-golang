package permissionController

import (
	"fmt"
	"permission-api/model"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var permViper *viper.Viper
var PermissionsMap = make(map[string]PermissionDef)
var apiToPermissionMap = make(map[string]ApiPermissionInfo)

type PermissionDef struct {
	Category string `json:"category" binding:"required" example:"USER"`   // 權限種類
	Code     string `json:"code" binding:"required" example:"createUser"` // 權限代號
}

// 讀、寫兩種執行權限
type PermissionOp string

const (
	READ  PermissionOp = "R"
	WRITE PermissionOp = "W"
)

type ApiToPermission struct {
	Url            string       `mapstructure:"Url"`
	Methods        string       `mapstructure:"Methods"`
	PermissionName string       `mapstructure:"PermissionName"`
	PermissionOp   PermissionOp `mapstructure:"PermissionOp"`
}

type ApiPermissionInfo struct {
	PermissionName string
	PermissionOp   PermissionOp
}

func InitViper() {
	permViper = viper.New()
	permViper.SetConfigName("permissionDefs")
	permViper.SetConfigType("yaml")
	permViper.AddConfigPath("etc/")
	err := permViper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config failed: %v", err)
	}

	loadPermissionsMap()
	loadApiToPermission()

	permViper.OnConfigChange(func(e fsnotify.Event) {
		err := permViper.ReadInConfig()
		if err != nil {
			fmt.Printf("read config failed: %v", err)
		}
		loadPermissionsMap()
		loadApiToPermission()
	})

	permViper.WatchConfig()
}

func loadPermissionsMap() {
	// 取得yaml的權限並存入map, reset PermissionsMap as empty map
	PermissionsMap = make(map[string]PermissionDef)
	permissionsYAML := permViper.GetStringMap("PermissionDefs")
	for key, value := range permissionsYAML {
		data := value.(map[string]interface{})
		PermissionsMap[key] = PermissionDef{
			Category: data["category"].(string),
			Code:     data["code"].(string),
		}
	}

	// 取得db現有的權限
	foundPermissions := []*model.Permission{}
	if err := FindPermission(FindPermissionOptions{}, &foundPermissions); err != nil {
		panic(err)
	}

	// yaml無但db有此權限 則停用
	for _, foundPermission := range foundPermissions {
		if *foundPermission.Status != model.NormalStatus { // 未啟用的權限不驗證
			continue
		}
		if exists := checkPermissionExistInYAML(foundPermission, PermissionsMap); !exists {
			if err := DeletePermission(foundPermission.ID); err != nil {
				panic(err)
			}
		}
	}

	// yaml有但db無此權限 則新增
	for _, permissionMap := range PermissionsMap {
		exist, foundPermission := checkPermissionExistInDB(foundPermissions, permissionMap)
		if !exist {
			createOpts := CreatePermissionOptions(permissionMap)
			if err := CreatePermission(createOpts, nil); err != nil {
				panic(err)
			}

		} else if foundPermission != nil {
			if err := EnablePermission(foundPermission.ID); err != nil {
				panic(err)
			}
		}
	}

}

func checkPermissionExistInYAML(foundPermission *model.Permission, PermissionsMap map[string]PermissionDef) bool {
	for _, permissionMap := range PermissionsMap {
		f := *foundPermission
		if permissionMap.Category == f.Category && permissionMap.Code == f.Code {
			return true
		}
	}

	return false
}

func checkPermissionExistInDB(foundPermissions []*model.Permission, p PermissionDef) (bool, *model.Permission) {
	for _, foundPermission := range foundPermissions {
		f := *foundPermission
		if p.Category == f.Category && p.Code == f.Code {
			if *f.Status == model.NormalStatus {
				return true, nil
			} else {
				// 須重新啟用, 則回傳permission以更新
				return true, &f
			}
		}
	}

	return false, nil
}

func loadApiToPermission() {
	var apiToPermissionsYAML []ApiToPermission
	if err := permViper.UnmarshalKey("APIToPermission", &apiToPermissionsYAML); err != nil {
		panic(err)
	}

	for _, apiToPermissionYAML := range apiToPermissionsYAML {
		key := fmt.Sprintf("%s-%s", apiToPermissionYAML.Methods, apiToPermissionYAML.Url)

		var permissionOp PermissionOp
		if apiToPermissionYAML.PermissionOp == READ {
			permissionOp = READ
		} else if apiToPermissionYAML.PermissionOp == WRITE {
			permissionOp = WRITE
		}

		apiToPermissionMap[key] = ApiPermissionInfo{
			PermissionName: apiToPermissionYAML.PermissionName,
			PermissionOp:   permissionOp,
		}
	}
}

func GetApiPermission(fullPath string, method string) (*PermissionDef, *PermissionOp) {
	apiRoute := fmt.Sprintf("%s-%s", method, fullPath)

	for api, permissionInfo := range apiToPermissionMap {
		if apiRoute == api {
			permissionDef, ok := PermissionsMap[strings.ToLower(permissionInfo.PermissionName)]
			if ok {
				return &permissionDef, &permissionInfo.PermissionOp
			}
		}
	}
	return nil, nil
}
