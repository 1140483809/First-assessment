package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	_ "xorm.io/xorm"
)

func main(){
	engine,err := xorm.NewEngine("mysql","root:Yy13883129603@/sqlsql?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//学生管理系统的学生数据
	type stundent_2 struct {
		Id int `xorm:"pk"` //学号
		Name string
		Age int
		School string  //学院
		Money float64
	}
	err = engine.Sync2(new(stundent_2))
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	for true {
		var code,yanzheng int
		var sno,sage int = 0 , 0
		var x string
	//	aa := false
		fmt.Println("选择你的操作：")
		fmt.Println("1.查找")
		fmt.Println("2.插入")
		fmt.Println("3.删除")
		fmt.Println("4.更新")
		fmt.Println("5.退出")
		fmt.Scanln(&code)
		switch code {
			case 1:
				fmt.Println("是否通过学号查询数据(Y/N)")
				YY:fmt.Scanln(&x)
				if !(x=="Y" || x=="N")  {
					fmt.Println("输入格式错误，请重新输入！")
					fmt.Println(x)
					goto YY
				}
				if x=="Y" {
					fmt.Printf("请输入想要查看学生的学号:")
					fmt.Scanln(&sno)
					chaxun := new(stundent_2)
					_,err = engine.Id(sno).Get(chaxun)
					if err != nil {
						log.Fatal(err.Error())
						return
					}
					fmt.Printf("Id:%d,Name:%s,Age:%d,School:%s,Money:%f",chaxun.Id,chaxun.Name,chaxun.Age,chaxun.School,chaxun.Money) //Id int/Name string/Age int/School string/Money float
				}else {
					var total int64
					chaxun := make([]stundent_2,0)
					fuzhu := new(stundent_2)
					fmt.Println("通过年龄查询(年龄+判断)(1大于，2小于)")
					fmt.Scanln(&sage,&yanzheng)
					if yanzheng==1 {
						err = engine.Where("age > ?",sage).Find(&chaxun)
						if err != nil {
							log.Fatal(err.Error())
							return
						}
						total, err = engine.Where("age > ?",sage).Count(fuzhu)
					}
					if yanzheng==2 {
						err = engine.Where("age < ?",sage).Find(&chaxun)
						if err != nil {
							log.Fatal(err.Error())
							return
						}
						total, err = engine.Where("age < ?",sage).Count(fuzhu)
					}
					fmt.Println("Id\tName\tAge\tSchool\tMoney\t")
					var i int64=0
					for ; i < total ; i++ {
						fmt.Printf("%d\t%s\t%d\t%s\t%f\n",chaxun[i].Id,chaxun[i].Name,chaxun[i].Age,chaxun[i].School,chaxun[i].Money)
					}
					fmt.Println(chaxun)
				}
			case 2:
				fmt.Println("是否插入多条数据(Y/N)")
				YN:fmt.Scanln(&x)
				if !(x=="Y" || x=="N")  {
					fmt.Println("输入格式错误，请重新输入！")
					fmt.Println(x)
					goto YN
				}
				if x=="N" {
					fmt.Println("Id int/Name string/Age int/School string/Money float(请一行输入):")
					a := new(stundent_2)
					fmt.Scanln(&a.Id,&a.Name,&a.Age,&a.School,&a.Money)
					_,err := engine.Insert(a)
					if err != nil {
						log.Fatal(err.Error())
						return
					}
				}else {
					a := make([]stundent_2,20)
					fmt.Println("输入最后一条数据后，双击回车结束输入(最多一次输入20条)")
					fmt.Println("Id int/Name string/Age int/School string/Money float(一组请一行输入):")
					for i := 0; i < 20; i++ {
						fmt.Scanln(&a[i].Id,&a[i].Name,&a[i].Age,&a[i].School,&a[i].Money)
						if a[i].Id==0 {
							b := make([]stundent_2, i)
							for j := 0; j < i; j++ {
								b[j] = a[j]
							}
							_ , err = engine.Insert(b)
							yanzheng = 1
							if err != nil {
								log.Fatal(err.Error())
								return
							}
							break
						}
					}
					if yanzheng != 1 {
						_,err := engine.Insert(a)
						if err != nil {
							log.Fatal(err.Error())
							return
						}
					}
				}
				fmt.Println("插入成功！")
			case 3:
				fmt.Printf("请输入想删除学生的学号：")
				fmt.Scanln(&sno)
				chaxun := new(stundent_2)
				_,err = engine.Id(sno).Get(chaxun)
				delete := new(stundent_2)
				_ , err := engine.Id(sno).Delete(delete)
				if err != nil {
					log.Fatal(err.Error())
					return
				}
				fmt.Printf("已删除:")
				fmt.Println(chaxun)
			case 4:
				fmt.Printf("请输入想要修改学生的学号：")
				fmt.Scanln(&sno)
				chaxun := new(stundent_2)
				_,err = engine.Id(sno).Get(chaxun)
				update := new(stundent_2)
				fmt.Println("请输入新的数据")
				fmt.Println("Id int/Name string/Age int/School string/Money float(一组一行输入):")
				fmt.Scanln(&update.Id,&update.Name,&update.Age,&update.School,&update.Money)
				_ , err := engine.Id(sno).Update(update)
				if err != nil {
					log.Fatal(err.Error())
					return
				}
				fmt.Println("将",chaxun,"改为",update)
			case 5:
				return
		}
	}
}