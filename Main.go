package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
	"strconv"
	"time"
	_ "time"
	_ "xorm.io/xorm"
)

type hhh struct {
	Id int `xorm:"autoincr not null pk INT"`
	Name string `xorm:"unique" json:"username"`
	CreatedAt time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type info struct {
	Username string `xorm:"pk" json:"username"`
	A_id int `xorm:"pk" json:"a_id"`
	P_id int `xorm:"pk" json:"p_id"`
}

type user struct {
	Username string `xorm:"pk" json:"username"`
	Password string `json:"password"`
}

type artical struct {
	A_id int `xorm:"autoincr pk" json:"a_in"`
	A_name string `xorm:"not null" json:"username"`
	A_title string `json:"a_title"`
	A_arti string `json:"a_arti"`
	A_dianzanshu int
	CreatedAt time.Time `xorm:"created"`
}

type pinglun struct {
	P_id int `xorm:"autoincr pk" json:"p_in"`
	P_name string `xorm:"not null" json:"username"`
	P_comment string `json:"p_comment"`
	P_dianzanshu int
	CreatedAt time.Time `xorm:"created"`
}

type shoucang struct {
	S_name string `xorm:"not null" json:"username"`
	S_aid int `xorm:"pk" json:"aid"`
}

type pinglun_2 struct {
	P_id2 int `xorm:"autoincr pk" json:"p_id2"`
	P_id int `json:"p_id"`
	P_name2 string `json:"username"` //评论人
	P_name string `json:"p_name"` //被评论的人
	P_comment string `json:"p_comment"`
}

var conn = "root:Yy13883129603@/sqlsql?charset=utf8"
var chaxun = make(map[string]user)
var yhm string
var aid int
var engine_sql *xorm.Engine

func main()  {
	engine := gin.Default()


	//连接sqlsql数据库
	 //tcp(127.0.0.1:3306),utf8
	 var err error
	engine_sql,err = xorm.NewEngine("mysql",conn)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = engine_sql.Sync2(new(hhh)) //同步表
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(user))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(artical))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(pinglun))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(info))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(pinglun_2))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	err = engine_sql.Sync2(new(shoucang))
	if err != nil {
		log.Fatal(err.Error())
		return
	}


	err = engine_sql.Find(&chaxun)
	if err != nil {
		log.Fatal(err.Error())
		return
	}


//	engine.LoadHTMLFiles("C:/Users/11404/Desktop/html/1-2104241H541/index.html")

	//http://localhost:8080/user
	routegroup := engine.Group("/user")


	//用户
	//http://localhost:8080/user/register  注册
	//{
	//  "username":,
	//  "password":
	// }
	routegroup.POST("/register", register)
	//http://localhost:8080/user/login   登录
	routegroup.POST("login",login)
	//http://localhost:8080/user/exit   退出登录
	routegroup.GET("exit", func(context *gin.Context) {
		yhm=""
		context.Writer.WriteString("已退出登录!")
		return
	})
	//http://localhost:8080/user/revise   修改密码
	//
	routegroup.POST("/revise",revise)


	//文章
	//http://localhost:8080/user/artical?id=?   用户id发表文章
	routegroup.POST("artical",uartical)
	//http://localhost:8080/user/artical_see/:aid  显示id为？aid的文章
	routegroup.GET("artical_see/:aid",sartical)
	//http://localhost:8080/user/artical_dianzan/:aid   给id为aid的文章点赞
	routegroup.GET("artical_dianzan/:aid",a_dianzan)
	//http://localhost:8080/user/artical_allcomment/:aid  显示aid全部评论
	routegroup.GET("artical_allcomment/:aid",a_allcom)
	//http://localhost:8080/user/artical_shoucang/:aid?id=?  用户id收藏aid
	routegroup.GET("/artical_shoucang/:aid",a_shoucang)


	//评论
	//http://localhost:8080/user/comment/:aid?id=?   给文章aid评论
	routegroup.POST("comment/:aid",comment)
	//http://localhost:8080/user/comment/com/:pid?id=?   回复pid
	routegroup.POST("/comment/com/:pid",com_com)
	//http://localhost:8080/user/comment_dianzan/:pid   给pid点赞
	routegroup.GET("comment_dianzan/:pid",c_dianzan)
	//http://localhost:8080/user/display/:id  id为用户名即username  显示id的发布的
	routegroup.GET("/display/:id",display)

	//显示SQL操作
//	engine_sql.ShowSQL()

	engine.Run()
}

