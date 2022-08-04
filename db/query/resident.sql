-- name: CreateResident :one
INSERT INTO residents (
    nic,
    fullName,
    rAddress,
    clearance
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetResident :one
SELECT * FROM residents
WHERE nic = $1 LIMIT 1;
