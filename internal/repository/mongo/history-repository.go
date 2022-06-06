package mongo

import "go.mongodb.org/mongo-driver/mongo"

type HistoryRepository struct {
	db *mongo.Database
}

func NewHistoryRepositories(db *mongo.Database) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}

func (r *HistoryRepository) Save() {

}
