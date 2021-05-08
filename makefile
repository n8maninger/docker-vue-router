static:
	go build -o ./dist/server -a -trimpath -ldflags="-s -w -extldflags='-static'" -tags='netgo timetzdata' .

docker:
	docker build -t n8maninger/vue-router:beta .
	docker push n8maninger/vue-router:beta