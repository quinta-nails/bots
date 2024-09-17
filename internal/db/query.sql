-- name: GetBotByToken :one
SELECT *
FROM bots
WHERE token = $1;

-- name: AddBot :one
INSERT INTO bots (
    token,
    first_name,
    username
)
VALUES (
    $1,
    $2,
    $3
)
ON CONFLICT (token) DO NOTHING
RETURNING *;