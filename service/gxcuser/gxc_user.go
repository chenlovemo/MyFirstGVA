package gxcuser

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	//"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	//"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	//systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	//systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gxcuser"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter gxcuser.GxcSysUser, err error

type GxcUserService struct{}

var UserServiceApp = new(GxcUserService)

func (userService *GxcUserService) Register(u gxcuser.GxcSysUser) (userInter gxcuser.GxcSysUser, err error) {
	var user gxcuser.GxcSysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *GxcUserService) Login(u *gxcuser.GxcSysUser) (userInter *gxcuser.GxcSysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user gxcuser.GxcSysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		//	MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *GxcUserService) ChangePassword(u *gxcuser.GxcSysUser, newPassword string) (userInter *gxcuser.GxcSysUser, err error) {
	var user gxcuser.GxcSysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *GxcUserService) GetUserInfoList(info systemReq.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&gxcuser.GxcSysUser{})
	var userList []gxcuser.GxcSysUser

	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

//func (userService *GxcUserService) SetUserAuthority(id uint, authorityId uint) (err error) {
//
//	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&gxcuser.GxcSysUserAuthority{}).Error
//	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
//		return errors.New("该用户无此角色")
//	}
//
//	var authority system.SysAuthority
//	err = global.GVA_DB.Where("authority_id = ?", authorityId).First(&authority).Error
//	if err != nil {
//		return err
//	}
//	var authorityMenu []system.SysAuthorityMenu
//	var authorityMenuIDs []string
//	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
//	if err != nil {
//		return err
//	}
//
//	for i := range authorityMenu {
//		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
//	}
//
//	var authorityMenus []system.SysBaseMenu
//	err = global.GVA_DB.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
//	if err != nil {
//		return err
//	}
//	hasMenu := false
//	for i := range authorityMenus {
//		if authorityMenus[i].Name == authority.DefaultRouter {
//			hasMenu = true
//			break
//		}
//	}
//	if !hasMenu {
//		return errors.New("找不到默认路由,无法切换本角色")
//	}
//
//	err = global.GVA_DB.Model(&gxcuser.GxcSysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
//	return err
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *GxcUserService) DeleteUser(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&gxcuser.GxcSysUser{}).Error; err != nil {
			return err
		}

		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *GxcUserService) SetUserInfo(req gxcuser.GxcSysUser) error {
	return global.GVA_DB.Model(&gxcuser.GxcSysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *GxcUserService) SetSelfInfo(req gxcuser.GxcSysUser) error {
	return global.GVA_DB.Model(&gxcuser.GxcSysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfSetting
//@description: 设置用户配置
//@param: req datatypes.JSON, uid uint
//@return: err error

func (userService *GxcUserService) SetSelfSetting(req common.JSONMap, uid uint) error {
	return global.GVA_DB.Model(&gxcuser.GxcSysUser{}).Where("id = ?", uid).Update("origin_setting", req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user gxcuser.GxcSysUser

func (userService *GxcUserService) GetUserInfo(uuid uuid.UUID) (user gxcuser.GxcSysUser, err error) {
	var reqUser gxcuser.GxcSysUser
	//err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	//if err != nil {
	//	return reqUser, err
	//}
	//MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *GxcUserService) FindUserById(id int) (user *gxcuser.GxcSysUser, err error) {
	var u gxcuser.GxcSysUser
	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *GxcUserService) FindUserByUuid(uuid string) (user *gxcuser.GxcSysUser, err error) {
	var u gxcuser.GxcSysUser
	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ResetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *GxcUserService) ResetPassword(ID uint, password string) (err error) {
	err = global.GVA_DB.Model(&gxcuser.GxcSysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash(password)).Error
	return err
}
