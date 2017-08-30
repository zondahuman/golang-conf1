package test


import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"strconv"
	"testing"
)

var db *leveldb.DB

const ID_FIELD  =  "id"

func init(){
	var err error
	db,err = leveldb.OpenFile("./db",nil)
	if (err != nil) {
		panic(err)
	}

	_,err = db.Get([]byte(ID_FIELD),nil)
	if (err!=nil) {
		db.Put([]byte(ID_FIELD),[]byte("10000"),nil)
	}
}

func GetNextId() int {
	ids,err := db.Get([]byte(ID_FIELD),nil)
	if (err != nil) {
		fmt.Println(err)
	}
	id := Byte2int(ids)
	db.Put([]byte(ID_FIELD),Int2byte(id+1),nil)
	return id
}

func Byte2int(val []byte) int {
	var result int
	result,_ = strconv.Atoi(string(val))
	return result
}

func Int2byte(val int) []byte {
	result := []byte(strconv.Itoa(val))
	return result
}


func Test_levelDbGetId(t *testing.T) {
	id := GetNextId()
	fmt.Println("id=", id)
}
