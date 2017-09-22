// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"github.com/b3log/solo.go/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

const tablePrefix = "b3_solo_go_"

var db *gorm.DB

func ConnectDB() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/solo.go?charset=utf8&parseTime=true&loc=Local")
	if nil != err {
		log.Error(err)

		return
	}

	tables := []interface{}{
		&model.User{}, &model.Article{}, &model.Comment{}, &model.Link{}, &model.Page{}, &model.Tag{},
		&model.Category{}, &model.Setting{}, &model.Correlation{},
	}

	db.DropTableIfExists(tables...)
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.AutoMigrate(tables...)
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Error("Disconnect from MySQL failed: " + err.Error())
	}
}
