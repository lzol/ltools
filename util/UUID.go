package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID()(string,error){
	id,err := uuid.NewV4()
	return id.String(),err
}

func Get32UUID()(string,error){
	id,err := uuid.NewV4()
	uuid := strings.Replace(id.String(),"-","",-1)
	return uuid,err
}