package subscriber

type EmailDto struct {
	Email string `form:"email" binding:"required,email"`
}

type SubscriberService struct {
	repository SubscriberRepositoryInterface
}

func NewService(repository SubscriberRepositoryInterface) SubscriberService {
	return SubscriberService{repository}
}

func (s *SubscriberService) Exists(email string) (bool, error) {
	exists, err := s.repository.CheckIfUserExists(email)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
