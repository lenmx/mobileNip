package util

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strings"
)

type logConfig struct {
	filename string // 保存的文件名
	maxlines int    // 每个文件保存的最大行数，默认值 1000000
	maxsize  int    // 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	daily    bool   // 是否按照每天 logrotate，默认是 true
	maxdays  int    // 文件最多保存多少天，默认保存 7 天
	rotate   bool   // 是否开启 logrotate，默认是 true
	level    string // 日志保存的时候的级别，默认是 Trace 级别
	perm     string // 日志文件权限
}

func InitLog() {
	logpath := `./logs/` + beego.AppConfig.String("appname") + `.log`
	//logpath, _ = GetFileAbsolutePath(logpath)

	logs.SetLogger(logs.AdapterConsole)
	//if beego.AppConfig.String("runmode") != "dev" {
	logs.SetLogger(logs.AdapterMultiFile, `{
"filename":"`+logpath+`",
"level":7,
"daily":true,
"maxdays":10,
"color":true,
"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]
}`)

	//}

	logs.Async()
}

func GetIPInfo(clientIP string) string {
	if clientIP == "" {
		return ""
	}

	ips := strings.Split(clientIP, ",")
	res := ""
	for i := range ips {
		ip := strings.TrimSpace(ips[i])
		desc := "" // GetDescFromIP(ip)
		ipstr := fmt.Sprintf("%s: %s", ip, desc)
		if i != len(ips)-1 {
			res += ipstr + " -> "
		} else {
			res += ipstr
		}
	}

	return res
}

func getIPFromRequest(req *http.Request) string {
	clientIP := req.Header.Get("x-forwarded-for")
	if clientIP == "" {
		ipPort := strings.Split(req.RemoteAddr, ":")
		if len(ipPort) >= 1 && len(ipPort) <= 2 {
			clientIP = ipPort[0]
		} else if len(ipPort) > 2 {
			idx := strings.LastIndex(req.RemoteAddr, ":")
			clientIP = req.RemoteAddr[0:idx]
			clientIP = strings.TrimLeft(clientIP, "[")
			clientIP = strings.TrimRight(clientIP, "]")
		}
	}

	return GetIPInfo(clientIP)
}

func LogInfoByContext(ctx *context.Context, f string, v ...interface{}) {
	ipString := fmt.Sprintf("(%s) ", getIPFromRequest(ctx.Request))
	logs.Info(ipString+f, v...)
}

func LogWarningByContext(ctx *context.Context, f string, v ...interface{}) {
	ipString := fmt.Sprintf("(%s) ", getIPFromRequest(ctx.Request))
	logs.Warning(ipString+f, v...)
}
