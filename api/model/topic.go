package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const topic_tablename string = "mac_topic"

type Topic_Field struct {
	TopicID        *int     `json:"topic_id,omitempty" validate:"min=1"`
	TopicName      *string  `json:"topic_name,omitempty"`
	TopicEn        *string  `json:"topic_en,omitempty"`
	TopicSub       *string  `json:"topic_sub,omitempty"`
	TopicStatus    *int     `json:"topic_status,omitempty"`
	TopicSort      *int     `json:"topic_sort,omitempty"`
	TopicLetter    *string  `json:"topic_letter,omitempty"`
	TopicColor     *string  `json:"topic_color,omitempty"`
	TopicTpl       *string  `json:"topic_tpl,omitempty"`
	TopicType      *string  `json:"topic_type,omitempty"`
	TopicPic       *string  `json:"topic_pic,omitempty"`
	TopicPicThumb  *string  `json:"topic_pic_thumb,omitempty"`
	TopicPicSlide  *string  `json:"topic_pic_slide,omitempty"`
	TopicKey       *string  `json:"topic_key,omitempty"`
	TopicDes       *string  `json:"topic_des,omitempty"`
	TopicTitle     *string  `json:"topic_title,omitempty"`
	TopicBlurb     *string  `json:"topic_blurb,omitempty"`
	TopicRemarks   *string  `json:"topic_remarks,omitempty"`
	TopicLevel     *int     `json:"topic_level,omitempty"`
	TopicUp        *int     `json:"topic_up,omitempty"`
	TopicDown      *int     `json:"topic_down,omitempty"`
	TopicScore     *float64 `json:"topic_score,omitempty"`
	TopicScoreAll  *int     `json:"topic_score_all,omitempty"`
	TopicScoreNum  *int     `json:"topic_score_num,omitempty"`
	TopicHits      *int     `json:"topic_hits,omitempty"`
	TopicHitsDay   *int     `json:"topic_hits_day,omitempty"`
	TopicHitsWeek  *int     `json:"topic_hits_week,omitempty"`
	TopicHitsMonth *int     `json:"topic_hits_month,omitempty"`
	TopicTime      *int     `json:"topic_time,omitempty"`
	TopicTimeAdd   *int     `json:"topic_time_add,omitempty"`
	TopicTimeHits  *int     `json:"topic_time_hits,omitempty"`
	TopicTimeMake  *int     `json:"topic_time_make,omitempty"`
	TopicTag       *string  `json:"topic_tag,omitempty"`
	TopicRelVod    *string  `json:"topic_rel_vod,omitempty"`
	TopicRelArt    *string  `json:"topic_rel_art,omitempty"`
	TopicContent   *string  `json:"topic_content,omitempty"`
	TopicExtend    *string  `json:"topic_extend,omitempty"`
}

func Topic_Select(_data string) *common.Response_format_res {

	type Topic_Select struct {
		*Topic_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Topic_Select{}
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
	sSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", topic_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
func Topic_Update(_data string) *common.Response_format_res {
	type Topic_Update struct {
		*Topic_Field
		TopicID int `json:"topic_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Topic_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", topic_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Topic_Create(_data string) *common.Response_format_res {
	type Topic_Create struct {
		*Topic_Field
		TopicID *int `json:"topic_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Topic_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", topic_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Topic_Delete(_data string) *common.Response_format_res {

	type Topic_Delete struct {
		*Topic_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Topic_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", topic_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
