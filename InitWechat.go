package requests

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func InitWechat(accessTokenHandle credential.AccessTokenHandle) {
	WC := wechat.NewWechat()
	memory := cache.NewMemory()

	cfg := &offConfig.Config{
		AppID:          CONFIG.Wechat.AppID,
		AppSecret:      CONFIG.Wechat.AppSecret,
		Token:          CONFIG.Wechat.Token,
		EncodingAESKey: CONFIG.Wechat.AesKey,
		Cache:          memory,
	}
	OA = WC.GetOfficialAccount(cfg)
	OA.SetAccessTokenHandle(accessTokenHandle)
}
