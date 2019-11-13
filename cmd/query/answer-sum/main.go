package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	surveyColl := client.Database("lycle_line").Collection("survey_log")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"$and": []bson.M{
					bson.M{
						"enqueteID": "bn5pga9b7a4qg6818iig",
					},
					// NOTE: questionID
					bson.M{
						"bn5pga9b7a4qg6818ij0": "B",
					},
					bson.M{
						"bn5pga9b7a4qg6818ijg": "D",
					},
					bson.M{
						"bn5pga9b7a4qg6818ik0": "C",
					},
					bson.M{
						"bn5pga9b7a4qg6818ikg": "E",
					},
				},
				// "$and": []bson.M{
				// 	bson.M{
				// 		"enqueteID": "bn5pga9b7a4qg6818iig",
				// 	},
				// 	bson.M{
				// 		"$or": []bson.M{
				// 			// NOTE: questionID
				// 			bson.M{
				// 				"bn5pga9b7a4qg6818ij0": bson.M{
				// 					"$in": bson.A{"A", "B"},
				// 				},
				// 			},
				// 			bson.M{
				// 				"bn5pga9b7a4qg6818ijg": bson.M{
				// 					"$in": bson.A{"C", "D"},
				// 				},
				// 			},
				// 			bson.M{
				// 				"bn5pga9b7a4qg6818ik0": bson.M{
				// 					"$in": bson.A{"B", "C"},
				// 				},
				// 			},
				// 			bson.M{
				// 				"bn5pga9b7a4qg6818ikg": bson.M{
				// 					"$in": bson.A{"A", "B", "D", "E"},
				// 				},
				// 			},
				// 		},
				// 	},
				// },
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "null",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
	}

	surveyAggre, err := surveyColl.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalln(err)
	}

	defer surveyAggre.Close(ctx)
	for surveyAggre.Next(ctx) {
		var result bson.M
		err := surveyAggre.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
		break
	}
	if err := surveyAggre.Err(); err != nil {
		log.Fatal(err)
	}

	// 処理
	end := time.Now()
	fmt.Printf("%f秒かかった\n", (end.Sub(start)).Seconds())
}
