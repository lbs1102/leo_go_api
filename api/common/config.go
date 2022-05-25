package common

import (
	"github.com/gofiber/fiber/v2"
)

/* file location -> /api/config */

/* api settings */
const Name string = "mdys010"
const Port string = "3010"
const Allow_HostName string = "192.168.1.165"

/* logger settings */
const OutputDir string = "log"

/* database settings */
const Username string = "root"
const Password string = "123456"
const Hostname string = "127.0.0.1"
const Hostport string = "3306"

/* PHP 語系檔路徑 */
const PHP_langfile_Dir string = "D:\\MDYS010\\application\\lang\\"

/* Aes cbc Key Iv 設定值 */
const AES_CBC_Key string = "42c2ebc50739308ebf15a862a8a173c3"
const AES_CBC_Iv string = "8028a5bb88910c1a"

/* 2022/03/04 leo 定義api渠道與對應的資料庫 */
var Api_chennel_list = map[string]struct {
	DB_name     string
	Fiber_group fiber.Router
}{
	"develop": {"testmdys", nil},
	"api_c1":  {"mdys010", nil},
	"api_c2":  {"mdys010", nil},
}
