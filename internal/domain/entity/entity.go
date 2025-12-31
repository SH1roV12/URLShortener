package entity



type Link struct{
	ID int
	OriginalURL string
	ShortURL string
}


func NewLink(short_url,original_url string)*Link{
	return &Link{
		OriginalURL: original_url,
		ShortURL: short_url,
	}
}