package BasicPractice

import "fmt"


type ITvAdInfoDB interface {
CreateTvAdInfo(a int )
}


type TvAdInfoDB struct{}

func (TAIDB *TvAdInfoDB) CreateTvAdInfo( a int ){

	fmt.Println("你没想到我虽然没有成员，但是有方法233333\n printf(int +1)", b+1 );
}

func NewTvAdInfoDB() ITvAdInfoDB {

	return &TvAdInfoDB{}
}

func tempmain() {

 	NewTvAdInfoDB().CreateTvAdInfo(4)
 	fmt.Println("测试结束，练练跳转吧同学");

}