-- name: GetAuthor :one
SELECT * FROM category
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM category
ORDER BY id desc;

-- name: CreateAuthor :execresult
INSERT INTO category (
    name
) VALUES (
    ?
);

-- name: DeleteAuthor :exec
DELETE FROM category
WHERE id = ?;