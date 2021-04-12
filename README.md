# OcrService
This repository contains the source code and docker container for an OCR REST API using gosseract, a GoLang implementation of Google's tesseract OCR library. 

## Compiling
To build and run the container, in your terminal run ```docker-compose up``` from the project directory. Please reference https://docs.docker.com/get-started/ to get started with docker.

## API Usage

1. The following request processes an OCR request on any filetype synchronously. The response time may be slow depending on the size of the image:

```
curl -XPOST "http://localhost:5000/image-sync" -d '{"image_data": "data:image/<imagetype>;base64,<your base64 encoded image data>}'
```

2. The following request processes and OCR request asynchronously, returning instead a taskid with which you can retrieve the OCR results in a separate GET request.
```
curl -XPOST "http://localhost:5000/image" -d '{"image_data": "data:image/<imagetype>;base64,<your base64 encoded image data>}'
```
3. The following request will retrieve the OCR results for the provided task id. In the event the OCR task has not finished, null will be returned.
```
curl -XGET "http://localhost:5000/image" -d '{"taskid": <your ocr task id>}'
```

## Database Administration

This webservice uses pgadmin which contains a web-based postgresql database management tool so you can view the database workload and contents. Navigate to ```localhost:16543``` and use the crendentials found in .env in this repository to log in. A new server instance must be created to view the database. Be sure to use ```db``` as the host name in the connection section, as the container it is running in docker and thus has a non-static ip address.