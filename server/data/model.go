package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//A ToDoList struct with three fields, which can be used to represent a MongoDB document with an _id, a task description, and a completion status.
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

// get all task from the DB and return it
func GetAllTask() []primitive.M {
	cur, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func InsertTask(task ToDoList) primitive.ObjectID {
	insertResult, err := Collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	return id
}

// task complete method, update task's status to true
func CompleteTask(task string) bool {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false

	}
	return true
}

// task undo method, update task's status to false
func UndoTask(task string) bool {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	_, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// delete one task from the DB, delete by ID
func DeleteTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	_, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

// delete all the tasks from the DB
func DeleteAllTasks() int64 {
	d, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	return d.DeletedCount
}

// task update method, update task's status to false
func UpdateTask(task ToDoList, taskId string) bool {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"task": task.Task}}
	_, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
