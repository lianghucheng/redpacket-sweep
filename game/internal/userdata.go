package internal

import "github.com/name5566/leaf/log"

type UserData struct {
	UserID        int `bson:"_id"`
	AccountID     int
	Account       string
	Password      string
	CircleID      int  // 圈圈ID
	Online        bool // 在线
	Nickname      string
	Headimgurl    string
	Sex           int // 1 男, 2 女
	Lang          string
	IPAddr        string
	Token         string
	ExpiredAt      int64 // Token 过期时间
	Role          int
	Chips         int64 // 筹码
	WinChips      int64 // 赢得的筹码
	Win           int   // 胜局
	GameId        int   // 总对局
	SubsidizedAt  int64 // 补助时间
	CreatedAt     int64
	UpdatedAt     int64
	LoginIP       string
	FreeChangedAt int64 // 免费重置时间
	UnionID       string
}

func updateUserData(id int, update interface{}) {
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)

		_, err := db.DB(DB).C("users").UpsertId(id, update)
		if err != nil {
			log.Error("update user %v data error: %v", id, err)
		}
	}, nil)
}
