package campaign

type Service interface {
	FindCampaign(UserID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaign, err := s.repository.FindByUserID(UserID)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}
	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
