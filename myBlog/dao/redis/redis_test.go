package redis

import (
	"fmt"
	"testing"
)

func init(){
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err : %v\n", err)
		return
	}
	fmt.Println("connect redis success!")
}

func Test(t *testing.T) {

}