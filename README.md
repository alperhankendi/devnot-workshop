README



### Todo

![Design](docs/api.png)

### service runner
```
$> ❯ ./devnot-workshop
Devnote workshop application

Usage:
  devnot-workshop [flags]
  devnot-workshop [command]

Available Commands:
  command     Command Service
  help              Help about any command
  query            Query Service

Flags:
  -c, --conn 	string     	database connection string
  -d, --dbname string   	database name (default "imdb")
  -h, --help            		help for devnot-workshop
  -p, --port 	string    	Service Port (default "5001")

Use "devnot-workshop [command] --help" for more information about a command.

```

movie data structure
```
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
```

### request sample

```
curl --location --request POST 'http://127.0.0.1:5001/api/v1/' \
--header 'Content-Type: application/json' \
--data-raw '{  
      "imdb_title_id": "tt0000009",
      "title": "Miss Jerry",
      "original_title": "Miss Jerry",
      "year": 1894,    
      "date_published": "1894-10-09",    
      "genre": "Romance",
      "duration": 45,    "country": "USA",
      "director": "Alexander Black",    "writer": "Alexander Black",
      "production_company": "Alexander Black Photoplays",
      "actors": "Blanche Bayliss, William Courtenay, Chauncey Depew",
      "description": "The adventures of a female reporter in the 1890s.",
      "votes": 154
    }'
```

```
curl --location --request GET 'http://127.0.0.1:5000/api/v1/tt0000009'
```