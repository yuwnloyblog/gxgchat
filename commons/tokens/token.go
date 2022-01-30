package tokens

import (
	"encoding/base64"

	"github.com/yuwnloyblog/gxgchat/commons/pbdefines/pbobjs"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
)

type ImToken struct {
	AppKey string
	pbobjs.TokenValue
}

func (t ImToken) ToTokenString(secureKey []byte) (string, error) {
	tokenValue := &pbobjs.TokenValue{
		UserId:    t.UserId,
		DeviceId:  t.DeviceId,
		TokenTime: t.TokenTime,
	}
	tokenBs, err := tools.PbMarshal(tokenValue)
	if err == nil {
		encryptToken, err := encrypt(tokenBs, secureKey)
		if err == nil {
			tokenWrap := &pbobjs.TokenWrap{
				AppKey:     t.AppKey,
				TokenValue: encryptToken,
			}
			tokenWrapBs, err := tools.PbMarshal(tokenWrap)
			if err == nil {
				// encoded := base64.StdEncoding.EncodeToString(strbytes)
				bas64TokenStr := base64.StdEncoding.EncodeToString(tokenWrapBs)
				return bas64TokenStr, nil
			}
		}
	}
	return "", err
}

func encrypt(dataBs, secureKeyBs []byte) ([]byte, error) {
	return tools.AesEncrypt(dataBs, secureKeyBs)
}
func decrypt(cryptedData, secureKeyBs []byte) ([]byte, error) {
	return tools.AesDecrypt(cryptedData, secureKeyBs)
}

func ParseTokenString(tokenStr string) (*pbobjs.TokenWrap, error) {
	tokenWrap := &pbobjs.TokenWrap{}
	tokenWrapBs, err := base64.StdEncoding.DecodeString(tokenStr)
	if err == nil {
		err = tools.PbUnMarshal(tokenWrapBs, tokenWrap)
	}
	return tokenWrap, err
}
func ParseToken(tokenWrap *pbobjs.TokenWrap, secureKey []byte) (ImToken, error) {
	token := ImToken{
		AppKey: tokenWrap.AppKey,
	}
	cryptedToken := tokenWrap.TokenValue
	tokenBs, err := decrypt(cryptedToken, secureKey)
	if err == nil {
		tokenValue := &pbobjs.TokenValue{}
		err = tools.PbUnMarshal(tokenBs, tokenValue)
		if err == nil {
			token.TokenValue = *tokenValue
		}
	}
	return token, err
}

/*
decoded, err := base64.StdEncoding.DecodeString(encoded)
decodestr := string(decoded)
fmt.Println(decodestr, err)
*/
