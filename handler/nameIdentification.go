package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg/util"
	"encoding/json"
	_ "errors"
	"github.com/labstack/echo"
	"log"
	_ "net/http"
)

func NameIdentificationHandler(c echo.Context) error {

	teamList := make([]*domain.TeamName, 0, 0)
	log.Println(teamList)

	// ファイル１を取得→チェックを行う
	bytes, err := util.ReadFile("json/folder.json")
	if err != nil {
		return err
	}

	// ファイル２を取得→チェックを行う

	// ファイルを名寄せする

	// CSVファイルを作成する

	// ダウンロード

	list := make([]*domain.Folder, 0, 0)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}

	log.Println(list)

	return c.File("json/folder.json")
}

func nameIdentification(list1 []*domain.TeamName, list2 []*domain.TeamName) ([]*domain.NameIdentificationFile, error) {
	// whiteswanに委託するところ
	return nil, nil
}
