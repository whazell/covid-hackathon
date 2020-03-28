build:
	go build -o covid ./cmd

run-docker-db:
	docker run -d --name mysql -p 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=pass \
		-e MYSQL_DATABASE=covid \
		-e MYSQL_USER=covid \
		-e MYSQL_PASSWORD=password \
		mysql:8.0

