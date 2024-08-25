IMAGE_NAME = plotly-data-generator

all: build run

build:
	@docker build -t $(IMAGE_NAME) .

run:
	@docker run --rm $(IMAGE_NAME)



clean:
	@docker rmi $(IMAGE_NAME)
