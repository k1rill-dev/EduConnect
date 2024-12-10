package repo

import (
	"EduConnect/internal/model"
	"EduConnect/pkg/config"
	"EduConnect/pkg/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepo struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoAccountRepo(log logger.Logger, cfg *config.Config, db *mongo.Client) *UserMongoRepo {
	return &UserMongoRepo{log: log, cfg: cfg, db: db}
}

func (m *UserMongoRepo) Create(ctx context.Context, user *model.User) error {
	_, err := m.getUserCollection().InsertOne(ctx, user, &options.InsertOneOptions{})
	if err != nil {
		m.log.Debugf("(MongoAccountRepo) error: %v", err)
		return err
	}
	return nil
}

func (m *UserMongoRepo) Update(ctx context.Context, user *model.User) error {
	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	if err := m.getUserCollection().FindOneAndUpdate(ctx, bson.M{"_id": user.Id}, bson.M{"$set": user}, ops).Err(); err != nil {
		return err
	}

	return nil
}

func (m *UserMongoRepo) GetById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := m.getUserCollection().FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserMongoRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := m.getUserCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserMongoRepo) getUserCollection() *mongo.Collection {
	return m.db.Database(m.cfg.Mongo.Db).Collection(m.cfg.MongoCollections.Users)
}
