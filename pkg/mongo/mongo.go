package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//go:generate mockgen -destination=../../mocks/mongo_mock.go -package=mocks go_project_template/pkg/mongo MongoInf

type MongoInf interface {
	FindOne(col string, filter bson.M, ops ...*options.FindOneOptions) (*mongo.SingleResult, error)
	Find(col string, filter bson.M, ops ...*options.FindOptions) (*mongo.Cursor, error)
	FindMany(col string, filter bson.M, ops ...*options.FindOptions) (bson.A, error)
	InsertOne(col string, doc interface{}, ops ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(col string, docs []interface{}, ops ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(col string, filter interface{}, update interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(col string, filter interface{}, update interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(col string, filter interface{}, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(col string, filter interface{}, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	FindOneAndUpdate(col string, filter interface{}, update interface{}, ops ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error)
	CountDocuments(col string, filter interface{}, ops ...*options.CountOptions) (int64, error)
	Aggregate(col string, pipeline interface{}, ops ...*options.AggregateOptions) (*mongo.Cursor, error)
}

var _ MongoInf = &MongoClient{}

type PageFilter struct {
	SortBy     string
	SortMode   int8
	Limit      *int64
	Skip       *int64
	Filter     map[string]interface{}
	RegexFiler map[string]string
}

type MongoDBConfig struct {
	// 唯一连接名
	DbConnectName string                 `json:"dbconnectname"`
	DbUser        string                 `json:"dbuser"`
	DbPass        string                 `json:"dbpass"`
	DbName        string                 `json:"dbname"`
	DbHost        string                 `json:"dbhost"`
	ClientOptions *options.ClientOptions `json:"clientoptions"`
}

type MongoClient struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mgo *MongoClient

var DbClients = make(map[string]*MongoClient)

// 初始化数据库连接
func InitConnect(mongoConfig ...*MongoDBConfig) (*MongoClient, error) {
	var once sync.Once
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // bug may happen
		var (
			dbconnectname string                 = ""
			dbuser        string                 = "mongo"
			dbpass        string                 = "pass"
			dbhost        string                 = "127.0.0.1:27017"
			dbname        string                 = "local"
			clientoptions *options.ClientOptions = options.Client()
		)
		if mongoConfig != nil {
			for _, v := range mongoConfig {
				if v.DbConnectName != "" {
					dbconnectname = v.DbConnectName
				}
				if v.DbUser != "" {
					dbuser = v.DbUser
				}
				if v.DbPass != "" {
					dbpass = v.DbPass
				}
				if v.DbHost != "" {
					dbhost = v.DbHost
				}
				if v.DbName != "" {
					dbname = v.DbName
				}
				if v.ClientOptions != nil {
					clientoptions = v.ClientOptions
				}
			}
		} else {
			dbconnectname = os.Getenv("dbconnectname")
			dbuser = os.Getenv("dbuser")
			dbpass = os.Getenv("dbpass")
			if os.Getenv("dbhost") != "" {
				dbhost = os.Getenv("dbhost")
			} else {
				dbhost = os.Getenv("dbip") + ":" + os.Getenv("dbport")
			}
			dbname = os.Getenv("dbname")
		}
		connecturi := "mongodb://" + dbuser + ":" + dbpass + "@" + dbhost + "/" + dbname
		clientoptions.ApplyURI(connecturi)
		log.Println("Trying connect to: mongodb://******:******@" + dbhost + "/" + dbname)
		client, err := mongo.Connect(ctx, clientoptions)
		if err != nil {
			log.Fatal(err)
		}
		Mgo = &MongoClient{}
		Mgo.Client = client
		Mgo.Db = client.Database(dbname) // 默认为env设置的db
		// Check the connection
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		if dbconnectname != "" {
			DbClients[dbconnectname] = Mgo
		} else {
			log.Println("Warn: db connect name is not allow empty string")
		}
		log.Println("Connected to MongoDB!")
	})
	return Mgo, nil
}

// 连接到新的数据库
func (mgo *MongoClient) SetNewDb(dbname string) {
	if mgo.Client == nil {
		panic("mongo connnect empty...")
	}
	mgo.Db = mgo.Client.Database(dbname)
}

// FindOne
func (mgo *MongoClient) FindOne(col string, filter bson.M, ops ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var singleResult = table.FindOne(ctx, filter, ops...)
	if singleResult.Err() == nil {
		return singleResult, nil
	}
	return nil, singleResult.Err()
}

// Find: return mongo.Cursor
func (mgo *MongoClient) Find(col string, filter bson.M, ops ...*options.FindOptions) (*mongo.Cursor, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := table.Find(ctx, filter, ops...)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

// FindMany: return primative.A
func (mgo *MongoClient) FindMany(col string, filter bson.M, ops ...*options.FindOptions) (bson.A, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := table.Find(ctx, filter, ops...)
	if err != nil {
		return nil, err
	}
	var result bson.A
	defer cur.Close(context.TODO())
	for cur.Next(context.Background()) {
		var r bson.M
		if err = cur.Decode(&r); err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}
	return result, nil
}

// InsertOne
func (mgo *MongoClient) InsertOne(col string, doc interface{}, ops ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.InsertOne(ctx, doc, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InsertMany
func (mgo *MongoClient) InsertMany(col string, docs []interface{}, ops ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.InsertMany(ctx, docs, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateOne
func (mgo *MongoClient) UpdateOne(col string, filter interface{}, update interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.UpdateOne(ctx, filter, update, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateMany
func (mgo *MongoClient) UpdateMany(col string, filter interface{}, update interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.UpdateMany(ctx, filter, update, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteOne
func (mgo *MongoClient) DeleteOne(col string, filter interface{}, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.DeleteOne(ctx, filter, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteMany
func (mgo *MongoClient) DeleteMany(col string, filter interface{}, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := table.DeleteMany(ctx, filter, ops...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneAndUpdate
func (mgo *MongoClient) FindOneAndUpdate(col string, filter interface{}, update interface{}, ops ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result *mongo.SingleResult
	err := table.FindOneAndUpdate(ctx, filter, update, ops...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CountDocuments
func (mgo *MongoClient) CountDocuments(col string, filter interface{}, ops ...*options.CountOptions) (int64, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return -1, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := table.CountDocuments(ctx, filter, ops...)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (mgo *MongoClient) Aggregate(col string, pipeline interface{}, ops ...*options.AggregateOptions) (*mongo.Cursor, error) {
	if mgo.Db == nil || mgo.Client == nil {
		return nil, fmt.Errorf("no init connect and db info! ")
	}
	table := mgo.Db.Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := table.Aggregate(ctx, pipeline, ops...)
	if err != nil {
		return nil, err
	}
	return cur, nil
}
