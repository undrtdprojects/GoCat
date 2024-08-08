package middlewares

import "time"

type UserLoginRedis struct {
	UserId    int
	Username  string
	RoleId    int
	LoginAt   time.Time
	ExpiredAt time.Time
}
