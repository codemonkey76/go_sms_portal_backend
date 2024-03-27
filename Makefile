# Makefile for seting up the Go web server systemd service and running database migrations
# Author: Shane Poppleton
# Created: 2024-03-27

# Variables
SERVICE_NAME := sms_backend_service
RELATIVE_EXECUTABLE_PATH := ./bin/${SERVICE_NAME}
ABSOLUTE_EXECUTABLE_PATH := $(shell realpath .)/bin/${SERVICE_NAME}
WORKING_DIRECTORY := $(shell realpath .)/bin
USERNAME := $(USER)
GROUP_NAME := ${USERNAME}

include .env
export

DATABASE_URL="postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

.PHONY: migrate-up migrate-down migrate-reset install uninstall
echo:
	echo ${USERNAME}
migrate-up:
	migrate -path db/migration -database ${DATABASE_URL} -verbose up

migrate-down:
	migrate -path db/migration -database ${DATABASE_URL} -verbose down

migrate-reset:
	docker-compose down
	docker-compose up -d
	sleep 2
	migrate -path db/migration -database ${DATABASE_URL} -verbose up

install:
	mkdir -p ./bin
	go build -o ${RELATIVE_EXECUTABLE_PATH}
	sudo cp ${SERVICE_NAME}.service /etc/systemd/system/${SERVICE_NAME}.service
	sudo sed -i 's|{{EXECUTABLE_PATH}}|${ABSOLUTE_EXECUTABLE_PATH}|g' /etc/systemd/system/${SERVICE_NAME}.service
	sudo sed -i 's|{{WORKING_DIRECTORY}}|${WORKING_DIRECTORY}|g' /etc/systemd/system/${SERVICE_NAME}.service
	sudo sed -i 's|{{USERNAME}}|${USERNAME}|g' /etc/systemd/system/${SERVICE_NAME}.service
	sudo sed -i 's|{{GROUP_NAME}}|${GROUP_NAME}|g' /etc/systemd/system/${SERVICE_NAME}.service

	sudo systemctl daemon-reload
	sudo systemctl enable --now ${SERVICE_NAME}

uninstall:
	sudo systemctl stop ${SERVICE_NAME}
	sudo systemctl disable ${SERVICE_NAME}

	sudo rm /etc/systemd/system/${SERVICE_NAME}.service

	sudo deluser --remove-home ${USERNAME}
	sudo delgroup ${GROUP_NAME}

	sudo systemctl daemon-reload
