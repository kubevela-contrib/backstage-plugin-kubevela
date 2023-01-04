build:
	 docker build . -t wonderflow/backstage-plugin-kubevela

push:
	docker push wonderflow/backstage-plugin-kubevela

build-push: build push
	echo "done"