package service

import (
	"encoding/csv"
	"github.com/astaxie/beego/logs"
	"mobileNip/util"
	"os"
)

var mobiles []MobileItem

type MobileItem struct {
	Id           int    `json:"id"`
	MobileRegion string `json:"mobile_region"`
	Province     string `json:"province"`
	City         string `json:"city"`
	CardType     string `json:"card_type"`
	AreaZone     string `json:"area_zone"`
	ZipCode      string `json:"zip_code"`
}

func InitMobileResource() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("init mobile error: ", err)
		}
	}()
	ReadMobile()
}

func ReadMobile() {
	filename := "/resources/mobile.csv"
	path, err := util.GetFileAbsolutePath(filename)
	if err != nil {
		logs.Error("get absolute path err: ", err)
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		logs.Error("open file err: ", err)
		panic(err)
	}

	defer file.Close()

	w := csv.NewReader(file)
	datas, err := w.ReadAll()
	if err != nil {
		logs.Error("read mobile csv err: ", err)
		panic(err)
	}

	mobiles = make([]MobileItem, 0)
	for _, data := range datas[1:] {
		item := MobileItem{
			Id:           util.ParseInt(data[0]),
			MobileRegion: data[1],
			Province:     data[2],
			City:         data[3],
			CardType:     data[4],
			AreaZone:     data[5],
			ZipCode:      data[6],
		}

		mobiles = append(mobiles, item)
	}

	logs.Info("read mobile count: %d", len(mobiles))
}

func FindMobileInfo(mobile string) *MobileItem {
	var res *MobileItem

	tmp := mobile[:7]
	for _, item := range mobiles {
		if item.MobileRegion == tmp {
			res = &item
			break
		}
	}

	return res
}

func FindMobileInfoByBinarySearch(mobile string, subMobiles []MobileItem) *MobileItem {
	if subMobiles == nil {
		subMobiles = mobiles
	}

	i := len(subMobiles) / 2
	mobileRegion := util.ParseInt(subMobiles[i].MobileRegion)
	userMobileRegion := util.ParseInt(mobile[:7])

	if mobileRegion == userMobileRegion {
		return &subMobiles[i]
	}

	if mobileRegion > userMobileRegion {
		return FindMobileInfoByBinarySearch(mobile, subMobiles[:i])
	}

	if mobileRegion < userMobileRegion {
		return FindMobileInfoByBinarySearch(mobile, subMobiles[i:])
	}

	return nil
}
