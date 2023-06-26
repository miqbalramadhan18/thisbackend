package thisbackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertJalurPenerimaan(db string, jalurpenerimaan JalurPenerimaan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("jalurpenerimaan").InsertOne(context.TODO(), jalurpenerimaan)
	if err != nil {
		fmt.Printf("InsertJalurPenerimaan: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertInformasi(db string, informasi Informasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("informasi").InsertOne(context.TODO(), informasi)
	if err != nil {
		fmt.Printf("InsertInformasi: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertBiaya(db string, biaya Biaya) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("biaya").InsertOne(context.TODO(), biaya)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
	}
	return insertResult.InsertedID
}
// func InsertAbout(db string, about About) (insertedID interface{}) {
// 	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
// 	if err != nil {
// 		fmt.Printf("InsertAbout: %v\n", err)
// 	}
// 	return insertResult.InsertedID
// }
// func InsertContactus(db string, contacus Contactus) (insertedID interface{}) {
// 	insertResult, err := MongoConnect(db).Collection("contactus").InsertOne(context.TODO(), contacus)
// 	if err != nil {
// 		fmt.Printf("InsertContactus: %v\n", err)
// 	}
// 	return insertResult.InsertedID
// }
func GetDataJalurPenerimaan(stats string) (data []JalurPenerimaan) {
	user := MongoConnect("qobel").Collection("jalurpenerimaan")
	filter := bson.M{"jalurtes": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetJalurPenerimaan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataInformasi(stats string) (data []Informasi) {
	user := MongoConnect("qobel").Collection("informasi")
	filter := bson.M{"catatan": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataInformasi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataBiaya(stats string) (data []Biaya) {
	user := MongoConnect("qobel").Collection("biaya")
	filter := bson.M{"biayasemester": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDatabiaya :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

