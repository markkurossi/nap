NAME := NAP

all:
	@echo "Targets: deploy delete"

deploy:
	gcloud functions deploy $(NAME) --gen2 --runtime go121 --trigger-http

delete:
	gcloud functions delete $(NAME)
