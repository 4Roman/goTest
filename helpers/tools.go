package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/gusleein/golog"
	"go.mongodb.org/mongo-driver/bson"
	"goTest/db"
	"goTest/models"
)

func PiplineTop100(field string) bson.A {
	return bson.A{
		bson.M{"$group": bson.M{"_id": "$" + field, "count": bson.M{"$sum": 1}}},
		bson.M{"$sort": bson.M{"count": -1}},
		bson.M{"$limit": 100},
	}
}
func StringFromPipeline(ctx context.Context, pipeline any) string {
	var result string
	cursor, err := db.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Error(err)
	}
	defer cursor.Close(ctx)
	var results []models.TopResult
	if err = cursor.All(ctx, &results); err != nil {
		log.Error(err)
	}
	for _, r := range results {
		result += fmt.Sprintf("%s: %d\n", r.ID, r.Count)
	}
	return result
}

func PiplineTop100WithTwoFields(field1, field2 string) bson.A {
	pipeline := bson.A{
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					field1: "$" + field1,
					field1: "$" + field1,
				},
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{"count": -1},
		}, //osname 1
		bson.M{
			"$group": bson.M{
				"_id": "$_id." + field1,
				"top_versions": bson.M{"$push": bson.M{
					"version": "$_id." + field2,
					"count":   "$count",
				}},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":          0,
				field1:         "$_id",
				"top_versions": bson.M{"$slice": []interface{}{"$top_versions", 100}},
			},
		},
	}
	return pipeline
}

func StringFromPipelineWith2Field(ctx context.Context, pipeline any) []byte {
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
