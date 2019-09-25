NAMESPACE=eddygt
IMAGE=apphostname
TAG=1.0

build: clean build-binary build-image

build-binary:
	GOOS=linux go build -ldflags="-s -w" -o bin/apphostname

build-image:
	docker build -t ${NAMESPACE}/${IMAGE}:${TAG} .

clean:
	rm bin/${IMAGE} || true
	docker rmi ${NAMESPACE}/${IMAGE}:${TAG} || true

publish:
	docker push ${NAMESPACE}/${IMAGE}:${TAG}

run:
	docker run --rm -p 8080:8080 ${NAMESPACE}/${IMAGE}:${TAG}
print:
	echo ${NAMESPACE}/${IMAGE}:${TAG}
