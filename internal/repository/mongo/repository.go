package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Histories interface {
	Save()
}

type Repositories struct {
	Histories Histories
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Histories: NewHistoryRepositories(db),
	}
}
