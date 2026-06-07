package multi

import (
	"github.com/cox722/go-fullstack-cox/internal/errs"
	log "github.com/cox722/go-fullstack-cox/package/log"
)

func New(handlers ...log.Handler) *Handler {
	return &Handler{handlers}
}

// Handler logs by level
type Handler struct {
	handlers []log.Handler
}

func (f *Handler) Log(entry *log.Entry) (err error) {
	for _, handler := range f.handlers {
		err = errs.Join(err, handler.Log(entry))
	}
	return err
}
