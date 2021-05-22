# First-assessment
//用户
     
     //http://localhost:8080/user/register  注册
     POST
     {
        "username":"",
        "password":""
     }
     //http://localhost:8080/user/login   登录
     POST
    {
      "username":"",
      "password":""
    }
     //http://localhost:8080/user/exit   退出登录
     GET
     //http://localhost:8080/user/revise   修改密码
    POST
    {
      "password":""
    }


//文章

        //http://localhost:8080/user/artical?id=?   用户id发表文章
        POST
        {
           "a_title":"这是一篇关于祭奠袁老先生的文章",
           "a_arti":"袁隆平……"
        }
	//http://localhost:8080/user/artical_see/:aid  显示id为？aid的文章
        GET
	//http://localhost:8080/user/artical_dianzan/:aid   给id为aid的文章点赞
        GET
	//http://localhost:8080/user/artical_allcomment/:aid  显示aid全部评论
        GET
	//http://localhost:8080/user/artical_shoucang/:aid?id=?  用户id收藏aid
        GET


//评论

	//http://localhost:8080/user/comment/:aid?id=?   给文章aid评论
        POST
        {
           "p_comment":""
        }
	//http://localhost:8080/user/comment/com/:pid?id=?   回复pid
        POST
        {
           "p_comment":""
        }
	//http://localhost:8080/user/comment_dianzan/:pid   给pid点赞
        GET
	//http://localhost:8080/user/:id  id为用户名即username  显示id的发布的
        GET
