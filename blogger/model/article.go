package model

import "time"

// 定义文章结构体
// 加载首页提高效率
type ArticleInfo struct {
	Id   int64  `db:"id"`
	CategoryId int64 `db:"category_id"`
	// 文章摘要
	Summary   string    `db:"summary"`
	Title   string    `db:"title"`
	ViewCount uint32	`db:"view_count"`
	CommentCount uint32	`db:"comment_count"`
	UserName	string    `db:"username"`
	// 创建时间
	CreateTime	time.Time	`db:"create_time"`
}

// 文章详情页结构体
type ArticleDetail struct {
	ArticleInfo
	Category
	Content	string    `db:"content"`
	
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}