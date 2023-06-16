# JiberSoz - Pastebin alternative, written in Go

**JiberSoz** is a pasting service similar to [Pastebin](https://github.com/lordelph/pastebin) that you can host yourself. The API allows you to share snippets of text with others. You paste your text and get a short URL that you can share with anyone. This is the point. But that is not all! The service is implemented according to microservice architecture. Used **Redis** for caching and **Amazon S3** for storing snippets.
[Hash Generator](https://github.com/jumagaliev1/hash-generator) mircoservice with connect gRPC.
![Untitled-2023-06-16-1927](https://github.com/jumagaliev1/jiberSoz/assets/71185943/285a9bc9-6a17-43c5-bb6f-0202f7803145)

## Setup and Installation
To set up the project, follow these steps:
1. Clone the project repository: `git clone https://github.com/jumagaliev1/jiberSoz.git`
2. Clone hash generator microservice: `git clone https://github.com/jumagaliev1/hash-generator.git`
3. Install Docker, docker compose engine.
4. `docker compose up --build app` (if have errors try `docker compose up`) for both project.

## API Endpoints
The project includes the following REST API endpoints that can be accessed for various operations:
- `POST /api/create`: Creates new Snippet.
```
  {
    "message": "some text",
    "day": 1
  }
```
- `GET /api/text/hAuaQnf`: Get Snippet.
```
  {
    "message": "some text"
  }
```
