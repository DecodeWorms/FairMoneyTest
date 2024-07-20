package storage

import (
	"context"
	"fairmoneytest/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const (
	transactionCollections = "transactions"
)

var _ DataStore = &MongoStorage{}

type MongoStorage struct {
	mongoClient  *mongo.Client
	databaseName string
}

func New(connectURI, databaseName string) (DataStore, *mongo.Client, error) {
	//Connecting to MongoDB...
	log.Println("Connecting to MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, err
	}
	//Connected to MongoDB successfully ..
	log.Println("Connected to mongoDB Successfully")
	return &MongoStorage{mongoClient: client, databaseName: databaseName}, client, nil
}

func (repo MongoStorage) RecordTransaction(data *model.TransactionRecords) error {
	_, err := repo.col(transactionCollections).InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (repo MongoStorage) GetTransactionByID(ID string) (*model.TransactionRecords, error) {
	filter := bson.M{"id": ID}
	var trx model.TransactionRecords
	if err := repo.col(transactionCollections).FindOne(context.Background(), filter).Decode(&trx); err != nil {
		return nil, err
	}
	return &trx, nil
}

func (repo MongoStorage) GetTransactionByReference(ref string) (*model.TransactionRecords, error) {
	filter := bson.M{"reference": ref}
	var trx model.TransactionRecords
	if err := repo.col(transactionCollections).FindOne(context.Background(), filter).Decode(&trx); err != nil {
		return nil, err
	}
	return &trx, nil
}

func (repo MongoStorage) UpdateAccountBalance(ID string, data *model.TransactionRecords) error {
	oldRec, err := repo.GetTransactionByID(ID)
	if err != nil {
		return err
	}
	filter := bson.M{"id": ID}
	_, err = repo.col(transactionCollections).ReplaceOne(context.TODO(), filter, repo.buildAccountBalancePayload(data, oldRec))
	if err != nil {
		return err
	}
	return nil

}

func (repo MongoStorage) col(collectionName string) *mongo.Collection {
	return repo.mongoClient.Database(repo.databaseName).Collection(collectionName)
}

func (repo MongoStorage) buildAccountBalancePayload(newTrx, oldTrx *model.TransactionRecords) *model.TransactionRecords {
	if newTrx == nil {
		return oldTrx
	}
	if newTrx.Balance != 0 {
		oldTrx.Balance = newTrx.Balance
	}
	return oldTrx
}
