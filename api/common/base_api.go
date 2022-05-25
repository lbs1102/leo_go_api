package common

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response_format_res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*leo 2022/02/25  Validator 套件 Json欄位驗證*/
func Check_json_data_validator(data interface{}) (string, bool) {
	/* data 為結構體的指標
	* 此函式將依據結構體內的資料欄位validate定義來篩選data內的傳入參數，
	* 將會過濾掉不存在結構體內的其餘異常資料
	* 須注意format傳入格式須為結構體指標 struct(&pointer) 否則在reflect時會出問題
	* string : 錯誤信息欄位 若驗證正常則為空
	* bool : 若欄位驗證有問題則為true
	 */
	validate := validator.New()
	res := "驗證失敗"
	check := false
	if err := validate.Struct(data); err != nil {
		filed_json_name := ""
		for _, err := range err.(validator.ValidationErrors) {
			ref_val := reflect.ValueOf(data).FieldByName(err.StructField())
			varValue := ref_val.Interface()
			switch varValue.(type) {
			case string:
				varValue = ref_val.String()
			case int:
				varValue = fmt.Sprintf("%v", ref_val.Int())
			case float32, float64:
				varValue = fmt.Sprintf("%v", ref_val.Float())
			case *string:
				varValue = fmt.Sprintf("%s", ref_val.Elem())
			case *int, *float32, *float64:
				varValue = fmt.Sprintf("%v", ref_val.Elem())
			}
			if fmt.Sprintf("%v", err.Value()) == varValue {
				check = true
				ref_type, _ := reflect.TypeOf(data).FieldByName(err.StructField())
				filed_json_name = strings.Split(ref_type.Tag.Get("json"), ",")[0]
			}
		}

		res = fmt.Sprintf("%s 欄位驗證失敗", filed_json_name)
	}
	if !check {
		res = ""
	}
	return res, check
}

/*leo 2022/03/02  計算傳入值數量*/
func Get_json_data_count(struct_type interface{}, data interface{}) int {
	/* data 為結構體的指標
	* 此函式將依據結構體內的資料欄位validate定義來篩選data內的傳入參數，
	* 將會過濾掉不存在結構體內的其餘異常資料
	* 須注意format傳入格式須為結構體指標 struct(&pointer) 否則在reflect時會出問題
	* int : 欄位計算值 若驗證正常則為空
	 */
	var count_val = 0
	e := reflect.TypeOf(struct_type)
	for i := 0; i < e.NumField(); i++ {
		ValueName := e.Field(i).Name
		ref_type, _ := reflect.TypeOf(data).FieldByName(ValueName)
		ref_val := reflect.ValueOf(data).FieldByName(ValueName)
		ValuType := ref_type.Type.Kind().String()
		if ValuType == "ptr" {
			ValueCheck := ref_val.IsNil()
			if !ValueCheck {
				count_val++
			}
		} else {
			ValueCheck := ref_val.IsZero()
			if !ValueCheck {
				count_val++
			}
		}
	}
	return count_val
}

/*leo 2022/01/21 API回傳資料格式化*/
func Response_Format(code int, msg string, data interface{}) *Response_format_res {
	/* code 為狀態碼 1為正常 其餘狀態碼皆為異常 狀態碼由程式編寫者自訂
	* msg 為狀態敘述
	* data 為回傳資料或異常狀態log紀錄資料
	* 回傳成功範例:
	* {
	* 	"code": 1
	* 	"msg": "成功"
	* 	"data":	{json資料}
	* }
	* 回傳異常範例: (異常時寫入log，log檔案位置於根目錄log資料夾內)
	* {
	* 	"code": 0
	* 	"msg": "失敗"
	* 	"data":	nil
	* }
	 */
	res_data := &Response_format_res{}
	res_data.Code = code
	res_data.Msg = msg
	if code == 1 {
		if data == nil {
			res_data.Data = nil
		} else {
			check := reflect.ValueOf(data)
			if check.IsZero() {
				res_data.Data = nil
			} else {
				res_data.Data = data
			}
		}
	} else if code != 0 {
		LogError(fmt.Sprintf("\n%s", data))
	}

	return res_data
}

/*leo 2022/03/29 字串單雙引號轉譯*/
func Quotation_mark_repleace(str string) string {
	res := ""
	a, _ := regexp.Compile(`\\`)
	s_a := a.ReplaceAllString(str, "")
	r_a, _ := regexp.Compile(`'`)
	s_b := r_a.ReplaceAllString(s_a, "\\'")
	r_b, _ := regexp.Compile(`"`)
	s_final := r_b.ReplaceAllString(s_b, "\\\"")
	res = s_final
	return res
}

/*leo 2022/03/30 IP轉INT*/
func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

/*leo 2022/03/30 INT轉IP*/
func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
