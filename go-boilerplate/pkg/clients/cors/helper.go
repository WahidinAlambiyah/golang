/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cors

import "github.com/WahidinAlambiyah/golang/pkg/utils/constants"

var client *CorsClient

func init() {
	client = &CorsClient{
		name: constants.FEATURE_CORS,
	}
}

func GetClient() *CorsClient {
	return client
}
