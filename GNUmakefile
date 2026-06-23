NAME := NAP

all:
	@echo "Targets: deploy delete"

deploy:
	gcloud functions deploy $(NAME) --gen2 --runtime go126 --trigger-http --allow-unauthenticated

delete:
	gcloud functions delete $(NAME)
