all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o parikshitg .

prod:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o parikshitg .
	- docker image rm reg.urantiatech.com/me/parikshitg
	docker build -t reg.urantiatech.com/me/parikshitg .
	docker push reg.urantiatech.com/me/parikshitg

dev:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o parikshitg .
	- docker image rm localhost:5000/me/parikshitg
	docker build -t localhost:5000/me/parikshitg .
	docker push localhost:5000/me/parikshitg
