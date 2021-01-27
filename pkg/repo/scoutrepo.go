package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// ScoutServer wraps the dependcies needed to access Scount info
type ScoutServer struct {
	client *mongo.Client
}

// ConfirmScoutClient yada yada
func ConfirmScoutClient() error {
	var scoutServer = new(ScoutServer)
	err := scoutServer.client.Ping(context.Background(), nil)
	return err
}
