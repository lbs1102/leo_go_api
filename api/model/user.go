package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const user_tablename string = "mac_user"

type User_Field struct {
	UserID               *int    `json:"user_id,omitempty" validate:"min=1"`
	GroupID              *int    `json:"group_id,omitempty" validate:"min=1"`
	UserAnswer           *string `json:"user_answer,omitempty"`
	UserEmail            *string `json:"user_email,omitempty"`
	UserEndTime          *int    `json:"user_end_time,omitempty"`
	UserPresentationTime *int    `json:"user_presentation_time,omitempty"`
	UserFreePointTime    *int    `json:"user_free_point_time,omitempty" validate:"min=0"`
	UserExtend           *int    `json:"user_extend,omitempty"`
	UserLastLoginIP      *int    `json:"user_last_login_ip,omitempty"`
	UserLastLoginTime    *int    `json:"user_last_login_time,omitempty"`
	UserLoginIP          *int    `json:"user_login_ip,omitempty"`
	UserLoginNum         *int    `json:"user_login_num,omitempty"`
	UserLoginTime        *int    `json:"user_login_time,omitempty"`
	UserName             *string `json:"user_name,omitempty"`
	UserNickName         *string `json:"user_nick_name,omitempty"`
	UserOpenidQq         *string `json:"user_openid_qq,omitempty"`
	UserOpenidWeixin     *string `json:"user_openid_weixin,omitempty"`
	UserPhone            *string `json:"user_phone,omitempty"`
	UserPid              *int    `json:"user_pid,omitempty"`
	UserPid2             *int    `json:"user_pid_2,omitempty"`
	UserPid3             *int    `json:"user_pid_3,omitempty"`
	UserPoints           *int    `json:"user_points,omitempty" validate:"min=0"`
	UserFreeWatch        *int    `json:"user_free_watch,omitempty" validate:"min=0"`
	UserPointsFroze      *int    `json:"user_points_froze,omitempty" validate:"min=0"`
	UserPortrait         *string `json:"user_portrait,omitempty"`
	UserPortraitThumb    *string `json:"user_portrait_thumb,omitempty"`
	UserPwd              *string `json:"user_pwd,omitempty"`
	UserQq               *string `json:"user_qq,omitempty"`
	UserQuestion         *string `json:"user_question,omitempty"`
	UserRandom           *string `json:"user_random,omitempty"`
	UserRegIP            *int    `json:"user_reg_ip,omitempty"`
	UserRegTime          *int    `json:"user_reg_time,omitempty"`
	UserStatus           *int    `json:"user_status,omitempty"`
}

