package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const art_tablename string = "mac_art"

type Art_Field struct {
	ArtID            *int     `json:"art_id,omitempty" validate:"min=1"`
	TypeID           *int     `json:"type_id,omitempty"`
	TypeID1          *int     `json:"type_id_1,omitempty"`
	GroupID          *int     `json:"group_id,omitempty"`
	ArtName          *string  `json:"art_name,omitempty"`
	ArtSub           *string  `json:"art_sub,omitempty"`
	ArtEn            *string  `json:"art_en,omitempty"`
	ArtStatus        *int     `json:"art_status,omitempty"`
	ArtLetter        *string  `json:"art_letter,omitempty"`
	ArtColor         *string  `json:"art_color,omitempty"`
	ArtFrom          *string  `json:"art_from,omitempty"`
	ArtAuthor        *string  `json:"art_author,omitempty"`
	ArtTag           *string  `json:"art_tag,omitempty"`
	ArtClass         *string  `json:"art_class,omitempty"`
	ArtPic           *string  `json:"art_pic,omitempty"`
	ArtPicThumb      *string  `json:"art_pic_thumb,omitempty"`
	ArtPicSlide      *string  `json:"art_pic_slide,omitempty"`
	ArtBlurb         *string  `json:"art_blurb,omitempty"`
	ArtRemarks       *string  `json:"art_remarks,omitempty"`
	ArtJumpurl       *string  `json:"art_jumpurl,omitempty"`
	ArtTpl           *string  `json:"art_tpl,omitempty"`
	ArtLevel         *int     `json:"art_level,omitempty"`
	ArtLock          *int     `json:"art_lock,omitempty"`
	ArtPoints        *int     `json:"art_points,omitempty"`
	ArtPointsDetail  *int     `json:"art_points_detail,omitempty"`
	ArtUp            *int     `json:"art_up,omitempty"`
	ArtDown          *int     `json:"art_down,omitempty"`
	ArtHits          *int     `json:"art_hits,omitempty"`
	ArtHitsDay       *int     `json:"art_hits_day,omitempty"`
	ArtHitsWeek      *int     `json:"art_hits_week,omitempty"`
	ArtHitsMonth     *int     `json:"art_hits_month,omitempty"`
	ArtTime          *int     `json:"art_time,omitempty"`
	ArtTimeAdd       *int     `json:"art_time_add,omitempty"`
	ArtTimeHits      *int     `json:"art_time_hits,omitempty"`
	ArtTimeMake      *int     `json:"art_time_make,omitempty"`
	ArtScore         *float64 `json:"art_score,omitempty"`
	ArtScoreAll      *int     `json:"art_score_all,omitempty"`
	ArtScoreNum      *int     `json:"art_score_num,omitempty"`
	ArtRelArt        *string  `json:"art_rel_art,omitempty"`
	ArtRelVod        *string  `json:"art_rel_vod,omitempty"`
	ArtPwd           *string  `json:"art_pwd,omitempty"`
	ArtPwdURL        *string  `json:"art_pwd_url,omitempty"`
	ArtTitle         *string  `json:"art_title,omitempty"`
	ArtNote          *string  `json:"art_note,omitempty"`
	ArtContent       *string  `json:"art_content,omitempty"`
	ArtPicScreenshot *string  `json:"art_pic_screenshot,omitempty"`
}

func Art_Select(_data string) *common.Response_format_res {

	type Art_Select struct {
		*Art_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Art_Select{}
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
	sSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", art_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
func Art_Update(_data string) *common.Response_format_res {
	type Art_Update struct {
		*Art_Field
		ArtID int `json:"art_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Art_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", art_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Art_Create(_data string) *common.Response_format_res {
	type Art_Create struct {
		*Art_Field
		ArtID *int `json:"art_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Art_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", art_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Art_Delete(_data string) *common.Response_format_res {

	type Art_Delete struct {
		*Art_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Art_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", art_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
