CREATE TABLE dogs (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   breed TEXT NOT NULL DEFAULT 'outbred'
);

-- name: ListDoggos :many

SELECT *
FROM dogs;

-- name: GetDoggo :one 
SELECT *
FROM dogs
WHERE id = $1;

-- name: AddDoggo :execresult
INSERT INTO dogs (name, breed)
VALUES($1, $2);

-- name: UpdateDoggoInfo :execresult
UPDATE dogs
SET name = $1, breed = $2
WHERE id = $3;

-- name: DeleteDoggoInfo :execresult
DELETE FROM dogs
WHERE id = $1;
