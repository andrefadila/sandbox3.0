package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (wh *WebHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, errId := strconv.Atoi(id)
	if errId != nil {
		fmt.Fprintf(w, "Validation error: %s", errId.Error())
		return
	}

	dept, err := wh.rs.GetDepartment(idInt)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Dept Name = %s", dept.Name)
}
