package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const order_tablename string = "mac_order"

type Order_Field struct {
	OrderID      *int     `json:"order_id,omitempty" validate:"min=1"`
	UserID       *int     `json:"user_id,omitempty" validate:"min=1"`
	OrderStatus  *int     `json:"order_status,omitempty"`
	OrderCode    *string  `json:"order_code,omitempty"`
	OrderPrice   *float64 `json:"order_price,omitempty"`
	OrderTime    *int     `json:"order_time,omitempty"`
	OrderPoints  *int     `json:"order_points,omitempty"`
	OrderPayType *string  `json:"order_pay_type,omitempty"`
	OrderPayTime *int     `json:"order_pay_time,omitempty"`
	OrderPayID   *string  `json:"order_pay_id,omitempty"`
	OrderRemarks *string  `json:"order_remarks,omitempty"`
}

func Order_Select(_data string) *common.Response_format_res {

	type Order_Select struct {
		*Order_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Order_Select{}
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
		log_data := fmt.Sprintf("資料驗證錯誤 : %s\n資料內容 : %s\n", msg, json_data)
		res := common.Response_Format(3, msg, log_data)
		return res
	}
	//step.2 取得sql字段陣列並執行sql
	Sql_arr := common.Get_sql_str_arr(req_data)
	sSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", order_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
func Order_Update(_data string) *common.Response_format_res {
	type Order_Update struct {
		*Order_Field
		OrderID *int `json:"order_id" validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Order_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", order_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Order_Create(_data string) *common.Response_format_res {
	type Order_Create struct {
		*Order_Field
		OrderID *int `json:"order_id" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Order_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", order_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Order_Delete(_data string) *common.Response_format_res {

	type Order_Delete struct {
		*Order_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Order_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", order_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
