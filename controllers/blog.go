package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strconv"
	"time"
)

type BlogController struct {
	baseController
}

func (c *BlogController) list() {
	var (
		page       int
		pageSize   int = 6
		offset     int
		list       []*models.Post
		categoryId int
		keyword    string
	)

	// 获取全部目录
	categories := []*models.Category{}
	c.o.QueryTable(new(models.Category).TableName()).All(&categories)
	c.Data["categories"] = categories

	// 从 controller 里获取数据
	// page 第几页, 博文所属的目录 category_id，搜索关键词 keyword
	page, _ = c.GetInt("page")
	categoryId, _ = c.GetInt("category_id")
	keyword = c.Input().Get("keyword")

	// 如果页数小于 1，则设为第一页
	if page < 1 {
		page = 1
	}

	offset = (page - 1) * pageSize

	// query 全部博文
	query := c.o.QueryTable(new(models.Post).TableName())

	// 根据目录id过滤
	if categoryId != 0 {
		query = query.Filter("category_id", categoryId)
	}

	// 根据关键词过滤
	if keyword != "" {
		query = query.Filter("title_contains", keyword)
	}

	// 如果是首页的请求，则从博文里筛选 is_top 的文章
	if c.actionName == "home" {
		query = query.Filter("is_top", 1)
	}

	// list：请求具体某页
	query.OrderBy("-updated").Limit(pageSize, offset).All(&list)

	// 获取当页的个数
	count, _ := query.Count()

	c.Data["count"] = count
	c.Data["list"] = list

}

func getPost(c *BlogController) models.Post {
	id, err := strconv.Atoi(c.Ctx.Input.Params()[":id"])
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("<script>alert('id 不合法')&amp</script>"))
	}
	post := models.Post{Id: id}
	err = c.o.Read(&post)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("<script>找不到id对应的文章</script>"))
	}
	return post
}

func (c *BlogController) Home() {
	config := models.Config{Name: "start"}
	c.o.Read(&config, "Name")

	// if config.Value != "1" {
	// 	c.Ctx.WriteString("系统维护中...")
	// 	return
	// }

	var notices []*models.Post

	c.o.QueryTable(new(models.Post).TableName()).
		// Filter("category_id", 2).
		All(&notices)
	c.Data["notices"] = notices

	c.list()

	c.TplName = c.controllerName + "/home.html"
	utils.Debugger("controllerName", c.controllerName)
}

func (c *BlogController) Article() {
	c.TplName = c.controllerName + "/article.html"
	post := getPost(c)
	utils.Debugger("show post", "\ntitle: %s\ncontent:%s", post.Title, post.Content)
	c.Data["Post"] = post
	var comments []*models.Comment

	c.o.QueryTable("comment").Filter("post_id", post.Id).OrderBy("-created").All(&comments)
	c.Data["comments"] = comments
}
func (c *BlogController) Detail()  {}
func (c *BlogController) About()   {}
func (c *BlogController) Comment() {}

func (c *BlogController) Create() {
	c.TplName = c.controllerName + "/create.html"
}

func (c *BlogController) Edit() {
	c.TplName = c.controllerName + "/edit.html"
	post := getPost(c)
	c.Data["Post"] = post
}

func (c *BlogController) Save() {
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	post := models.Post{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
		Views:   0,
		IsTop:   1,
		Url:     "",
	}
	fmt.Println(post)

	_, err := c.o.Insert(&post)

	if err != nil {
		c.History("博文插入数据错误："+err.Error(), "")
	} else {
		c.History("已创建成功", "/")
	}
}

func (c *BlogController) Update() {
	post := getPost(c)

	post.Title = c.Input().Get("title")
	post.Content = c.Input().Get("content")
	post.Updated = time.Now()

	_, err := c.o.Update(&post)
	if err != nil {
		c.History("修改失败："+err.Error(), "")
	} else {
		c.History("修改成功", "/")
	}
}

func (c *BlogController) Delete() {
	post := getPost(c)

	_, err := c.o.Delete(&post)
	if err != nil {
		utils.Debugger("删除失败：", err.Error())
		c.History("删除失败："+err.Error(), "")
	} else {
		utils.Debugger("删除成功", "")
		c.History("删除成功", "/")
	}
}
