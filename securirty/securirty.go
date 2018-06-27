package securirty

import (
	"io/ioutil"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

func Autenticar() []byte {
	bytes, _ := ioutil.ReadFile("./securirty/sample_key.priv")

	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(time.Duration(10) * time.Second))

	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(rsaPrivate)
	return b
}

func Verificar(token string) bool {

	bytes, _ := ioutil.ReadFile("./securirty/sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes)

	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		return false
	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		return false
	}

	return true
}
