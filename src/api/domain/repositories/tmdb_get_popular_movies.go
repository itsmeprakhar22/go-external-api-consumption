package repositories

type GetPopularMoviesRequest struct {
	ApiKey   string `json:"apiKey"`
	Language string `json:"language"`
	Page     int64  `json:"page"`
}

type GetPopularMoviesResponse struct {
	Page   int               `json:"page"`
	Result []GetPopularMovie `json:"results"`
}

type GetPopularMovie struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	BackdropPath string `json:"backdrop_path"`
	ReleaseDate  string `json:"release_date"`
}
