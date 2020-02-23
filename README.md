# Overlap
Overlap is a very simple service that detects intersection between rectangles and saves them.
Current storage for saving is memory but it's simple to implement your custom storage.

### API
# Show all rectangles

Get list of all available rectangles in db

**URL** : `/`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

```json
[
    {
        "id": 1,
        "rect": {
            "x1": -1,
            "y1": -1,
            "x2": 4,
            "y2": 3
        },
        "created_at": "2020-02-22T20:04:44.2683678-08:00"
    },
    {
        "id": 2,
        "rect": {
            "x1": 2,
            "y1": 18,
            "x2": 7,
            "y2": 22
        },
        "created_at": "2020-02-22T20:04:44.2683678-08:00"
    }
]
```

# Show first unique rectangle

Get first unique rectangle from db

**URL** : `/unique`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

```json
{
    "id": 1,
    "rect": {
        "x1": -1,
        "y1": -1,
        "x2": 4,
        "y2": 3
    },
    "created_at": "2020-02-22T20:04:44.2683678-08:00"
}
```

# Detect intersection

This api accepts a `main` rectangle and `input`. `input` is array of rectangles which should be saved in case they have any intersection with `main`.

**URL** : `/`

**Method** : `POST`

## Success Response

**Code** : `201 Created`

**Request body example**

```json
{
	"main": {"x1": 0, "y1": 0, "x2": 10, "y2": 20},
	"input": [
		{"x1": -1, "y1": -1, "x2": 4, "y2": 3},
		{"x1": 12, "y1": 18, "x2": 17, "y2": 22},
		{"x1": 2, "y1": 18, "x2": 7, "y2": 22}
	]
}
```

## NOTE
There isn't any payload in response of this api(for sake of simplicity xD)

### Installation

Overlap requires [golang](https://golang.org/) v1.13.8+ to run.

Install the dependencies and start the server.

```sh
$ go get github.com/golang/dep
$ go get github.com/sonemaro/overlap
$ cd $GOPATH/src/github.com/sonemaro/overlap
$ dep ensure
$ go build ./cmd/overlap/
$ ./overlap
```

You can also run `overlap` by `-h` param to see cli options.

```sh
$ ./overlap -h
```

### Docker
Overlap's docker file is insanely small and secure.
It uses docker multi-stage feature thus makes him strong and small.
```sh:wq
docker build -t sonemaro/overlap:latest .
```
Running the container based on previous image
```sh
docker run --rm -p 8000:41279 sonemaro/overlap:latest
```

License
----

MIT
