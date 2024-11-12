### s3-files-upload: simple Amazon S3 files uploader using gin 

* Versions:
```
go v1.16
gin/gonic v1.6.3
aws-sdk-go v1.55.5
```

* Build app:
```bash
go build 
```

* Execute app:
```bash
export AWS_ACCESS_KEY=your_access_key && export AWS_BUCKET_NAME=your_bucket_name && export AWS_REGION=your_region && export AWS_SECRET_KEY=your_secret_key && ./s3-files-upload
```