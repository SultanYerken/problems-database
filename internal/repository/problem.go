package repository

import (
	"git.01.alem.school/Sultanye/problems-database/internal/entity"
)

func (r *Repository) GetAllProblems() ([]entity.Problem, error) {
	var problems []entity.Problem
	query := `SELECT id, title, description, level, topics, samples, created_at, updated_at FROM problem`
	rows, err := r.db.Query(query)
	if err != nil {
		return problems, err
	}
	defer rows.Close()

	for rows.Next() {
		problem := entity.Problem{}
		err := rows.Scan(&problem.Id, &problem.Title, &problem.Description, &problem.Level, &problem.Topics, &problem.Samples, &problem.CreatedAt, &problem.UpdateAt)
		if err != nil {
			return problems, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func (r *Repository) GetProblemById(id int) (entity.Problem, error) {
	var problem entity.Problem
	query := `SELECT id, title, description, level, topics, samples, created_at, updated_at FROM problem WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&problem.Id, &problem.Title, &problem.Description, &problem.Level, &problem.Topics, &problem.Samples, &problem.CreatedAt, &problem.UpdateAt)
	if err != nil {
		return problem, err
	}
	return problem, nil
}

func (r *Repository) CreateProblem(problem entity.Problem) (int, error) {
	var id int
	query := `INSERT INTO problem (title, description, level, topics, samples, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	row := r.db.QueryRow(query, problem.Title, problem.Description, problem.Level, problem.Topics, problem.Samples, problem.CreatedAt, problem.UpdateAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) EditProblem(id int, problem entity.Problem) (int, error) {
	var problemId int
	query := `UPDATE problem SET title = $1, description = $2, level = $3, topics = $4, samples = $5, updated_at = $6 WHERE id = $7 RETURNING id`
	row := r.db.QueryRow(query, problem.Title, problem.Description, problem.Level, problem.Topics, problem.Samples, problem.UpdateAt, id)
	if err := row.Scan(&problemId); err != nil {
		return 0, err
	}
	return problemId, nil
}

func (r *Repository) DeleteProblem(id int) error {
	query := `DELETE FROM problem WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTopicIdByName(topic string) (int, error) {
	var topicId int
	query := `SELECT id FROM topic WHERE topic_name = $1`
	err := r.db.QueryRow(query, topic).Scan(&topicId)
	if err != nil {
		return 0, err
	}
	return topicId, nil
}

func (r *Repository) WriteiInProblemTopicTable(problemId, topicId int) error {
	query := `INSERT INTO problem_topic (problem_id, topic_id) VALUES ($1,$2)`
	_, err := r.db.Exec(query, problemId, topicId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromProblemTopicTable(problemId int) error {
	query := `DELETE FROM problem_topic WHERE problem_id = $1`
	_, err := r.db.Exec(query, problemId)
	if err != nil {
		return err
	}
	return nil
}

// func (r *Repository) WriteTopics() error {
// 	arrtopics := []string{"Array", "String", "Hash Table", "Dynamic Programming", "Math", "Sorting", "Tree", "Binary Tree", "Binary Search", "Database", "other"}

// 	for i := range arrtopics {
// 		query := `INSERT INTO topic (topic_name) VALUES ($1)`
// 		_, err := r.db.Exec(query, arrtopics[i])
// 		if err != nil {
// 			return nil
// 		}
// 		fmt.Println(arrtopics[i], "insert")
// 	}
// 	return nil
// }
