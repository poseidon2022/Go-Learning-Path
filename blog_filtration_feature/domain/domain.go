package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogResponse struct {
	ID									primitive.ObjectID				`json:"id" bson:"_id"`
	Title       						string							`json:"title" bson:"title"`
	Description 						string							`json:"description" bson:"description"`
	UpdatedAt   						time.Time						`json:"updated_at" bson:"updated_at"`
	CreatedAt							time.Time						`json:"created_at" bson:"created_at"`
}

type BlogRequest struct {
	Tags								[]string
	LikeLowerRange						uint
	ViewLowerRange						uint
	Date								*time.Time
}

type BlogFiltrationUseCase interface {
	FilterBlog(*BlogRequest)			([]*BlogResponse, error)
}

type BlogFiltrationRepository interface {
	FilterBlog(*BlogRequest)			([]*BlogResponse, error)
}