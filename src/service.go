package src

type service struct{}

type Service interface {
	Login(req *loginRequest) (*loginRequest, error)
}

func NewService() Service {
	return &service{}
}
func (s *service) Login(req *loginRequest) (*loginRequest, error) {

}
