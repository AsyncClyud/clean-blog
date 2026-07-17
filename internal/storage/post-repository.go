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

	rows, err := pr.db.Query("SELECT Id, Title FROM posts")
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		article := models.Article{}
		err := rows.Scan(&article.Id, &article.Title)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}

	result, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		log.Println(err)
	}

	return string(result)
}

func (pr *PostRepository) GetArticleById(Id int) (byid_article string) {
	var article models.Article

	rows, err := pr.db.Query("SELECT Title, Content, Created_At, Author FROM posts WHERE Id = $1", Id)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&article.Title, &article.Content, &article.Created_at, &article.Author)
		if err != nil {
			log.Println(err)
		}
	}
	result, err := json.MarshalIndent(article, "", " ")
	if err != nil {
		log.Println(err)
	}

	return string(result)
}

func (pr *PostRepository) InsertArticle(article models.Article, author int) sql.Result {
	result, err := pr.db.Exec(
		"INSERT INTO Posts(Title, Content, Created_at, Author) VALUES ($1, $2, $3, $4)", article.Title, article.Content, time.Now(), author)
	if err != nil {
		log.Println("Insert article query error:", err)
	}
	log.Printf("Inserted new article with title %v; Article author: %v", article.Title, author)

	return result
}

func (pr *PostRepository) UpdateArticle(article models.Article) sql.Result {
	result, err := pr.db.Exec(
		"UPDATE Posts SET Title = $1, Content = $2 WHERE Id = $3", article.Title, article.Content, article.Id)
	if err != nil {
		log.Println("Update article query error:", err)
	}
	log.Printf("Updated article with title: %v", article.Title)

	return result
}

func (pr *PostRepository) DeleteArticle(article models.Article) sql.Result {
	result, err := pr.db.Exec("DELETE FROM Posts WHERE Id = $1", article.Id)
	if err != nil {
		log.Println("Delete article query error:", err)
	}
	log.Printf("Deleted article with id: %v", article.Id)

	return result
}

func (pr *PostRepository) GetArticleCommentsById(id int) (comment string) {
	var comments []models.Comment

	rows, err := pr.db.Query("SELECT Comment_content, Created_at, Author FROM Comments WHERE Post_id = $1", id)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.Comment_content, &comment.Created_at, &comment.Author)
		if err != nil {
			log.Println(err)
		}
		comments = append(comments, comment)
	}
	result, err := json.MarshalIndent(comments, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	return string(result)

}

func (pr *PostRepository) InsertComment(comment models.Comment, author int) sql.Result {
	result, err := pr.db.Exec(
		"INSERT INTO Comments(Comment_content, Created_at, Post_id, Author) VALUES($1, $2, $3, $4)", comment.Comment_content, time.Now(), comment.Post_id, author)
	if err != nil {
		log.Println("Insert comment query error:", err)
	}
	log.Printf("Inserted comment in post with id: %v", comment.Post_id)

	return result
}
