package repository

import (
	"blog_filtration_feature/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogFiltrationRepository struct {
	Database		*mongo.Database
	collection		string
}

func NewFiltrationRepo(dataBase *mongo.Database, collection string) domain.BlogFiltrationRepository {
	return &BlogFiltrationRepository{
		Database : dataBase,
		collection : collection,
	}
}

func (bf *BlogFiltrationRepository) FilterBlog(blogRequest *domain.BlogRequest) ([]*domain.BlogResponse, error) {
	collection := bf.Database.Collection(bf.collection)
	filter := bson.M{}
	if blogRequest.LikeLowerRange > 0 {
		filter["like"] = bson.M{
			"$gt" :blogRequest.LikeLowerRange,
		}
	}

	if blogRequest.ViewLowerRange > 0 {
		filter["view"] = bson.M{
			"$gt" :blogRequest.ViewLowerRange,
		}
	}

	if blogRequest.Date != nil  {
		filter["created_at"] = blogRequest.Date
	}

	if blogRequest.Tags != nil {
		filter["tags"] = bson.M {
			"$in" : blogRequest.Tags,
		}
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var blogResponse []*domain.BlogResponse
	for cur.Next(context.TODO()) {
		var singleResponse domain.BlogResponse
		err := cur.Decode(&singleResponse)
		if err != nil {
			return nil, err
		}

		blogResponse = append(blogResponse, &singleResponse)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return blogResponse, nil
}