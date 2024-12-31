package service

import (
	"fmt"
	"github.com/bojurgess/bard/internal/database"
	"github.com/bojurgess/bard/internal/model"
	"log"
	"sync"
	"time"
)

var BrokerService = &brokerService{
	subscribers:  make(map[string][]chan any),
	stopChannels: make(map[string]chan bool),
}

type brokerService struct {
	mut          sync.RWMutex
	subscribers  map[string][]chan any
	stopChannels map[string]chan bool
}

func (b *brokerService) Publish(userId string, msg any) {
	b.mut.RLock()
	defer b.mut.RUnlock()

	for _, ch := range b.subscribers[userId] {
		select {
		case ch <- msg:
		default:
		}
	}
}

func (b *brokerService) BeginBroadcasting(userId string) {
	log.Printf("Beginning broadcast for user: %s", userId)
	stopCh := make(chan bool)
	b.stopChannels[userId] = stopCh

	tokens, err := database.TokenService.Find(userId)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to begin broadcasting: %v", err))
		return
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				currentlyPlaying, err := SpotifyService.GetCurrentlyPlaying(tokens.AccessToken)
				if err != nil {
					if err.Error() == "The access token expired" {
						oauthTokens, err := SpotifyService.RefreshAccessToken(tokens.RefreshToken)
						if err != nil {
							fmt.Println(fmt.Errorf("failed to refresh access token: %v", err))
						}
						tokens = model.OAuthToDatabaseTokens(oauthTokens, userId)
						err = database.TokenService.Update(tokens)
						if err != nil {
							fmt.Println(fmt.Errorf("failed to update db tokens: %v", err))
						}
						continue
					}

					errMsg := fmt.Errorf("failed to broadcast currently playing: %v", err)
					b.Publish(userId, "stop")
					delete(b.stopChannels, userId)
					delete(b.subscribers, userId)
					fmt.Println(errMsg)
					continue
				}
				b.Publish(userId, currentlyPlaying)
			case <-stopCh:
				log.Printf("Stopping broadcast for user with id: %s", userId)
				delete(b.stopChannels, userId)
				delete(b.subscribers, userId)
				return
			}
		}
	}()
}

func (b *brokerService) Subscribe(userId string) <-chan any {
	b.mut.Lock()
	defer b.mut.Unlock()

	ch := make(chan any, 1)
	if len(b.subscribers[userId]) == 0 {
		b.BeginBroadcasting(userId)
	}
	b.subscribers[userId] = append(b.subscribers[userId], ch)
	return ch
}

func (b *brokerService) Unsubscribe(userId string, ch <-chan any) {
	log.Println("Unsubscribing a client from user:", userId)
	b.mut.Lock()
	defer b.mut.Unlock()

	channels := b.subscribers[userId]
	for i, subscriber := range channels {
		if subscriber == ch {
			b.subscribers[userId] = append(channels[:i], channels[i+1:]...)
			if len(b.subscribers[userId]) == 0 {
				b.stopChannels[userId] <- true
			}
			close(subscriber)
			break
		}
	}
}
