package middleware

import (
	"fmt"
	"leo_go_api/api/common"
	"leo_go_api/api/model"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/* 2022/03/25 leo 定義個動作方法所對應的func */
var Api_model_list = map[string]map[string]func(_data string) *common.Response_format_res{
	"User":    {"GET": model.User_Select, "POST": model.User_Create, "PUT": model.User_Update, "DELECT": model.User_Delete},
	"Order":   {"GET": model.Order_Select, "POST": model.Order_Create, "PUT": model.Order_Update, "DELECT": model.Order_Delete},
	"Group":   {"GET": model.Group_Select, "POST": model.Group_Create, "PUT": model.Group_Update, "DELECT": model.Group_Delete},
	"Plog":    {"GET": model.Plog_Select, "POST": model.Plog_Create, "PUT": model.Plog_Update, "DELECT": model.Plog_Delete},
	"Actor":   {"GET": model.Actor_Select, "POST": model.Actor_Create, "PUT": model.Actor_Update, "DELECT": model.Actor_Delete},
	"Admin":   {"GET": model.Admin_Select, "POST": model.Admin_Create, "PUT": model.Admin_Update, "DELECT": model.Admin_Delete},
	"Annex":   {"GET": model.Annex_Select, "POST": model.Annex_Create, "PUT": model.Annex_Update, "DELECT": model.Annex_Delete},
	"Art":     {"GET": model.Art_Select, "POST": model.Art_Create, "PUT": model.Art_Update, "DELECT": model.Art_Delete},
	"Card":    {"GET": model.Card_Select, "POST": model.Card_Create, "PUT": model.Card_Update, "DELECT": model.Card_Delete},
	"Cash":    {"GET": model.Cash_Select, "POST": model.Cash_Create, "PUT": model.Cash_Update, "DELECT": model.Cash_Delete},
	"Collect": {"GET": model.Collect_Select, "POST": model.Collect_Create, "PUT": model.Collect_Update, "DELECT": model.Collect_Delete},
	"Topic":   {"GET": model.Topic_Select, "POST": model.Topic_Create, "PUT": model.Topic_Update, "DELECT": model.Topic_Delete},
	"Type":    {"GET": model.Type_Select, "POST": model.Type_Create, "PUT": model.Type_Update, "DELECT": model.Type_Delete},
	"Ulog":    {"GET": model.Ulog_Select, "POST": model.Ulog_Create, "PUT": model.Ulog_Update, "DELECT": model.Ulog_Delete},
	"Vod":     {"GET": model.Vod_Select, "POST": model.Vod_Create, "PUT": model.Vod_Update, "DELECT": model.Vod_Delete},
}

func Intercept(c *fiber.Ctx) error {
	/* allow only requests from localhost */
	var RequestIP = c.IP()
	var LocalHost = common.Allow_HostName
	var RequestMethod = c.Method()
	var RequestUri = c.OriginalURL()
	if RequestIP == LocalHost {
		common.LogInfo(fmt.Sprintf("ip: %s connection success, RequestMethod: %s, RequestUri: %s", RequestIP, RequestMethod, RequestUri))
		return c.Next()
	} else {
		common.LogWarn(fmt.Sprintf("ip: %s connection denied, RequestMethod: %s, RequestUri: %s", RequestIP, RequestMethod, RequestUri))
		return c.SendStatus(http.StatusUnauthorized)
	}
}

func Open_DB_Connect(c *fiber.Ctx) error {
	/* allow only requests from localhost */
	var RequestIP = c.IP()
	var RequestMethod = c.Method()
	var RequestUri = c.OriginalURL()
	request_path := strings.Split(strings.TrimLeft(c.OriginalURL(), "/"), "/")[0]
	common.Select_db_for_api(request_path)
	if c.Body() == nil {
		common.LogWarn(fmt.Sprintf("ip: %s Request empty, RequestMethod: %s, RequestUri: %s", RequestIP, RequestMethod, RequestUri))
		return c.SendStatus(http.StatusExpectationFailed)
	}
	return c.Next()
}
func Open_DB_Connect_Vue(c *fiber.Ctx) error {
	/* allow only requests from localhost */
	var RequestIP = c.IP()
	var RequestMethod = c.Method()
	var RequestUri = c.OriginalURL()
	common.Select_db_for_api("develop")
	if c.Body() == nil && RequestMethod != "GET" {
		common.LogWarn(fmt.Sprintf("ip: %s Request empty, RequestMethod: %s, RequestUri: %s", RequestIP, RequestMethod, RequestUri))
		return c.SendStatus(http.StatusExpectationFailed)
	}
	return c.Next()
}
func Transition_DB(c *fiber.Ctx) error {
	request_method := c.Method()
	request_path := strings.Split(strings.TrimLeft(c.Path(), "/"), "/")[1]
	request_path_2 := strings.Split(strings.TrimLeft(c.OriginalURL(), "/"), "/")[0]

	decode_data := common.Aes_decode(string(c.Body()))
	if request_path_2 == "api_c2" && request_path == "Vod" {
		common.LogWarn(decode_data)
	}
	res := Api_model_list[request_path][request_method](decode_data)
	c.Locals("res", res)
	return c.Next()
}
func Close_DB_Connect(c *fiber.Ctx) error {
	res := c.Locals("res")
	return c.JSON(res)
}
