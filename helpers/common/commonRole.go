package common

import "GoCat/helpers/constant"

func CheckRole(roleId int, action string) bool {
	if roleId == constant.AdminRoleIdUser.Int() {
		return true
	} else if roleId != constant.AdminRoleIdUser.Int() {
		if action == constant.ReadActionUser.String() {
			return true
		}
	}
	return false
}
