package tools

import (
	"fmt"

	"github.com/sqids/sqids-go"
)

//Генерация уникального ID для БД
func GenerateUniqueID(counter uint64)string{

	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	//Создание ID по количеству строк в БД+1 
	id, _ := s.Encode([]uint64{counter})
	fmt.Println(id)
	return id
}
