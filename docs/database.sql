CREATE TABLE `users` (
  `uuid` varchar(50) DEFAULT '' COMMENT 'uuid',
  `role` varchar(1) DEFAULT '0' COMMENT '角色 0:普通用户,1:管理员',
  `username` varchar(50) DEFAULT '' COMMENT '账户',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `nickname` varchar(50) DEFAULT '昵称' COMMENT '昵称',
  `phone` varchar(20) DEFAULT '' COMMENT '手机号',
  `sex` varchar(1) DEFAULT '0' COMMENT '性别 0:男,1:女',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `header_img` varchar(50) DEFAULT 'http://xiaohubai.cn/head.png' COMMENT '头像',
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


INSERT INTO users (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `role`, `username`, `password`, `nickname`, `phone`, `sex`, `email`, `header_img`) VALUES ('2', '2020-11-26 10:23:41', '2020-11-26 10:23:41', NULL, '7614068a-3992-4874-86f0-4ae1e289b379', '1','user001', 'e10adc3949ba59abbe56e057f20f883e', '小虎','18300697959','0','','http://xiaohubai.cn/head.png');

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:varchar(50);not null;default:'';comment:'uuid'" json:"uuid"`
	Role     string `gorm:"type:varchar(1);not null;default:'0';comment:'角色 0:普通用户,1:管理员'" json:"role"`
	Username string `gorm:"type:varchar(50);not null;default:'';comment:'账户'" json:"username"`
	Password string `gorm:"type:varchar(50);not null;default:'';comment:'密码'" json:"password"`
  Nickname string `gorm:"type:varchar(50);not null;default:'昵称';comment:'昵称'" json:"nickname"`
	Phone string `gorm:"type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`
	Sex string `gorm:"type:varchar(1);not null;default:'0';comment:'性别 0:男,1:女'" json:"sex"`
	Email string `gorm:"type:varchar(50);not null;default:'';comment:'邮箱'" json:"email"`
  HeaderImg string `gorm:"type:varchar(50);not null;default:'http://xiaohubai.cn/head.png';comment:'头像'" json:"headerImg"`
}

