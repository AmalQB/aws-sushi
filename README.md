# Sushi: A simple image storage as microservice

### Usage
To build sushi image with tag "sushi"
```
docker build . -t sushi
```
To run sushi deamon 
```
docker run -d -p 9090:9090 --name sushi-instance --rm sushi
```

### Image specification

| Density                   | Thumbnail | Ad view |
|---------------------------|-----------|---------|
| xxlarge (xxhdpi): 960x960 |           | 960     |
| xlarge (xhdpi): 640x960   | 320       | 640     |
| large (hdpi): 480x800     | 240       | 480     |
| medium (mdpi): 320x480    | 160       | 320     |

### API

#### PUT / HTTP/1.1
```
Content-Length: 65534
[file content]
    
HTTP/1.1 200 OK
[FIDs]
{
  "status": "OK",
  "data": {
    "image": [
      {
         "field": "xxlarge_view",
         "value": "0001-webp-xlarge-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-960-1034"
      },
      {
         "field": "xlarge_view",
         "value": "0001-webp-xlarge-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-640-871"
      },
      {
         "field": "xlarge_list",
         "value": "0001-webp-xlarge-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-320-435"
      },
      {
         "field": "large_view",
         "value": "0001-webp-large-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-480-653"
      },
      {
         "field": "large_list",
         "value": "0001-webp-large-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-240-327"
      },
      {
         "field": "medium_view",
         "value": "0001-webp-medium-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-320-435"
      },
      {
         "field": "medium_list",
         "value": "0001-webp-medium-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-160-218"
      }
    ]
  }
}
```

#### TEST

POST data using the Content-Type multipart/form-data
```	
curl -F image=@gopher.png http://localhost:9090
```	

#### OUTPUT
```
{
  "status": "OK",
  "data": {
    "image": [
      {
         "field": "xlarge_view",
         "value": "0001-jpeg-xlarge-view-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-640-871"
      },
      {
         "field": "xlarge_list",
         "value": "0001-jpeg-xlarge-list-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-320-435"
      },
      {
         "field": "large_view",
         "value": "0001-jpeg-large-view-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-480-653"
      },
      {
         "field": "large_list",
         "value": "0001-jpeg-large-list-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-240-327"
      },
      {
        "field": "medium_view",
        "value": "0001-jpeg-medium-view-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-320-435"
      },
      {
        "field": "medium_list",
        "value": "0001-jpeg-medium-list-738046cbd7a5a01b4a59f1e2c9b9c1a8851791da-EAEAEA-160-218"
      }
    ]
  }
}
```
