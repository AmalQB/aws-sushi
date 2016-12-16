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

### Image specification ([configurable](https://github.com/microservices-today/aws-sushi/blob/master/sushi.conf))

| Density                   | Thumbnail | Ad view |
|---------------------------|-----------|---------|
| xxlarge (xxhdpi): 960x960 |           | 960     |
| xlarge (xhdpi): 640x960   | 320       | 640     |
| large (hdpi): 480x800     | 240       | 480     |
| medium (mdpi): 320x480    | 160       | 320     |

### API

#### PUT / HTTP/1.1
```JSON
Content-Length: 65534
[file content]
    
HTTP/1.1 200 OK
[FIDs]
{
  "status": "OK",
  "data": {
    "image":[
      {"field":"xxhdpi_view","value":"0001-jpeg-xxhdpi-view-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-960-960"},
      {"field":"xlarge_view","value":"0001-jpeg-xlarge-view-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-640-640"},
      {"field":"xlarge_list","value":"0001-jpeg-xlarge-list-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-320-320"},
      {"field":"large_view","value":"0001-jpeg-large-view-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-480-480"},
      {"field":"large_list","value":"0001-jpeg-large-list-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-240-240"},
      {"field":"medium_view","value":"0001-jpeg-medium-view-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-320-320"},
      {"field":"medium_list","value":"0001-jpeg-medium-list-f9a5fdfe5cb7de1d3fe2d77baebdc84efcad4058-49515C-160-160"}
    ]
  }
}
```

#### TEST

POST data using the Content-Type multipart/form-data
```	
curl -F image=@gopher.png http://localhost:9090
```	
