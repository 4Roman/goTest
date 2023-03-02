package helpers

import (
	"context"
	"encoding/json"
	log "github.com/gusleein/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goTest/db"
)

// пайплайн для топ 100 по одному полю
func PiplineTop100(field string) bson.A {
	return bson.A{
		bson.M{"$group": bson.M{"_id": "$" + field, "count": bson.M{"$sum": 1}}},
		bson.M{"$sort": bson.M{"count": -1}},
		bson.M{"$limit": 100},
	}
}

func BytesFromPipeline(ctx context.Context, pipeline any) []byte {
	cursor, err := db.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error(err)
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Error(err)
	}
	textRsult, err := json.Marshal(results)
	if err != nil {
		log.Error(err)
	}
	return textRsult
}

func PipelineForOS() mongo.Pipeline {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", bson.D{{"name", "$os_name"}, {"ver", "$os_version"}}},
			{"count", bson.D{{"$sum", 1}}},
		}}},
		{{"$sort", bson.D{{"count", -1}}}},
		{{"$group", bson.D{
			{"_id", "$_id.name"},
			{"top_versions", bson.D{{"$push", "$_id.ver"}}},
		}}},
		{{"$project", bson.D{
			{"_id", 0},
			{"name", "$_id"},
			{"top_versions", bson.D{{"$slice", []interface{}{"$top_versions", 100}}}},
		}}},
	}
	return pipeline
}

func PipelineForBrowser() mongo.Pipeline {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", bson.D{{"name", "$browser_name"}, {"ver", "$browser_ver"}}},
			{"count", bson.D{{"$sum", 1}}},
		}}},
		{{"$sort", bson.D{{"count", -1}}}},
		{{"$group", bson.D{
			{"_id", "$_id.name"},
			{"top_versions", bson.D{{"$push", "$_id.ver"}}},
		}}},
		{{"$project", bson.D{
			{"_id", 0},
			{"name", "$_id"},
			{"top_versions", bson.D{{"$slice", []interface{}{"$top_versions", 100}}}},
		}}},
	}
	return pipeline
}

func PipelineForPhone() mongo.Pipeline {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", bson.D{{"brand", "$phone_brand"}, {"model", "$phone_model"}}},
			{"count", bson.D{{"$sum", 1}}},
		}}},
		{{"$sort", bson.D{{"count", -1}}}},
		{{"$group", bson.D{
			{"_id", "$_id.brand"},
			{"top_models", bson.D{{"$push", "$_id.model"}}},
		}}},
		{{"$project", bson.D{
			{"_id", 0},
			{"brand", "$_id"},
			{"top_models", bson.D{{"$slice", []interface{}{"$top_models", 100}}}},
		}}},
	}
	return pipeline
}
