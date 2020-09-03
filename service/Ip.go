package service

import (
	"encoding/csv"
	"github.com/astaxie/beego/logs"
	"mobileNip/util"
	"os"
	"strings"
)

var IpsMap map[string][]*IpItem

type IpItem struct {
	Ip       string `json:"ip"`
	Province string `json:"province"`
}

func InitIpResource() {
	ReadIp()
}

func ReadIp() {
	filename := "/resources/ip-china.csv"
	path, err := util.GetFileAbsolutePath(filename)
	if err != nil {
		logs.Error("get file absolute path err: ", err)
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		logs.Error("open file err: ", err)
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			logs.Error("ip init error: ", err)
		}
		file.Close()
	}()

	w := csv.NewReader(file)
	datas, err := w.ReadAll()
	if err != nil {
		panic(err)
	}

	IpsMap = make(map[string][]*IpItem)
	for _, data := range datas[1:] {
		i := strings.Index(data[0], ".")
		ipKey := data[0][:i]
		item := IpItem{
			Ip:       data[0],
			Province: data[1],
		}

		_map, ok := IpsMap[ipKey]
		if ok {
			_map = append(_map, &item)
			IpsMap[ipKey] = _map
		} else {
			IpsMap[ipKey] = []*IpItem{&item}
		}
	}

	logs.Info("read ip count: %d", len(IpsMap))
}

func FindIpInfo(ip string) *IpItem {
	var res *IpItem

	firstIndex := strings.Index(ip, ".")
	lastIndex := strings.LastIndex(ip, ".")

	key := ip[:firstIndex]
	ip = ip[:lastIndex+1] + "0"

	_map, ok := IpsMap[key]
	if !ok {
		return nil
	}

	for _, item := range _map {
		if item.Ip == ip {
			res = item
			break
		}
	}

	return res
}
