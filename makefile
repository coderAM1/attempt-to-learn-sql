.PHONY: startPostgres stopPostgres

startPostgres:
	docker compose up --detach

stopPostgres:
	docker compose down