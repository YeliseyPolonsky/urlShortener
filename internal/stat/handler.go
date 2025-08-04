package stat

import (
	"fmt"
	"go-advance/configs"
	"go-advance/pkg/middlware"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandler struct {
	*StatRepository
	*configs.Config
}

type StatHandlerDep struct {
	*StatRepository
	*configs.Config
}

func NewStatHandler(router *http.ServeMux, dep StatHandlerDep) {
	handler := &StatHandler{
		StatRepository: dep.StatRepository,
		Config:         dep.Config,
	}

	router.Handle("GET /stat", middlware.IsAuth(handler.GetStat(), dep.Config))
}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
		}

		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
		}

		fmt.Println("\n", "from:", from, "to:", to, "by:", by)
	}
}
