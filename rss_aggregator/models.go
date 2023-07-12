package main

import (
	"time"

	"github.com/0xOthmane/golang_projects/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func dbFeedFToFeedF(dbFeedF database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedF.ID,
		CreatedAt: dbFeedF.CreatedAt,
		UpdatedAt: dbFeedF.UpdatedAt,
		UserID:    dbFeedF.UserID,
		FeedID:    dbFeedF.FeedID,
	}
}

func dbFeedFsToFeedFs(dbFeedFs []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedF := range dbFeedFs {
		feedFollows = append(feedFollows, dbFeedFToFeedF(dbFeedF))
	}

	return feedFollows
}
