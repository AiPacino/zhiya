package models

import "github.com/jinzhu/gorm"

// 触发指令
type CmdType struct {
	gorm.Model
	TypeFlag string `gorm:"size:50;unique_index" json:"type_flag"` //定义标识，方便系统回调，命令进入回调队列
	TypeName string `gorm:"size:50" json:"type_name"`              //定义名称
}

/*
  获取触发指令
  如果没有则初始化
*/

func (c *CmdType) Ensure(db *gorm.DB, typeFlag string) error {
	return db.Where(CmdType{TypeFlag: typeFlag}).FirstOrInit(c).Error
}

/*
  群指令
*/
type ChatRoomCmd struct {
	gorm.Model
	ChatRoomSerialNo string `gorm:"size:100" json:"chat_room_serial_no"`
	CmdType          string `gorm:"size:50;index:idx_cmd_type" json:"cmd_type"`
	CmdValue         string `gorm:"size:100" json:"cmd_value"`
	CmdReply         string `gorm:"type:text(10000)" json:"cmd_reply"`
	IsOpen           bool   `gorm:"index:idx_is_open" json:"is_open"`
}

/*
  供应商指令
*/
type MyCmd struct {
	gorm.Model
	MyId     string `gorm:"size:100" json:"my_id"`
	CmdType  string `gorm:"size:50;index:idx_cmd_type" json:"cmd_type"`
	CmdValue string `gorm:"size:100" json:"cmd_value"`
	CmdReply string `gorm:"type:text(10000)" json:"cmd_reply"`
	IsOpen   bool   `gorm:"index:idx_is_open" json:"is_open"`
}

/*
  代理商指令
*/
type SubCmd struct {
	gorm.Model
	SubId    string `gorm:"size:100" json:"sub_id"`
	CmdType  string `gorm:"size:50;index:idx_cmd_type" json:"cmd_type"`
	CmdValue string `gorm:"size:100" json:"cmd_value"`
	CmdReply string `gorm:"type:text(10000)" json:"cmd_reply"`
	IsOpen   bool   `gorm:"index:idx_is_open" json:"is_open"`
}

// TAG指令
type TagCmd struct {
	gorm.Model
	TagId    string `gorm:"size:100" json:"tag_id"`
	CmdType  string `gorm:"size:50;index:idx_cmd_type" json:"cmd_type"`
	CmdValue string `gorm:"size:100" json:"cmd_value"`
	CmdReply string `gorm:"type:text(10000)" json:"cmd_reply"`
	IsOpen   bool   `gorm:"index:idx_is_open" json:"is_open"`
}