func a_shoucang(context *gin.Context) {
	fullPath := "收藏文章" + context.FullPath()
	fmt.Println(fullPath)
	sc := new(shoucang)
	sc.S_name = context.Query("id")
	s_id := context.Param("aid")
	var err1 error
	sc.S_aid , err1 = strconv.Atoi(s_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	if sc.S_name != yhm{
		context.Writer.WriteString("请先登录!")
		return
	}
	engine_sql.Insert(sc)
	context.String(200,"收藏成功")
}

func com_com(context *gin.Context) {
	fullPath := "回复评论" + context.FullPath()
	fmt.Println(fullPath)
	com := new(pinglun)
	com2 := new(pinglun_2)
	c_id := context.Param("pid")
	if yhm != c_id {
		context.Writer.WriteString("请先登录!")
		return
	}
	var err1 error
	com2.P_id , err1 =strconv.Atoi(c_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	if err := context.BindJSON(com2);err !=nil{
		log.Fatal(err.Error())
		return
	}
	engine_sql.Id(com2.P_id).Get(com)
	com2.P_name=com.P_name
	com2.P_name2 = context.Query("id")
	_ , err := engine_sql.Insert(com2)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	context.String(200,"评论成功")
}

func a_allcom(context *gin.Context) {
	fullPath := "显示文章评论" + context.FullPath()
	fmt.Println(fullPath)
	art := new(artical)
	a_id := context.Param("aid")
	var err1 error
	art.A_id , err1 = strconv.Atoi(a_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	engine_sql.Where("a_id = ?",art.A_id).Get(art)
	context.String(http.StatusOK,"文章全部评论\n")
	context.JSON(200,gin.H{
		"id号":art.A_id,
		"标题":art.A_title,
		"作者":art.A_name,
		"正文":"…………",
		"点赞数":art.A_dianzanshu,
	})
	context.String(http.StatusOK,"\n")
	in := make([]info,0)
	engine_sql.Where("a_id = ?",art.A_id).Find(&in)
	in_1 := new(info)
	total , err := engine_sql.Where("a_id = ?",art.A_id).Count(in_1)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	var j int64
	for j = 0; j < total; j++ {
		com := new(pinglun)
		engine_sql.Id(in[j].P_id).Get(com)
		context.JSON(200,gin.H{
			"id号":com.P_id,
			"用户":com.P_name,
			"评论":com.P_comment,
			"点赞数":com.P_dianzanshu,
		})
		context.Writer.WriteString("\n")
	}
}

func c_dianzan(context *gin.Context) {
	fullPath := "评论点赞" + context.FullPath()
	fmt.Println(fullPath)
	com := new(pinglun)
	c_id := context.Param("pid")
	var err1 error
	com.P_id , err1 = strconv.Atoi(c_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	engine_sql.Where("p_id = ?",com.P_id).Get(com)
	com.P_dianzanshu++
	engine_sql.Id(com.P_id).Update(com)
	context.String(200,"点赞成功！")
}

func a_dianzan(context *gin.Context) {
	fullPath := "文章点赞" + context.FullPath()
	fmt.Println(fullPath)
	art := new(artical)
	a_id := context.Param("aid")
	var err1 error
	art.A_id , err1 = strconv.Atoi(a_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	engine_sql.Where("a_id = ?",art.A_id).Get(art)
	art.A_dianzanshu++
	engine_sql.Id(art.A_id).Update(art)
	context.String(200,"点赞成功！")
}

func sartical(context *gin.Context) {
	fullPath := "显示文章" +context.FullPath()
	fmt.Println(fullPath)
	art := new(artical)
	a_id := context.Param("aid")
	var err1 error
	art.A_id , err1 = strconv.Atoi(a_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	sc := new(shoucang)
	total , err := engine_sql.Where("s_aid = ?",art.A_id).Count(sc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	engine_sql.Where("a_id = ?",art.A_id).Get(art)
	context.String(http.StatusOK,"文章页面\n")
	context.JSON(200,gin.H{
		"id号":art.A_id,
		"标题":art.A_title,
		"作者":art.A_name,
		"正文":art.A_arti,
		"点赞数":art.A_dianzanshu,
		"收藏数":total,
	})
	context.String(200,"\n")
	com := new(pinglun)
	info := new(info)
	engine_sql.Where("a_id = ?",art.A_id).Get(info)
	engine_sql.Where("p_id = ?",info.P_id).Get(com)
	if com.P_id == 0 {
		context.String(200,"还没有评论。")
	}else{
		context.String(200,"评论：\n")
		context.JSON(200,gin.H{
			"id号":com.P_id,
			"用户":com.P_name,
			"评论":com.P_comment,
			"点赞数":com.P_dianzanshu,
		})
	}
}

func display(context *gin.Context) {
	fullPath := "显示用户界面" + context.FullPath()
	fmt.Println(fullPath)
	user := new(user)
	com := make([]pinglun,0)
	art := make([]artical,0)
	com_1 := new(pinglun)
	art_1 := new(artical)
	user.Username = context.Param("id")
	engine_sql.Where("a_name = ?",user.Username).Find(&art)
	total , err := engine_sql.Where("a_name = ?",user.Username).Count(art_1)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	sc := new(shoucang)
	context.String(http.StatusOK,"这是%s的页面\n",user.Username)
	var i,j,k int64
	context.JSON(200,gin.H{
		"发布文章数":total,
	})
	context.String(200,"\n发布的文章：\n")
	for i = 0; i < total; i++ {
		total1 , err := engine_sql.Where("s_aid = ?",art[i].A_id).Count(sc)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		context.JSON(200,gin.H{
			"id号":art[i].A_id,
			"标题":art[i].A_title,
			"作者":art[i].A_name,
			"正文":art[i].A_arti,
			"发布时间":art[i].CreatedAt,
			"点赞数":art[i].A_dianzanshu,
			"收藏数":total1,
		})
		context.Writer.WriteString("\n")
	}
	sc1 := make([]shoucang,0)
	engine_sql.Where("s_name = ?",user.Username).Find(&sc1)
	total2 , _ :=engine_sql.Where("s_name = ?",user.Username).Count(sc)
	context.String(200,"收藏的文章：\n")
	for i = 0; i < total2; i++ {
		engine_sql.Where("a_id = ?",sc1[i].S_aid).Find(&art)
		context.JSON(200,gin.H{
			"标题":art[i+total].A_title,
		})
		context.Writer.WriteString("\n")
	}
	engine_sql.Where("p_name = ?",user.Username).Find(&com)
	total , err = engine_sql.Where("p_name = ?",user.Username).Count(com_1)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	inf := new(info)
	context.String(200,"发布的评论：\n")
	for j = 0; j < total; j++ {
		engine_sql.Where("p_id = ?",com[j].P_id).Get(inf)
		context.JSON(200,gin.H{
			"id号":com[j].P_id,
			"评论的文章号":inf.A_id,
			"用户":com[j].P_name,
			"评论":com[j].P_comment,
			"发布时间":com[j].CreatedAt,
			"点赞数":com[j].P_dianzanshu,
		})
		context.Writer.WriteString("\n")
	}
	com_2 := new(pinglun_2)
	com2 := make([]pinglun_2,0)
	engine_sql.Where("p_name = ?",user.Username).Find(&com2)
	total , err = engine_sql.Where("p_name = ?",user.Username).Count(com_2)
	for k = 0; k < total; k++ {
		context.JSON(200,gin.H{
			"id_2号":com2[k].P_id2,
			"id号":com2[k].P_id,
			"被回复人":com2[k].P_name,
			"用户":com2[k].P_name2,
			"回复":com2[k].P_comment,
		})
		context.Writer.WriteString("\n")
	}
}

func revise(context *gin.Context) {
	fullPath := "修改密码" + context.FullPath()
	fmt.Println(fullPath)
	user := new(user)
	if err := context.BindJSON(user); err != nil {
		log.Fatal(err.Error())
		return
	}
	if user.Username != yhm{
		context.Writer.WriteString("请先登录!")
		return
	}
	_ , err := engine_sql.Id(user.Username).Update(user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	context.Writer.WriteString("密码修改成功!")
	fmt.Println("密码更改为:" + user.Password)
}

func comment(context *gin.Context) {
	fullPatn := "发表评论" + context.FullPath()
	fmt.Println(fullPatn)
	if _,ok :=chaxun[yhm]; !ok {
		context.Writer.WriteString("请先登录!")
		return
	}
	com := new(pinglun)
	com_user := new(info)
	com.P_name = context.Query("id")
	a_id := context.Param("aid")
	if com.P_name != yhm{
		context.Writer.WriteString("请先登录!")
		return
	}
	var err1 error
	com_user.A_id , err1 = strconv.Atoi(a_id)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}
	if err := context.BindJSON(com);err !=nil{
		log.Fatal(err.Error())
		return
	}
	_ , err := engine_sql.Insert(com)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	com_user.Username = com.P_name
	com_user.P_id = com.P_id
	_ , err = engine_sql.Insert(com_user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	context.Writer.WriteString("发布成功！")
}

func uartical(context *gin.Context) {
	fullPatn := "发表文章" + context.FullPath()
	fmt.Println(fullPatn)
//	if _,ok :=chaxun[yhm]; !ok{
//		context.Writer.WriteString("请先登录!")
//		return
//	}
	art := new(artical)
	art.A_name = context.Query("id")
	if yhm != art.A_name {
		context.Writer.WriteString("请先登录!")
		return
	}
	if err := context.BindJSON(art);err !=nil{
		log.Fatal(err.Error())
		return
	}
	fmt.Println(art)
	_ , err := engine_sql.Insert(art)
	if err !=nil{
		log.Fatal(err.Error())
		return
	}
	aid = art.A_id
	context.Writer.WriteString("发布成功！")
}

func login(context *gin.Context) {
	fullPath := "用户登录" + context.FullPath()
	fmt.Println(fullPath)
	user := new(user)
	if err := context.BindJSON(user); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	if _,ok := chaxun[user.Username]; !ok {
		context.Writer.WriteString("查无此人!")
		return
	}
	if chaxun[user.Username] != *user {
		context.Writer.WriteString("密码错误!")
		return
		}
	yhm = user.Username
	context.Writer.WriteString("登录成功!")
}

func register(context *gin.Context) {
	fullpath := "用户注册" + context.FullPath()
	fmt.Println(fullpath)
	user := new(user)
	if err := context.BindJSON(&user); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	if _,ok := chaxun[user.Username]; ok{
		context.Writer.WriteString("用户已存在!")
		return
	}
	context.JSON(200,gin.H{
		"username":user.Username,
		"password":user.Password,
	})
	_,err := engine_sql.Insert(user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	context.Writer.WriteString("注册成功!")
}


//{
//    "username":"11404",
//    "password":"138831"
//}
