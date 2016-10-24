Personal Go Blog
================

This is supposed to become a simple blog written in Go and Angular 2 hosted by AppEngine.

The frontend and backend can be tested separately or together.

## Testing & Development
### Testing everyting
To testing everyting like it will run on AppEngine first build the frontend manually:
```
cd frontend
ng build
```
Then it's ready to be tested:
```
goapp serve dispatch.yaml frontend/frontend.yaml backend/backend.yaml
```

### Testing the backend
To only test the backend, you can run:
```
goapp serve backend/backend.yaml
```

### Testing the frontend
To develop and test the frontend you probably wanna use the angular tools. You should therefore do:
```
cd frontend
ng serve
```
Then if you want the backend running alongside it, simply test the backend in a new window.

## Deployment
First build the frontend:
```
cd frontend
ng build -prod
```
Then deploy the two modules:
```
goapp deploy frontend/frontend.yaml backend/backend.yaml
```
Then update the routing/dispatch:
```
appcfg.py update_dispatch
```
