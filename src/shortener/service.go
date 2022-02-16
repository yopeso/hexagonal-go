package shortener

type RedirectService interface {
	Find(code string) (*LinkRedirect, error)
	Store(redirect *LinkRedirect) error
}
