-- name: GetBotByToken :one
SELECT *
FROM bots
WHERE token = $1;

-- name: AddBot :exec
