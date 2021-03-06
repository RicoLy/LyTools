package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	a, err := xormadapter.NewAdapter("mysql", "root:ly@tcp(127.0.0.1:3306)/test?charset=utf8", true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, _ := casbin.NewEnforcer("./CasbinDemo/rbac_models.conf", a)

	//从DB加载策略
	_ = e.LoadPolicy()

	//获取router路由对象
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		name := c.Query("name")
		api := c.Query("api")
		fmt.Printf("%s ==%s \n", name, api)
		fmt.Println("增加Policy")
		if ok, _ := e.AddPolicy(name, api, "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})

	// 添加角色
	r.POST("/api/v1/addRole", func(c *gin.Context) {
		name := c.Query("name")
		role := c.Query("role")
		fmt.Println("增加Policy")
		if ok, _ := e.AddRoleForUser(name, role); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		name := c.Query("name")
		api := c.Query("api")
		fmt.Println("删除Policy")
		if ok, _ := e.RemovePolicy(name, api, "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})

	//删除policy
	r.DELETE("/api/v1/deleteRole", func(c *gin.Context) {
		name := c.Query("name")
		role := c.Query("role")
		fmt.Println("删除Policy")
		if ok, _ := e.DeleteRoleForUser(name, role); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
	//使用自定义拦截器中间件
	r.Use(Authorize(e))
	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	//创建请求
	r.GET("/api/v1/hello2", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	_ = r.Run(":9000") //参数为空 默认监听8080端口
}


//拦截器
func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		name := c.Query("name")
		//获取请求的URI
		obj := c.Request.URL.Path
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := name
		fmt.Printf("name==%s; obj==%s; act==%s", name, obj, act)
		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}