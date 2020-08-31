package biz

import "time"

func GetBizDate() int {
	yyyy, mm, dd := time.Now().UTC().Date()
	return yyyy*10000 + int(mm)*100 + dd
}

func GetBizTime() int {
	hh, mm, ss := time.Now().UTC().Clock()
	return hh*10000 + mm*100 + ss
}

func GetBizDatetime() (int, int) {
	now := time.Now().UTC()
	yyyy, mm, dd := now.Date()
	dateNow := yyyy*10000 + int(mm)*100 + dd

	hh, mm1, ss := now.Clock()
	timeNow := hh*10000 + mm1*100 + ss

	return dateNow, timeNow
}
