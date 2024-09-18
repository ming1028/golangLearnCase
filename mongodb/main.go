package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type student struct {
	Name string
	Age  int32
}

func main() {
	// 链接配置
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOption.SetMaxPoolSize(25)
	// 链接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Sprintf("connect mongodb err:%v\n", err)
		return
	}
	// 链接检查
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Sprintf("ping mongodb err:%v\n", err)
		return
	}

	fmt.Println("connect mongodb success")

	// 指定数据库
	collection := client.Database("smh").Collection("student")
	s1 := student{
		Name: "张三1",
		Age:  22,
	}
	s2 := student{
		Name: "王无误",
		Age:  44,
	}
	s3 := student{
		Name: "李四四",
		Age:  55,
	}
	insertResult, err := collection.InsertOne(ctx, s1)
	if err != nil {
		fmt.Sprintf("insertone err:%v\n", err)
		return
	}
	fmt.Println(insertResult, insertResult.InsertedID)

	students := []interface{}{s2, s3}
	insertManyResult, err := collection.InsertMany(ctx, students)
	if err != nil {
		fmt.Sprintf("insertMany err:%v\n", err)
		return
	}
	fmt.Println(insertManyResult.InsertedIDs)

	//更新
	filter := bson.D{{
		"name", "张三1",
	}}
	update := bson.D{
		{"$inc", bson.D{{"age", 11}}},
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Sprintf("updateOne err:%v\n", err)
		return
	}
	fmt.Println(updateResult)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	var s student
	err = collection.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		fmt.Sprintf("findOne err:%v\n", err)
		return
	}
	fmt.Println(s)

	// 查询多个
	findOpetion := options.Find()
	findOpetion.SetLimit(2)

	var stus []student
	cur, err := collection.Find(ctx, bson.D{{}}, findOpetion)
	if err != nil {
		fmt.Sprintf("find err:%v\n", err)
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var ele student
		err = cur.Decode(&ele)
		if err != nil {
			fmt.Sprintf("Decode err:%v\n", err)
			return
		}
		stus = append(stus, ele)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stus)

	fmt.Println("删除")
	delResult, err := collection.DeleteOne(ctx, bson.D{{"name", "李四"}})
	if err != nil {
		fmt.Sprintf("del err:%v\n", err)
		return
	}
	fmt.Println(delResult)

	delMany, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		fmt.Sprintf("del err:%v\n", err)
		return
	}
	fmt.Println(delMany)
}
