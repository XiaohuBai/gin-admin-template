package v1

import (
	"gin-admin-template/global"
	"gin-admin-template/middleware"
	"gin-admin-template/model"
	"gin-admin-template/model/request"
	"gin-admin-template/model/response"
	"gin-admin-template/service"
	"gin-admin-template/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

//Login 登录接口
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//base64Captcha 自带验证
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		// 将数据放到model结构体,为了保证 新对象的初始数据的完整性
		U := &model.User{Username: L.Username, Password: L.Password}
		if user, err := service.Login(U); err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			token(*user, c)
		}
	} else {
		global.GVA_LOG.Error("验证码错误")
		response.FailWithMessage("验证码错误", c)
	}
}

//token 登录以后签发jwt
func token(user model.User, c *gin.Context) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		ID:         user.ID,
		UUID:       user.UUID,
		Role:       user.Role,
		Username:   user.Username,
		Phone:      user.Phone,
		BufferTime: 60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    "xiaohubai",                    // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	userInfo := response.UserInfo{
		Username:  user.Username,
		Phone:     user.Phone,
		Role:      user.Role,
		Nickname:  user.Nickname,
		UUID:      user.UUID,
		HeaderImg: user.HeaderImg,
		Sex:       user.Sex,
	}

	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			UserInfo:  userInfo,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := service.GetRedisJWT(user.Username); err == redis.Nil { //redis 没有获取到token
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			UserInfo:  userInfo,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil { //redis 获取token 异常
		global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JSONInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			UserInfo:  userInfo,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var user request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.User{Username: user.Username, Password: user.Password}
	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
