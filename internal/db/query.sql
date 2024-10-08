-- name: GetBotById :one
SELECT *
FROM bots
WHERE id = $1;

-- name: GetBotByToken :one
SELECT *
FROM bots
WHERE token = $1;

-- name: AddBot :one
INSERT INTO bots (
    token,
    studio_id,
    first_name,
    username
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
ON CONFLICT (token) DO NOTHING
RETURNING *;