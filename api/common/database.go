package common

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Sql_str_struct struct {
	Index  int
	WHERE  []string
	UPDATE struct {
		Filed []string
		Value []string
	}
	CREATE struct {
		Filed []string
		Value []string
	}
}

var DB_chennel_list = map[string]*gorm.DB{
	"develop": nil,
	"api_c1":  nil,
	"api_c2":  nil,
}

func InitDB() {
	for k, v := range Api_chennel_list {
		DNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", Username, Password, Hostname, Hostport, v.DB_name)
		Database, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error())
			LogPanic(err.Error())
		}
		DB_chennel_list[k] = Database
	}
}
func Select_db_for_api(API_name string) {
	Db = DB_chennel_list[API_name]
}

/*leo 建立增、刪、改、查 所需的sql字段陣列*/
func Get_sql_str_arr(struct_data interface{}) Sql_str_struct {
	count := count_data(reflect.TypeOf(struct_data), reflect.ValueOf(struct_data))
	sql_data := Sql_str_struct{}
	sql_data.CREATE.Filed = make([]string, count)
	sql_data.CREATE.Value = make([]string, count)
	sql_data.data_to_sql(reflect.TypeOf(struct_data), reflect.ValueOf(struct_data))
	return sql_data
}

/*leo 計算資料內實際有效欄位數量(用來初始化[]string長度)*/
func count_data(ref_type reflect.Type, ref_val reflect.Value) (Index int) {
	e := ref_type.Elem()
	f := ref_val.Elem()
	for i := 0; i < e.NumField(); i++ {
		ValueName := e.Field(i).Name
		if !strings.Contains(ValueName, "Sql_") {
			if strings.Contains(ValueName, "_Field") {
				test_child_type := e.Field(i).Type
				test_child_val := f.Field(i)
				if !test_child_val.IsNil() {
					Index = count_data(test_child_type, test_child_val)
				}
			} else {
				ref_type, _ := e.FieldByName(ValueName)
				ref_val := f.FieldByName(ValueName)
				ValuType := ref_type.Type.Kind().String()
				if ValuType == "ptr" {
					if !ref_val.IsNil() {
						Index++
					}
				} else {
					if !ref_val.IsZero() {
						Index++
					}
				}
			}
		}
	}
	return Index
}

/*leo 增、刪、改、查 sql字段遞迴函式(用來建立sql陣列)*/
func (sql_arr *Sql_str_struct) data_to_sql(ref_type reflect.Type, ref_val reflect.Value) {

	e := ref_type.Elem()
	f := ref_val.Elem()
	for i := 0; i < e.NumField(); i++ {
		ValueName := e.Field(i).Name
		if !strings.Contains(ValueName, "Sql_") {
			if strings.Contains(ValueName, "_Field") {
				test_child_type := e.Field(i).Type
				test_child_val := f.Field(i)
				if !test_child_val.IsNil() {
					sql_arr.data_to_sql(test_child_type, test_child_val)
				}
			} else {
				ref_type, _ := e.FieldByName(ValueName)
				ref_val := f.FieldByName(ValueName)
				ValusJsonName := strings.Split(ref_type.Tag.Get("json"), ",")[0]
				varValue := ref_val.Interface()
				ValuType := ref_type.Type.Kind().String()
				if ValuType == "ptr" {
					if !ref_val.IsNil() {
						switch varValue.(type) {
						case *string:
							Quotation_value := Quotation_mark_repleace(fmt.Sprintf("%s", ref_val.Elem()))
							sql_arr.WHERE = append(sql_arr.WHERE, fmt.Sprintf("%s = '%s'", ValusJsonName, Quotation_value))
							sql_arr.UPDATE.Value = append(sql_arr.UPDATE.Value, fmt.Sprintf("%s = '%s'", ValusJsonName, Quotation_value))
							sql_arr.CREATE.Filed[sql_arr.Index] = fmt.Sprintf("%v", ValusJsonName)
							sql_arr.CREATE.Value[sql_arr.Index] = fmt.Sprintf("'%s'", Quotation_value)
							sql_arr.Index++
						case *int, *float32, *float64:
							sql_arr.WHERE = append(sql_arr.WHERE, fmt.Sprintf("%s = %v", ValusJsonName, ref_val.Elem()))
							sql_arr.UPDATE.Value = append(sql_arr.UPDATE.Value, fmt.Sprintf("%s = %v", ValusJsonName, ref_val.Elem()))
							sql_arr.CREATE.Filed[sql_arr.Index] = fmt.Sprintf("%v", ValusJsonName)
							sql_arr.CREATE.Value[sql_arr.Index] = fmt.Sprintf("%v", ref_val.Elem())
							sql_arr.Index++
						}
					}
				} else {
					if !ref_val.IsZero() {
						switch ref_val.Interface().(type) {
						case string:
							Quotation_value := Quotation_mark_repleace(fmt.Sprintf("%s", ref_val))
							sql_arr.UPDATE.Filed = append(sql_arr.UPDATE.Filed, fmt.Sprintf("%s = '%s'", ValusJsonName, Quotation_value))
							sql_arr.CREATE.Filed[sql_arr.Index] = fmt.Sprintf("%v", ValusJsonName)
							sql_arr.CREATE.Value[sql_arr.Index] = fmt.Sprintf("'%s'", Quotation_value)
							sql_arr.Index++
						case int, float32, float64:
							sql_arr.UPDATE.Filed = append(sql_arr.UPDATE.Filed, fmt.Sprintf("%s = %v", ValusJsonName, ref_val))
							sql_arr.CREATE.Filed[sql_arr.Index] = fmt.Sprintf("%v", ValusJsonName)
							sql_arr.CREATE.Value[sql_arr.Index] = fmt.Sprintf("%v", ref_val)
							sql_arr.Index++
						}
					}
				}
			}

		}
	}
}
