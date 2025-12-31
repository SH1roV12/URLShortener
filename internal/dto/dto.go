package dto


type NewLink struct{
	OriginalURL string `json:"url" binding:"required"`
}

type GetOriginalURL struct{
	ShortURL string `json:"short_url" binding:"required"`
}

