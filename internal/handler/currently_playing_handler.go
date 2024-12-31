package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bojurgess/bard/internal/database"
	"github.com/bojurgess/bard/internal/model"
	"github.com/bojurgess/bard/internal/service"
	"log"
	"net/http"
)

func CurrentlyPlaying(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	if !database.UserService.Exists(userId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	ch := service.BrokerService.Subscribe(userId)

	var lastMessage *model.SpotifyCurrentlyPlaying

	for {
		select {
		case event := <-ch:
			if event == "stop" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			msg := event.(*model.SpotifyCurrentlyPlaying)

			var e string
			var buf []byte
			var err error

			if lastMessage == nil {
				e = "new_track"
				buf, err = json.Marshal(msg)
				if err != nil {
					log.Println(err)
				}
			} else if lastMessage.Item.Name != msg.Item.Name {
				e = "new_track"
				buf, err = json.Marshal(msg)
				if err != nil {
					log.Println(err)
				}
			} else {
				e = "track_update"
				buf, err = json.Marshal(&model.SpotifyCurrentlyPlayingTrackUpdate{
					ProgressMs: msg.ProgressMs,
					Timestamp:  msg.Timestamp,
				})
				if err != nil {
					log.Println(err)
				}
			}

			_, err = fmt.Fprintf(w, "event: %s\n\n data: %s\n\n", e, string(buf))
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			lastMessage = msg
			flusher.Flush()
		case <-r.Context().Done():
			service.BrokerService.Unsubscribe(userId, ch)
			log.Println("Client disconnected")
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
