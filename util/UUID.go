package util

import "github.com/satori/go.uuid"

func GetUUID()(string,error){
	id,err := uuid.NewV4()
	return id.String(),err
}
