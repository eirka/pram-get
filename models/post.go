package models

import (
	"database/sql"

	e "github.com/techjanitor/pram-get/errors"
	u "github.com/techjanitor/pram-get/utils"
)

// PostModel holds the parameters from the request and also the key for the cache
type PostModel struct {
	Thread uint
	Id     uint
	Result PostType
}

// IndexType is the top level of the JSON response
type PostType struct {
	Body Post `json:"post"`
}

type Post struct {
	ThreadId    uint    `json:"thread_id"`
	PostId      uint    `json:"post_id"`
	Name        *string `json:"name"`
	Num         uint    `json:"num"`
	Time        *string `json:"time"`
	Text        *string `json:"comment"`
	ImgId       *uint   `json:"img_id,omitempty"`
	File        *string `json:"filename,omitempty"`
	Thumb       *string `json:"thumbnail,omitempty"`
	ThumbHeight *uint   `json:"tn_height,omitempty"`
	ThumbWidth  *uint   `json:"tn_width,omitempty"`
}

// Get will gather the information from the database and return it as JSON serialized data
func (i *PostModel) Get() (err error) {

	// Initialize response header
	response := PostType{}

	// Get Database handle
	db, err := u.GetDb()
	if err != nil {
		return
	}

	post := Post{}

	err = db.QueryRow(`SELECT thread_id,posts.post_id,post_name,post_num,post_time,post_text,image_id,image_file,image_thumbnail,image_tn_height,image_tn_width
	FROM posts
	LEFT JOIN images on posts.post_id = images.post_id
	WHERE posts.post_num = ? AND posts.thread_id = ?`, i.Id, i.Thread).Scan(&post.ThreadId, &post.PostId, &post.Name, &post.Num, &post.Time, &post.Text, &post.ImgId, &post.File, &post.Thumb, &post.ThumbHeight, &post.ThumbWidth)
	if err == sql.ErrNoRows {
		return e.ErrNotFound
	} else if err != nil {
		return
	}

	// Add pagedresponse to the response struct
	response.Body = post

	// This is the data we will serialize
	i.Result = response

	return

}
