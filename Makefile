.PHONY: up
up:
	# docker network create --subnet=172.88.0.0/16 lynet
	docker-compose up -d
.PHONY: down
down:
	docker-compose down

.PHONY: genssl
genssl:
	openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj /CN=localhost
