help: ## This help dialog
help h:
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-30s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#==========================#
#== CONTAINER INTERACTION ==#
#==========================#

environment:
environment:
	@echo "======================="
	@echo "===== Docker Init ====="
	@echo "======================="
	docker compose -f ${DOCKER_COMPOSE_FILE} up --build -d

#==========================#
#== CONTAINER INTERACTION ==#
#==========================#

shell-db: ## Enter to database console
shell-db:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

run:
run:
	go run main.go
