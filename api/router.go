package api

import (
	"fmt"
	"leo_go_api/api/common"

	"leo_go_api/api/middleware"

	"github.com/gofiber/fiber/v2"
)

/* 2022/03/04 leo 定義各動作可使用方法 */
var Action_mether_list = map[string][]string{
	"/User":    {"GET", "POST", "PUT", "DELECT"},
	"/Order":   {"GET", "POST", "PUT", "DELECT"},
	"/Group":   {"GET", "POST", "PUT", "DELECT"},
	"/Plog":    {"GET", "POST", "PUT", "DELECT"},
	"/Actor":   {"GET", "POST", "PUT", "DELECT"},
	"/Admin":   {"GET", "POST", "PUT", "DELECT"},
	"/Annex":   {"GET", "POST", "PUT", "DELECT"},
	"/Art":     {"GET", "POST", "PUT", "DELECT"},
	"/Card":    {"GET", "POST", "PUT", "DELECT"},
	"/Cash":    {"GET", "POST", "PUT", "DELECT"},
	"/Collect": {"GET", "POST", "PUT", "DELECT"},
	"/Topic":   {"GET", "POST", "PUT", "DELECT"},
	"/Type":    {"GET", "POST", "PUT", "DELECT"},
	"/Ulog":    {"GET", "POST", "PUT", "DELECT"},
	"/Vod":     {"GET", "POST", "PUT", "DELECT"},
}

func InitRouter(app *fiber.App) {
	common.InitDB()
	common.Initlogger()
	/* application */
	app.All("*", middleware.Intercept)
	for api_path := range common.Api_chennel_list {
		api_chennel := common.Api_chennel_list[api_path]
		api_chennel.Fiber_group = app.Group(fmt.Sprintf("/%s", api_path), middleware.Open_DB_Connect)
		for api_action, api_methed_list := range Action_mether_list {
			for _, api_methed := range api_methed_list {
				switch api_methed {
				case "GET":
					api_chennel.Fiber_group.Get(api_action, middleware.Transition_DB, middleware.Close_DB_Connect)
				case "POST":
					api_chennel.Fiber_group.Post(api_action, middleware.Transition_DB, middleware.Close_DB_Connect)
				case "PUT":
					api_chennel.Fiber_group.Put(api_action, middleware.Transition_DB, middleware.Close_DB_Connect)
				case "DELECT":
					api_chennel.Fiber_group.Delete(api_action, middleware.Transition_DB, middleware.Close_DB_Connect)
				}
			}
		}
	}
	// Test_vue_baseapi := app.Group("/BaseVue")
	// Test_vue_api := app.Group("/Test_vue", middleware.Open_DB_Connect_Vue)
	// Test_vue_baseapi.Get("/Test", static.Get_auth)
	// Test_vue_baseapi.Get("/GetLevel", static.Get_level)
	// Test_vue_baseapi.Get("/GetMaccms", static.Get_maccms)
	// Test_vue_baseapi.Post("/Set_Config", static.Set_config)
	// Test_vue_baseapi.Post("/Set_ConfigSeo", static.Set_config_seo)
	// Test_vue_baseapi.Post("/UploadImg", static.Test_img)
	// Test_vue_api.Get("/GetVodType", method.Vod_Type_List)
	// Test_vue_api.Get("/GetTypeList", method.TypeList)
	// Test_vue_api.Post("/Login", method.Login)
	// Test_vue_api.Post("/VodList", method.VodList)
}
