package auth

type AuthService interface {
    signUp(*SignUp) (string, error)
    signIn()
    signOut()
}

type AuthServiceImpl struct {
	repo AuthRepo
}

var _ AuthService = (*AuthServiceImpl)(nil)

func (svc *AuthServiceImpl) signUp(params *SignUp) (string, error) {
    svc.repo.Create()
    return "ok", nil
}

func (svc *AuthServiceImpl) signIn() {
    svc.repo.GetByUsername()
}

func (svc *AuthServiceImpl) signOut() {

}
