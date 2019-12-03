package dto

type Link struct {
    Url string `json:"url"`
}

type GetLinkResult struct {
    Val string
    Err error
}

type CreateLinkResult struct {
    ShortLink string `json:"shortLink"`
    Err       error  `json:"-"`
}
