// Anx package - anx.go file specifically to reference the REST API and websocket data for the anx exchange.
// Documentation at http://docs.anxv2.apiary.io/#
//                  http://docs.anxv3.apiary.io/#

package anx

import "reflect"

// Base context default values.
const (
	v2APIURL          = "https://anxpro.com/api/2"
	v3APIURL          = "https://anxpro.com/api/3"
	tickerEndpoint    = "money/ticker"
	orderbookEndpoint = "money/depth/full"
)

// Anx holds the configuration values
type Anx struct {
	BaseURL      string
	APIKeyID     string
	APIKeySecret string
}

// New -
func New(keyID, secretKey string) *Anx {
	return &Anx{
		APIKeyID:     keyID,
		APIKeySecret: secretKey,
	}
}

// GetField gets a field from a struct
func (a *Anx) getField(structure interface{}, field string) interface{} {
	return reflect.Indirect(reflect.ValueOf(structure)).FieldByName(field).Interface()
}
