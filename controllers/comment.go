package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
	"time"
)

type CommentController struct {
	baseController
}

func (c *CommentController) CommentSave() {
	postId, err := strconv.Atoi(c.Input().Get("post_id"))
	fmt.Println(c.Input().Get("post_id"), c.Input().Get("content"))
	if err != nil {
		c.History(fmt.Sprintf("post id 不合法: %d", postId), "")
		return
	}
	// return
	content := c.Input().Get("content")
	//! 以后要把 post 和 comment 关联上
	comment := models.Comment{
		PostId:  postId,
		Content: content,
		Created: time.Now(),
	}
	_, err = c.o.Insert(&comment)
	if err != nil {
		c.History("评论失败："+err.Error(), "")
	} else {
		c.History("评论成功", "/article/"+c.Input().Get("post_id"))
	}
}

func (c *CommentController) CommentDelete() {
	id, err := strconv.Atoi(c.Input().Get("id"))
	postId, err := strconv.Atoi(c.Input().Get("post_id"))

	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("<script>alert('id 不合法')&amp</script>"))
	}

	_, err = c.o.Delete(&models.Comment{Id: id})

	if err != nil {
		c.History("评论删除失败："+err.Error(), "")
	} else {
		c.History("评论已删除", fmt.Sprintf("/article/%d", postId))
	}
}
