package storage

import (
	"blog/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) GetAllArticles() (all_articles string) {
	articles := []models.Article{}

	rows, err := pr.db.Query("SELECT * FROM posts")
	if err != nil {
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		article := models.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Created_at, &article.Author)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}

	result, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	return string(result)

}

func (pr *PostRepository) GetArticleById(Id int) (article string) {
	articles := []models.Article{}

	rows, err := pr.db.Query("SELECT * FROM posts WHERE Id = $1", Id)
	if err != nil {
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		article := models.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Created_at, &article.Author)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}
	result, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	return string(result)

}

func (pr *PostRepository) InsertArticle(article models.Article, Author_Id int) sql.Result {
	result, err := pr.db.Exec(
		"INSERT INTO Posts(Title, Content, Created_at, Author) VALUES ($1, $2, $3, $4)", article.Title, article.Content, time.Now(), Author_Id)
	if err != nil {
		log.Fatalln("Insert query error:", err)
	}

	return result
}

func (pr *PostRepository) UpdateArticle(article models.Article) sql.Result {
	result, err := pr.db.Exec(
		"UPDATE Posts SET Title = $1, Content = $2 WHERE Id = $3", article.Title, article.Content, article.Id)
	if err != nil {
		log.Fatalln("Update query error:", err)
	}

	return result
}

func (pr *PostRepository) DeleteArticle(article models.Article) sql.Result {
	log.Println(article)
	result, err := pr.db.Exec("DELETE FROM Posts WHERE Id = $1", article.Id)
	if err != nil {
		log.Fatalln("Delete query error:", err)
	}
	return result
}
