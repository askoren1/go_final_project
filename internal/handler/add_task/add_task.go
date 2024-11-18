package add_task

import (
	"net/http"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
