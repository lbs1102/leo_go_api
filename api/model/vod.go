package model

import (
	"encoding/json"
	"fmt"
	"leo_go_api/api/common"
	"strings"
)

const vod_tablename string = "mac_vod"

type Vod_Field struct {
	VodID            *int     `json:"vod_id,omitempty" validate:"min=1"`
	TypeID           *int     `json:"type_id,omitempty"`
	TypeID1          *int     `json:"type_id_1,omitempty"`
	TypeMode         *int     `json:"type_mode,omitempty"`
	GroupID          *int     `json:"group_id,omitempty"`
	VodName          *string  `json:"vod_name,omitempty"`
	VodSub           *string  `json:"vod_sub,omitempty"`
	VodEn            *string  `json:"vod_en,omitempty"`
	VodVi            *string  `json:"vod_vi,omitempty"`
	VodStatus        *int     `json:"vod_status,omitempty"`
	VodLetter        *string  `json:"vod_letter,omitempty"`
	VodColor         *string  `json:"vod_color,omitempty"`
	VodTag           *string  `json:"vod_tag,omitempty"`
	VodClass         *string  `json:"vod_class,omitempty"`
	VodPic           *string  `json:"vod_pic,omitempty"`
	VodPicThumb      *string  `json:"vod_pic_thumb,omitempty"`
	VodPicSlide      *string  `json:"vod_pic_slide,omitempty"`
	VodActor         *string  `json:"vod_actor,omitempty"`
	VodDirector      *string  `json:"vod_director,omitempty"`
	VodWriter        *string  `json:"vod_writer,omitempty"`
	VodBehind        *string  `json:"vod_behind,omitempty"`
	VodBlurb         *string  `json:"vod_blurb,omitempty"`
	VodRemarks       *string  `json:"vod_remarks,omitempty"`
	VodPubdate       *string  `json:"vod_pubdate,omitempty"`
	VodTotal         *int     `json:"vod_total,omitempty"`
	VodSerial        *string  `json:"vod_serial,omitempty"`
	VodTv            *string  `json:"vod_tv,omitempty"`
	VodWeekday       *string  `json:"vod_weekday,omitempty"`
	VodArea          *string  `json:"vod_area,omitempty"`
	VodLang          *string  `json:"vod_lang,omitempty"`
	VodYear          *string  `json:"vod_year,omitempty"`
	VodVersion       *string  `json:"vod_version,omitempty"`
	VodState         *string  `json:"vod_state,omitempty"`
	VodAuthor        *string  `json:"vod_author,omitempty"`
	VodJumpurl       *string  `json:"vod_jumpurl,omitempty"`
	VodTpl           *string  `json:"vod_tpl,omitempty"`
	VodTplPlay       *string  `json:"vod_tpl_play,omitempty"`
	VodTplDown       *string  `json:"vod_tpl_down,omitempty"`
	VodIsend         *int     `json:"vod_isend,omitempty"`
	VodLock          *int     `json:"vod_lock,omitempty"`
	VodLevel         *int     `json:"vod_level,omitempty"`
	VodCopyright     *int     `json:"vod_copyright,omitempty"`
	VodPoints        *int     `json:"vod_points,omitempty"`
	VodPointsPlay    *int     `json:"vod_points_play,omitempty"`
	VodPointsDown    *int     `json:"vod_points_down,omitempty"`
	VodHits          *int     `json:"vod_hits,omitempty"`
	VodHitsDay       *int     `json:"vod_hits_day,omitempty"`
	VodHitsWeek      *int     `json:"vod_hits_week,omitempty"`
	VodHitsMonth     *int     `json:"vod_hits_month,omitempty"`
	VodDuration      *string  `json:"vod_duration,omitempty"`
	VodUp            *int     `json:"vod_up,omitempty"`
	VodDown          *int     `json:"vod_down,omitempty"`
	VodScore         *float64 `json:"vod_score,omitempty"`
	VodScoreAll      *int     `json:"vod_score_all,omitempty"`
	VodScoreNum      *int     `json:"vod_score_num,omitempty"`
	VodTime          *int     `json:"vod_time,omitempty"`
	VodTimeAdd       *int     `json:"vod_time_add,omitempty"`
	VodTimeHits      *int     `json:"vod_time_hits,omitempty"`
	VodTimeMake      *int     `json:"vod_time_make,omitempty"`
	VodTrysee        *int     `json:"vod_trysee,omitempty"`
	VodDoubanID      *int     `json:"vod_douban_id,omitempty"`
	VodDoubanScore   *float64 `json:"vod_douban_score,omitempty"`
	VodReurl         *string  `json:"vod_reurl,omitempty"`
	VodRelVod        *string  `json:"vod_rel_vod,omitempty"`
	VodRelArt        *string  `json:"vod_rel_art,omitempty"`
	VodPwd           *string  `json:"vod_pwd,omitempty"`
	VodPwdURL        *string  `json:"vod_pwd_url,omitempty"`
	VodPwdPlay       *string  `json:"vod_pwd_play,omitempty"`
	VodPwdPlayURL    *string  `json:"vod_pwd_play_url,omitempty"`
	VodPwdDown       *string  `json:"vod_pwd_down,omitempty"`
	VodPwdDownURL    *string  `json:"vod_pwd_down_url,omitempty"`
	VodContent       *string  `json:"vod_content,omitempty"`
	VodPlayFrom      *string  `json:"vod_play_from,omitempty"`
	VodPlayServer    *string  `json:"vod_play_server,omitempty"`
	VodPlayNote      *string  `json:"vod_play_note,omitempty"`
	VodPlayURL       *string  `json:"vod_play_url,omitempty"`
	VodDownFrom      *string  `json:"vod_down_from,omitempty"`
	VodDownServer    *string  `json:"vod_down_server,omitempty"`
	VodDownNote      *string  `json:"vod_down_note,omitempty"`
	VodDownURL       *string  `json:"vod_down_url,omitempty"`
	VodPlot          *int     `json:"vod_plot,omitempty"`
	VodPlotName      *string  `json:"vod_plot_name,omitempty"`
	VodPlotDetail    *string  `json:"vod_plot_detail,omitempty"`
	VodPicScreenshot *string  `json:"vod_pic_screenshot,omitempty"`
}

