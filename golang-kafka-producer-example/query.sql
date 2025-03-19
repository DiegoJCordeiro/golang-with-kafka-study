-- name: QueryAllPosts :many
SELECT * FROM POSTS LIMIT ? OFFSET ?
