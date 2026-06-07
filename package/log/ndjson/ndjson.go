package json

import log "github.com/cox722/go-fullstack-cox/package/log"

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) Log(entry *log.Entry) error {
	return nil
}
