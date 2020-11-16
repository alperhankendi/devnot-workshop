package movies

type Movie struct {
	ImdbTitleID       string `json:"imdb_title_id" bson:"_id"`
	Actors            string `json:"actors"`
	Country           string `json:"country"`
	DatePublished     string `json:"date_published"`
	Description       string `json:"description"`
	Director          string `json:"director"`
	Duration          int64  `json:"duration"`
	Genre             string `json:"genre"`
	OriginalTitle     string `json:"original_title"`
	ProductionCompany string `json:"production_company"`
	Title             string `json:"title"`
	Writer            string `json:"writer"`
	Year              int64  `json:"year"`
	Votes             int64  `json:"votes"`
}
