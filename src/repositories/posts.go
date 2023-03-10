package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (repository Posts) FindPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) FindPostById(postID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id
		where p.id = ?`, 
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) UpdatePost(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"update posts set title = ?, content = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID)
	err != nil {
		return err
	}
	
	return nil
}

func (repository Posts) DeletePost(postID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from posts where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) FindPostsByUserId(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?
		order by 1 desc`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) LikePost(postID uint64) error {
	statement, err := repository.db.Prepare(
		"update posts set likes = likes + 1 where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID)
	err != nil {
		return err
	}
	
	return nil
}

func (repository Posts) UnlikePost(postID uint64) error {
	statement, err := repository.db.Prepare(`
		update posts set likes = 
		case 
			when likes > 0 then  likes - 1
			else 0 
		end
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID)
	err != nil {
		return err
	}
	
	return nil
}