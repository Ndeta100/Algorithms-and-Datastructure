package datastructures

import "time"

type Post struct {
	Body        string
	PublishDate time.Time
	Next        *Post //link to the next post
}

// Feed consist of post that are linked to another post
type Feed struct {
	Length int
	Start  *Post
}

func (f *Feed) Append(post *Post) {
	if f.Length == 0 {
		f.Start = post
	} else {
		currentPost := f.Start
		for currentPost.Next != nil {
			currentPost = currentPost.Next
		}
		currentPost.Next = post
	}
	f.Length++
}