func Vod_Get_Count(_data string) *common.Response_format_res {
	type Vod_Select struct {
		*Vod_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Vod_Select{}
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
		sSQL = fmt.Sprintf("SELECT COUNT(1) FROM %s", vod_tablename)
	} else {
		sSQL = fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s ", vod_tablename, strings.Join(Sql_arr.WHERE, " AND "))
	}
	var res_data = 0
	db_row := common.Db.Raw(sSQL).Scan(&res_data)
	if db_row.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_row.Error, sSQL)
		res := common.Response_Format(4, "資料撈取失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料撈取成功", res_data)
	return res
}
func Vod_Select(_data string) *common.Response_format_res {

	type Vod_Select struct {
		*Vod_Field
		Sql_page  int    `json:"limit_page,omitempty"`
		Sql_count int    `json:"limit_count,omitempty"`
		Sql_order string `json:"order_field,omitempty"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Vod_Select{}
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
	Sql_limit := ""
	if req_data.Sql_page > 0 && req_data.Sql_count > 0 {
		Sql_limit = fmt.Sprintf(" Limit %v,%v ", ((req_data.Sql_page - 1) * req_data.Sql_count), req_data.Sql_count)
	}
	sSQL := ""
	if Sql_arr.Index == 0 {
		sSQL = fmt.Sprintf("SELECT * FROM %s %s", vod_tablename, Sql_limit)
	} else {
		sSQL = fmt.Sprintf("SELECT * FROM %s WHERE %s %s", vod_tablename, strings.Join(Sql_arr.WHERE, " AND "), Sql_limit)
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
func Vod_Update(_data string) *common.Response_format_res {
	type Vod_Update struct {
		*Vod_Field
		VodID int `json:"vod_id"  validate:"min=1"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Vod_Update{}
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
	sSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", vod_tablename, strings.Join(Sql_arr.UPDATE.Value, ","), strings.Join(Sql_arr.UPDATE.Filed, " AND "))
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
func Vod_Create(_data string) *common.Response_format_res {
	type Vod_Create struct {
		*Vod_Field
		VodID *int `json:"vod_id,omitempty" validate:"excluded_with_all"`
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Vod_Create{}
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
	sSQL := fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( %s ) ", vod_tablename, strings.Join(Sql_arr.CREATE.Filed, ","), strings.Join(Sql_arr.CREATE.Value, ","))
	db_res := common.Db.Exec(sSQL)
	if db_res.Error != nil {
		log_data := fmt.Sprintf("SQL錯誤信息 : %s \nSQL : %s", db_res.Error, sSQL)
		res := common.Response_Format(4, "資料新增失敗", log_data)
		return res
	}
	res := common.Response_Format(1, "資料新增成功", nil)
	return res
}
func Vod_Delete(_data string) *common.Response_format_res {

	type Vod_Delete struct {
		*Vod_Field
	}

	req_data_byte := []byte(_data)
	// req_data_byte := c.Body()
	req_data := &Vod_Delete{}
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
	sSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", vod_tablename, strings.Join(Sql_arr.WHERE, " AND "))
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
