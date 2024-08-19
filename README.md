# youtube-go
Go Client for Google's Private InnerTube API. Works with **YouTube**, **YouTube Music**, **YouTube Kids**, **YouTube Studio** and more!

## About
This library handles low-level interactions with the underlying InnerTube API used by each of the YouTube services.

### Latest Release
```console
go get github.com/wslyyy/youtube-go
```

## Usage
```go
>>> import "github.com/@wslyyy/youtube-go"
>>>
>>> // Construct a client
>>> clientName := "Web"
>>> clientVersion := "2.20230728.00.00"
>>> apiKey := ""
>>> userAgent := ""
>>> referer := ""
>>> auto := true
>>> client, err := NewInnerTube("WEB", "2.20230728.00.00", "", "", "", nil, true)
>>>
>>> // Get some data!
>>> PARAMS_TYPE_VIDEO := "EgIQAQ%3D%3D"
>>> //PARAMS_TYPE_CHANNEL := "EgIQAg%3D%3D"
>>> //PARAMS_TYPE_PLAYLIST := "EgIQAw%3D%3D"
>>> //PARAMS_TYPE_FILM := "EgIQBA%3D%3D"
>>> query := "foo fighters"
>>> data, err := client.Search(&query, &PARAMS_TYPE_VIDEO, nil)
>>>
>>> // The core endpoints are implemented, so the above is equivalent to:
>>> channelId := "xxxxxxxxxxxx"
>>> channelData, err := client.Browse(&channelId, nil, nil)
>>> // if use params
>>> videosParams := "xxxxxxxxxxxx"
>>> videosData, err := client.Browse(&channelId, &videosParams, nil)
>>> // if use continuation
>>> continuation := "xxxxxxxxxxxx"
>>> continuedVideosData, err := client.Browse(nil, nil, &continuation)
```

## Comparison with the [YouTube Data API](https://developers.google.com/youtube/v3/)
The InnerTube API provides access to data you can't get from the Data API, however it comes at somewhat of a cost *(explained below)*.
|                                       | This Library | YouTube Data API |
| ------------------------------------- | ------------ | ---------------- |
| Google account required               | No           | Yes              |
| Request limit                         | No           | Yes              |
| Clean data                            | No           | Yes              |

The InnerTube API is used by a variety of YouTube services and is not designed for consumption by users. Therefore, the data returned by the InnerTube API will need to be parsed and sanitised to extract data of interest.

## Endpoints
Currently only the following core, unauthenticated endpoints are implemented:
|                                | YouTube | YouTubeMusic | YouTubeKids | YouTubeStudio |
| ------------------------------ | ------- | ------------ | ----------- | ------------- |
| config                         | &check; | &check;      | &check;     | &check;       |
| browse                         | &check; | &check;      | &check;     | &check;       |
| player                         | &check; | &check;      | &check;     | &check;       |
| next                           | &check; | &check;      | &check;     |               |
| search                         | &check; | &check;      | &check;     |               |
| guide                          | &check; | &check;      |             |               |
| get_transcript                 | &check; |              |             |               |
| music/get_search_suggestions   |         | &check;      |             |               |
| music/get_queue                |         | &check;      |             |               |

## Authentication
The InnerTube API uses OAuth2, however this has not yet been implemented, therefore this library currently only provides unauthenticated API access.

## Thanks to the open source [innertube](https://github.com/tombulled/innertube/)