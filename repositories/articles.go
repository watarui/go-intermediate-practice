package repositories

import (
	"database/sql"

	"github.com/watarui/go-intermediate-practice/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
  INSERT INTO articles (title, contents, username, nice, created_at) values
  (?, ?, ?, 0, now());
  `

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
  SELECT article_id, title, contents, username, nice
  FROM articles
  LIMIT ? OFFSET ?;
  `

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum); err != nil {
			return nil, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
  SELECT *
  FROM articles
  WHERE article_id = ?;
  `

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
  SELECT nice
  FROM articles
  WHERE article_id = ?;
  `
	const sqlUpdateNice = `
  UPDATE articles
  SET nice = ?
  WHERE article_id = ?;
  `

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
