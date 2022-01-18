react:
	npm run --prefix frontend build
	cp -R frontend/build/* backend/spa 
	
go:
	cd backend
	go build 