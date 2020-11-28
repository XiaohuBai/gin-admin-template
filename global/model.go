/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-25 23:45:28
 * @Description: 描述
 */
package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
