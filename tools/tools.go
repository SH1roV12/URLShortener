package tools

import (
	"fmt"

	"github.com/sqids/sqids-go"
)


func GenerateUniqueID(counter uint64)string{
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	id, _ := s.Encode([]uint64{counter})
	fmt.Println(id)
	return id
}
