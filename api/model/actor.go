package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const actor_tablename string = "mac_actor"

type Actor_Field struct {
	ActorID        *int     `json:"actor_id,omitempty" validate:"min=1"`
	TypeID         *int     `json:"type_id,omitempty"`
	TypeID1        *int     `json:"type_id_1,omitempty"`
	ActorName      *string  `json:"actor_name,omitempty"`
	ActorEn        *string  `json:"actor_en,omitempty"`
	ActorAlias     *string  `json:"actor_alias,omitempty"`
	ActorStatus    *int     `json:"actor_status,omitempty"`
	ActorLock      *int     `json:"actor_lock,omitempty"`
	ActorLetter    *string  `json:"actor_letter,omitempty"`
	ActorSex       *string  `json:"actor_sex,omitempty"`
	ActorColor     *string  `json:"actor_color,omitempty"`
	ActorPic       *string  `json:"actor_pic,omitempty"`
	ActorBlurb     *string  `json:"actor_blurb,omitempty"`
	ActorRemarks   *string  `json:"actor_remarks,omitempty"`
	ActorArea      *string  `json:"actor_area,omitempty"`
	ActorHeight    *string  `json:"actor_height,omitempty"`
	ActorWeight    *string  `json:"actor_weight,omitempty"`
	ActorBirthday  *string  `json:"actor_birthday,omitempty"`
	ActorBirtharea *string  `json:"actor_birtharea,omitempty"`
	ActorBlood     *string  `json:"actor_blood,omitempty"`
	ActorStarsign  *string  `json:"actor_starsign,omitempty"`
	ActorSchool    *string  `json:"actor_school,omitempty"`
	ActorWorks     *string  `json:"actor_works,omitempty"`
	ActorTag       *string  `json:"actor_tag,omitempty"`
	ActorClass     *string  `json:"actor_class,omitempty"`
	ActorLevel     *int     `json:"actor_level,omitempty"`
	ActorTime      *int     `json:"actor_time,omitempty"`
	ActorTimeAdd   *int     `json:"actor_time_add,omitempty"`
	ActorTimeHits  *int     `json:"actor_time_hits,omitempty"`
	ActorTimeMake  *int     `json:"actor_time_make,omitempty"`
	ActorHits      *int     `json:"actor_hits,omitempty"`
	ActorHitsDay   *int     `json:"actor_hits_day,omitempty"`
	ActorHitsWeek  *int     `json:"actor_hits_week,omitempty"`
	ActorHitsMonth *int     `json:"actor_hits_month,omitempty"`
	ActorScore     *float64 `json:"actor_score,omitempty"`
	ActorScoreAll  *int     `json:"actor_score_all,omitempty"`
	ActorScoreNum  *int     `json:"actor_score_num,omitempty"`
	ActorUp        *int     `json:"actor_up,omitempty"`
	ActorDown      *int     `json:"actor_down,omitempty"`
	ActorTpl       *string  `json:"actor_tpl,omitempty"`
	ActorJumpurl   *string  `json:"actor_jumpurl,omitempty"`
	ActorContent   *string  `json:"actor_content,omitempty"`
}

func Actor_Select(_data string) *common.Response_format_res {

	type Actor_Select struct {
		*Actor_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Actor_Select{}
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
	sSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", actor_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
func Actor_Update(_data string) *common.Response_format_res {
	type Actor_Update struct {
		*Actor_Field
		ActorID int `json:"actor_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Actor_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", actor_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Actor_Create(_data string) *common.Response_format_res {
	type Actor_Create struct {
		*Actor_Field
		ActorID *int `json:"actor_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Actor_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", actor_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Actor_Delete(_data string) *common.Response_format_res {

	type Actor_Delete struct {
		*Actor_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Actor_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", actor_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
