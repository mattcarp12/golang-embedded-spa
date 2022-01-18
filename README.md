# golang-embedded-spa

## 1. create react app

```
npx create-react-app frontend
```

## 2. create go app
```
mkdir backend && cd backend
go mod init backend
```

## 3. Build frontend
```
cd frontend
npm run build
cp -R build/* ../backend/spa
```

