* Project Structure
```
project/
├── main.go
├── handlers/
│      ├── form.go
│      └── video.go
├── static/
       ├── index.html
       └── form.js
```
* Versions
```
go v1.16
gin/gonic v1.6.3
aws-sdk-go v1.55.5

```
* Build app
```bash
go build
```

* Run app
```bash
export AWS_ACCESS_KEY=your_access_key && export AWS_BUCKET_NAME=your_bucket_name && export AWS_REGION=your_region && export AWS_SECRET_KEY=your_secret_key && ./s3-video-upload
```

* Run server
```bash
http://localhost:8080/static/ 
```