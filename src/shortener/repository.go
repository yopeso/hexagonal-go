package shortener

type LinkRepository interface {
	Find(code string) (*LinkRedirect, error)
	Store(link *LinkRedirect) error
}
