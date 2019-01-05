package apierrors

type Unauthorized struct{}

func NewUnauthorized() *Unauthorized {
	return &Unauthorized{}
}

func (u *Unauthorized) Error() string {
	return ""
}
