package server

import (
	"context"
	"strings"
)

func (s *server) initMongoDBCollections(ctx context.Context) {
	err := s.mongoClient.Database(s.cfg.Mongo.Db).CreateCollection(ctx, s.cfg.MongoCollections.Accounts)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			s.log.Fatalf("(CreateCollection) err: %v", err)
		}
	}
	err = s.mongoClient.Database(s.cfg.Mongo.Db).CreateCollection(ctx, s.cfg.MongoCollections.RefreshTokens)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			s.log.Fatalf("(CreateCollection) err: %v", err)
		}
	}
}
