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
	return s.repository.CheckIfUserExists(email)
}
