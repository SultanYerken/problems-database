package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.01.alem.school/Sultanye/problems-database/internal/entity"
	"github.com/gorilla/mux"
)

type Responses struct {
	AllProblems []entity.Problem `json:"all_problems"`
	Problem     entity.Problem   `json:"problems"`
	ProblemId   int              `json:"problem_id"`
}

func (h *Handler) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	allProblems, err := h.usecases.GetAllProblems()
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Problems:", allProblems)

	output := Responses{
		AllProblems: allProblems,
	}
	writeJSON(w, r, output)
}

func (h *Handler) GetProblemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	problem, err := h.usecases.GetProblemById(id)
	if err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Problem:", problem)

	output := Responses{
		Problem: problem,
	}
	writeJSON(w, r, output)
}

func (h *Handler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	level := r.FormValue("level")
	topicsarr := r.Form["topics"]
	topics := strings.Join(topicsarr, " ")
	samples := r.FormValue("samples")

	if !chekLevel(level) {
		NewErrorResponse(w, r, http.StatusBadRequest, "levels: hard, medium, or easy")
		return
	}

	problem := entity.Problem{
		Title:       title,
		Description: description,
		Level:       level,
		Topics:      topics,
		Samples:     samples,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Time{},
	}
	fmt.Println("update_at", problem.UpdateAt)

	id, err := h.usecases.CreateProblem(problem, topicsarr)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Problem #%d created\n", id)

	output := Responses{
		ProblemId: id,
	}
	writeJSON(w, r, output)
}

func (h *Handler) EditProblem(w http.ResponseWriter, r *http.Request) {
	problemId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.usecases.GetProblemById(problemId)
	if err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, "Problem not Exist")
		return
	}

	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	level := r.FormValue("level")
	topicsarr := r.Form["topics"]
	topics := strings.Join(topicsarr, " ")
	samples := r.FormValue("samples")

	if !chekLevel(level) {
		NewErrorResponse(w, r, http.StatusBadRequest, "levels: hard, medium, or easy")
		return
	}

	problem := entity.Problem{
		Id:          problemId,
		Title:       title,
		Description: description,
		Level:       level,
		Topics:      topics,
		Samples:     samples,
		UpdateAt:    time.Now(),
	}

	id, err := h.usecases.EditProblem(problem, topicsarr)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Problem #%d edited\n", id)

	output := Responses{
		ProblemId: problemId,
	}
	writeJSON(w, r, output)
}

func (h *Handler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.usecases.DeleteProblem(id)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Problem #%d deleted\n", id)

	output := Responses{
		ProblemId: id,
	}
	writeJSON(w, r, output)
}

func writeJSON(w http.ResponseWriter, r *http.Request, resp Responses) {
	data, err := json.Marshal(resp)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func chekLevel(level string) bool {
	if level == "easy" || level == "medium" || level == "hard" {
		return true
	}
	return false
}
