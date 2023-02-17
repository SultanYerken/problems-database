package usecase

import "git.01.alem.school/Sultanye/problems-database/internal/entity"

type (
	ToDoProblem interface {
		GetAllProblems() ([]entity.Problem, error)
		GetProblemById(id int) (entity.Problem, error)
		CreateProblem(problem entity.Problem, topics []string) (int, error)
		EditProblem(problem entity.Problem, topics []string) (int, error)
		DeleteProblem(id int) error
	}
	TodoProblemRepository interface {
		GetAllProblems() ([]entity.Problem, error)
		GetProblemById(id int) (entity.Problem, error)
		CreateProblem(problem entity.Problem) (int, error)
		EditProblem(id int, problem entity.Problem) (int, error)
		DeleteProblem(id int) error
		GetTopicIdByName(topic string) (int, error)
		WriteiInProblemTopicTable(problemId, topicId int) error
		DeleteFromProblemTopicTable(problemId int) error
		// WriteTopics() error
	}
)

type TodoProblemUsecase struct {
	repo TodoProblemRepository
}

func NewTodoProblemUseCase(repo TodoProblemRepository) *TodoProblemUsecase {
	return &TodoProblemUsecase{repo: repo}
}

func (s *TodoProblemUsecase) GetAllProblems() ([]entity.Problem, error) {
	// if err := s.repo.WriteTopics(); err != nil {
	// 	return nil, err
	// }

	return s.repo.GetAllProblems()
}

func (s *TodoProblemUsecase) GetProblemById(id int) (entity.Problem, error) {
	return s.repo.GetProblemById(id)
}

func (s *TodoProblemUsecase) CreateProblem(problem entity.Problem, topics []string) (int, error) {
	var topicIds []int
	for i := range topics {
		id, err := s.repo.GetTopicIdByName(topics[i])
		if err != nil {
			return 0, err
		}
		topicIds = append(topicIds, id)
	}

	problemId, err := s.repo.CreateProblem(problem)
	if err != nil {
		return 0, err
	}

	for i := range topicIds {
		err := s.repo.WriteiInProblemTopicTable(problemId, topicIds[i])
		if err != nil {
			return 0, err
		}
	}
	return problemId, nil
}

func (s *TodoProblemUsecase) EditProblem(problem entity.Problem, topics []string) (int, error) {
	var topicIds []int
	for i := range topics {
		id, err := s.repo.GetTopicIdByName(topics[i])
		if err != nil {
			return 0, err
		}
		topicIds = append(topicIds, id)
	}

	id, err := s.repo.EditProblem(problem.Id, problem)
	if err != nil {
		return 0, err
	}

	err = s.repo.DeleteFromProblemTopicTable(problem.Id)
	if err != nil {
		return 0, err
	}

	for i := range topicIds {
		err := s.repo.WriteiInProblemTopicTable(problem.Id, topicIds[i])
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (s *TodoProblemUsecase) DeleteProblem(id int) error {
	return s.repo.DeleteProblem(id)
}