func User_Select(_data string) *common.Response_format_res {

	type User_Select struct {
		*User_Field
		UserID *int `json:"user_id,omitempty" validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &User_Select{}
	//step.1 針對資料解析與欄位進行驗證
	if err := json.Unmarshal(req_data_byte, req_data); err != nil {
		err_msg := ""
		if strings.Contains(err.Error(), "number") || strings.Contains(err.Error(), "string") {
			err_msg = "資料型態"
		}
		if strings.Contains(err.Error(), "invalid character") {
			err_msg = "資料解析"
		}
		log_data := fmt.Sprintf("%s錯誤 \n資料內容 : %s\n", err_msg, _data)
		res := common.Response_Format(2, fmt.Sprintf("%s錯誤 : %s", err_msg, _data), log_data)
		return res
	}
	if msg, err := common.Check_json_data_validator(*req_data); err {
		json_data, _ := json.Marshal(req_data)
		log_data := fmt.Sprintf("欄位驗證錯誤信息 : %s\n資料內容 : %s\n", msg, json_data)
		res := common.Response_Format(3, msg, log_data)
		return res
	}
	//step.2 取得sql字段陣列並執行sql
	Sql_arr := common.Get_sql_str_arr(req_data)
	sSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", user_tablename, strings.Join(Sql_arr.WHERE, " AND "))
	var res_data []map[string]interface{}
	db_row := common.Db.Raw(sSQL).Scan(&res_data)
	if db_row.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_row.Error, sSQL)
		res := common.Response_Format(4, "資料撈取失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料撈取成功", res_data)
	return res
}
func User_Update(_data string) *common.Response_format_res {
	type User_Update struct {
		*User_Field
		UserID int `json:"user_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &User_Update{}
	//step.1 針對資料解析與欄位進行驗證
	if err := json.Unmarshal(req_data_byte, req_data); err != nil {
		err_msg := ""
		if strings.Contains(err.Error(), "number") || strings.Contains(err.Error(), "string") {
			err_msg = "資料型態"
		}
		if strings.Contains(err.Error(), "invalid character") {
			err_msg = "資料解析"
		}
		log_data := fmt.Sprintf("%s錯誤 \n資料內容 : %s\n", err_msg, _data)
		res := common.Response_Format(2, fmt.Sprintf("%s錯誤 : %s", err_msg, _data), log_data)
		return res
	}
	if msg, err := common.Check_json_data_validator(*req_data); err {
		json_data, _ := json.Marshal(req_data)
		log_data := fmt.Sprintf("欄位驗證錯誤信息 : %s\n資料內容 : %s\n", msg, json_data)
		res := common.Response_Format(3, msg, log_data)
		return res
	}
	//step.2 取得sql字段陣列並執行sql
	Sql_arr := common.Get_sql_str_arr(req_data)
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", user_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料更新失敗", log_data)
		return res
	}
	if db_res.RowsAffected == 0 {
		res := common.Response_Format(0, "資料無異動", nil)
		return res
	}
	res := common.Response_Format(1, "資料更新成功", nil)
	return res
}
func User_Create(_data string) *common.Response_format_res {
	type User_Create struct {
		*User_Field
		UserID *int `json:"user_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &User_Create{}
	//step.1 針對資料解析與欄位進行驗證
	if err := json.Unmarshal(req_data_byte, req_data); err != nil {
		err_msg := ""
		if strings.Contains(err.Error(), "number") || strings.Contains(err.Error(), "string") {
			err_msg = "資料型態"
		}
		if strings.Contains(err.Error(), "invalid character") {
			err_msg = "資料解析"
		}
		log_data := fmt.Sprintf("%s錯誤 \n資料內容 : %s\n", err_msg, _data)
		res := common.Response_Format(2, fmt.Sprintf("%s錯誤 : %s", err_msg, _data), log_data)
		return res
	}
	if msg, err := common.Check_json_data_validator(*req_data); err {
		json_data, _ := json.Marshal(req_data)
		log_data := fmt.Sprintf("欄位驗證錯誤信息 : %s\n資料內容 : %s\n", msg, json_data)
		res := common.Response_Format(3, msg, log_data)
		return res
	}
	//step.2 取得sql字段陣列並執行sql
	Sql_arr := common.Get_sql_str_arr(req_data)
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", user_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func User_Delete(_data string) *common.Response_format_res {

	type User_Delete struct {
		*User_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &User_Delete{}
	//step.1 針對資料解析與欄位進行驗證
	if err := json.Unmarshal(req_data_byte, req_data); err != nil {
		err_msg := ""
		if strings.Contains(err.Error(), "number") || strings.Contains(err.Error(), "string") {
			err_msg = "資料型態"
		}
		if strings.Contains(err.Error(), "invalid character") {
			err_msg = "資料解析"
		}
		log_data := fmt.Sprintf("%s錯誤 \n資料內容 : %s\n", err_msg, _data)
		res := common.Response_Format(2, fmt.Sprintf("%s錯誤 : %s", err_msg, _data), log_data)
		return res
	}
	if msg, err := common.Check_json_data_validator(*req_data); err {
		json_data, _ := json.Marshal(req_data)
		log_data := fmt.Sprintf("欄位驗證錯誤信息 : %s\n資料內容 : %s\n", msg, json_data)
		res := common.Response_Format(3, msg, log_data)
		return res
	}
	//step.2 取得sql字段陣列並執行sql
	Sql_arr := common.Get_sql_str_arr(req_data)
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", user_tablename, strings.Join(Sql_arr.WHERE, " AND "))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料刪除失敗", log_data)
		return res
	}
	if db_res.RowsAffected == 0 {
		res := common.Response_Format(0, "資料無異動", nil)
		return res
	}
	res := common.Response_Format(1, "資料刪除成功", nil)
	return res
}
