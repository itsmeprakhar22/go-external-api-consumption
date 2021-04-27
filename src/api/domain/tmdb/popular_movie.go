package tmdb

/*
{
	"adult": false,
	"backdrop_path": "/9yBVqNruk6Ykrwc32qrK2TIE5xw.jpg",
	"genre_ids": [
	14,
	28,
	12,
	878,
	53
	],
	"id": 460465,
	"original_language": "en",
	"original_title": "Mortal Kombat",
	"overview": "Washed-up MMA fighter Cole Young, unaware of his heritage, and hunted by Emperor Shang Tsung's best warrior, Sub-Zero, seeks out and trains with Earth's greatest champions as he prepares to stand against the enemies of Outworld in a high stakes battle for the universe.",
	"popularity": 6861.583,
	"poster_path": "/6Wdl9N6dL0Hi0T1qJLWSz6gMLbd.jpg",
	"release_date": "2021-04-07",
	"title": "Mortal Kombat",
	"video": false,
	"vote_average": 8,
	"vote_count": 1286
},
*/

type PopularMovieRequest struct {
	ApiKey   string `json:"apiKey"`
	Language string `json:"language"`
	Page     int64  `json:"page"`
}

type PopularMovieResponse struct {
	Page   int            `json:"page"`
	Result []PopularMovie `json:"results"`
}
type PopularMovie struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	BackdropPath string `json:"backdrop_path"`
	ReleaseDate  string `json:"release_date"`
}
