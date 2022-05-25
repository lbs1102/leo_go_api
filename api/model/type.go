package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const type_tablename string = "mac_type"

type Type_Field struct {
	TypeID        *int    `json:"type_id,omitempty" validate:"min=1"`
	TypeName      *string `json:"type_name,omitempty"`
	TypeEn        *string `json:"type_en,omitempty"`
	TypeSort      *int    `json:"type_sort,omitempty"`
	TypeMid       *int    `json:"type_mid,omitempty"`
	TypePid       *int    `json:"type_pid,omitempty"`
	TypeStatus    *int    `json:"type_status,omitempty"`
	TypeTpl       *string `json:"type_tpl,omitempty"`
	TypeTplList   *string `json:"type_tpl_list,omitempty"`
	TypeTplDetail *string `json:"type_tpl_detail,omitempty"`
	TypeTplPlay   *string `json:"type_tpl_play,omitempty"`
	TypeTplDown   *string `json:"type_tpl_down,omitempty"`
	TypeKey       *string `json:"type_key,omitempty"`
	TypeDes       *string `json:"type_des,omitempty"`
	TypeTitle     *string `json:"type_title,omitempty"`
	TypeUnion     *string `json:"type_union,omitempty"`
	TypeExtend    *string `json:"type_extend,omitempty"`
	TypeLogo      *string `json:"type_logo,omitempty"`
	TypePic       *string `json:"type_pic,omitempty"`
	TypeJumpurl   *string `json:"type_jumpurl,omitempty"`
}

func Type_Select(_data string) *common.Response_format_res {

	type Type_Select struct {
		*Type_Field
		Sql_page  int    `json:"limit_page,omitempty"`
		Sql_count int    `json:"limit_count,omitempty"`
		Sql_order string `json:"order_field,omitempty"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Type_Select{}
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
	sSQL := ""
	if Sql_arr.Index == 0 {
		sSQL = fmt.Sprintf("SELECT * FROM %s ", type_tablename)
	} else {
		sSQL = fmt.Sprintf("SELECT * FROM %s WHERE %s", type_tablename, strings.Join(Sql_arr.WHERE, " AND "))
	}
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
func Type_Update(_data string) *common.Response_format_res {
	type Type_Update struct {
		*Type_Field
		TypeID int `json:"type_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Type_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", type_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Type_Create(_data string) *common.Response_format_res {
	type Type_Create struct {
		*Type_Field
		TypeID *int `json:"type_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Type_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", type_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Type_Delete(_data string) *common.Response_format_res {

	type Type_Delete struct {
		*Type_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Type_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", type_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
