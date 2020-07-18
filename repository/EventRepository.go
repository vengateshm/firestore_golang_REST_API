package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
)

type EventRepository interface {
	CreateEvent(event *entity.Event) (*entity.Event, error)
	GetAllEvents() (*[]entity.Event, error)
}

type repo struct{}

func NewEventRepository() EventRepository {
	return &repo{}
}

const (
	project_id       = "fir-samples-92b83"
	event_collection = "GoAppEvents"
)

func (*repo) CreateEvent(event *entity.Event) (*entity.Event, error) {
	ctx := context.Background()
	client, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(event_collection).Add(ctx, map[string]interface{}{
		"name":  event.Name,
		"venue": event.Venue,
	})
	if err != nil {
		log.Fatalf("failed to create new event : %v", err)
		return nil, err
	}
	return event, nil
}

func (r *repo) GetAllEvents() (*[]entity.Event, error) {
	ctx := context.Background()
	client, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var events []entity.Event
	docIterator := client.Collection(event_collection).Documents(ctx)
	for {
		docSnapshot, err := docIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to retrieve list of events : %v", err)
			return nil, err
		}
		event := entity.Event{
			Name:  docSnapshot.Data()["name"].(string),
			Venue: docSnapshot.Data()["venue"].(string),
		}
		events = append(events, event)
	}

	return &events, nil
}

func GetFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_id)
	if err != nil {
		log.Fatalf("failed to create firestore client : %v", err)
		return nil, err
	}
	return client, nil
}
