package handler

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"time"

	"github.com/samirghasemi/golang-for-simra/lottery/pkg/lottery"
	"github.com/samirghasemi/golang-for-simra/lottery/pkg/redis"
)

type lotteryPayload struct {
	UUID string `json:"uuid"`
}

var prizes = []lottery.Prize{
	{Name: "A", Weight: 0.1},
	{Name: "B", Weight: 0.3},
	{Name: "C", Weight: 0.2},
	{Name: "D", Weight: 0.15},
	{Name: "E", Weight: 0.25},
}

var client = redis.NewClient()

func LotteryHandler(w http.ResponseWriter, r *http.Request) {

	var payload lotteryPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	count, err := client.GetUserCount(ctx, payload.UUID)
	// fmt.Println("err: ", err)
	if err != nil {
		http.Error(w, "Failed to get user participation count", http.StatusInternalServerError)
		return
	} else if count >= 3 {
		http.Error(w, "User already participated 3 times today", http.StatusTooManyRequests)
		return
	}

	if err := client.IncrementUserCount(ctx, payload.UUID); err != nil {
		http.Error(w, "Failed to increment user participation count", http.StatusInternalServerError)
		return
	}

	time.Sleep(5 * time.Second)

	prize := lottery.Draw(prizes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"prize": prize})
}
