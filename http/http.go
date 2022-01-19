package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sql "gsWeb/sql"
	"html/template"
	"log"
	"net/http"
)

func TestGin() {
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"Blog":   "www.flysnow.org",
	//		"wechat": "flysnow_org",
	//	})
	//})
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "static/index.html", gin.H{})
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	r.Run(":8080")
}

func HttpStart() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/testFun1", getAjax)
	log.Fatal(http.ListenAndServe("localhost:9001", nil))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
func getAjax(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success")
	r.ParseForm()

	// 第一种方式
	// username := request.Form["username"][0]
	// password := request.Form["password"][0]

	// 第二种方式
	id := r.Form.Get("id")
	fmt.Println(id)
	// query := r.URL.Query()
	// id := query["id"][0]
	// fmt.Println(id)

}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.ConnectSql()
	if err == nil {
		fmt.Println("数据库连接成功")
		//查询数据库
		rows, rowsErr := sql.QueryRow(db)
		if rowsErr != nil {
			fmt.Printf("查询失败")
		} else {
			//获取数据
			for rows.Next() {
				var count string
				rows.Scan(&count)
				fmt.Fprintln(w, "Row"+count)
			}
		}
	} else {
		fmt.Println("数据库连接失败，请查询参数")
	}
}
