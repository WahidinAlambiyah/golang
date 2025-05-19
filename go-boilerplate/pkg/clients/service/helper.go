/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package service

import "github.com/WahidinAlambiyah/golang/pkg/utils/constants"

var client *ServiceClient

func init() {
	client = &ServiceClient{
		name: constants.FEATURE_SERVICE,
	}
}

func GetClient() *ServiceClient {
	return client
}
