Головачев Роман Романович, группа БПМИ216, выбран вариант с социальной сетью

hw2 curl:

`
curl -X POST 0.0.0.0:8080/register -d '{"username": "aboba3", "password": "12345"}' -H "Content-Type: application/json" -v
`

`
curl -X POST 0.0.0.0:8080/login -d '{"username": "aboba3", "password": "12345"}' -H "Content-Type: application/json" -v
`

`
curl -X PUT 0.0.0.0:8080/update -d '{"firstName": "kekus2", "telephone": "2281337"}' -H "Content-Type: application/json" -H "Cookie: <your cookie goes here>"  -v
`

Можно зайти внутрь контейнера и проверить что действительно я все обновил:

`docker exec -it auth_db bash`

`psql -U auth -W super_secret_pass -d auth_db` (уже внутри контейнера)

`select * from users;` (уже внутри бд)

hw3 curl:

`
curl -X POST 0.0.0.0:8080/create_post -d '{"text": "test_post"}' -H "Content-Type: application/json" -H "Cookie: <your cookie goes here>"  -v
`

`
curl -X GET 0.0.0.0:8080/get_post/<post_id of your post> -H "Cookie: <your cookie goes here>"
`

`
curl -X PUT 0.0.0.0:8080/update_post/<post_id of your post> -H "Cookie: <your cookie goes here>" -d '{"text": "test_post_updated"}' -H "Content-Type: application/json" -v
`

`
curl -X DELETE 0.0.0.0:8080/delete_post/<post_id of your post> -H "Cookie: <your cookie goes here>"
`

`
curl -X GET 0.0.0.0:8080/get_posts -H "Cookie: <your cookie goes here>" -H "From: 0" -H "Count: 100"
`

hw4:

`
curl -X POST 0.0.0.0:8080/register -d '{"username": "aboba52", "password": "12345"}' -H "Content-Type: application/json" -v
`

`
curl -X POST 0.0.0.0:8080/login -d '{"username": "aboba52", "password": "12345"}' -H "Content-Type: application/json" -v
`

`
curl -X POST 0.0.0.0:8080/create_post -d '{"text": "post from aboba52"}' -H "Content-Type: application/json" -H "Cookie: <your cookie goes here>"
`

`
curl -X POST 0.0.0.0:8080/like_post/<your post_id goes here> -H "Cookie: <your cookie goes here>" -v
`

view data in CH:

`
docker exec -it clickhouse clickhouse-client
`

`
SELECT * FROM likes;
`
