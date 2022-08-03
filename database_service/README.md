# Database service

---

Database service represents a minimal service responsible for storing scheduled releases.
It offers two simple operations.

---

## Configuration

To successfully run the application it will expect to have the following
environment variables set. They are used to establish a successful connection to the database.

### Environment variables

`DATABASE_HOST` = URL to the mysql database, example `jdbc:mysql://127.0.0.1:3306/releases`

`DATABASE_USERNAME` = Username to use to log in to the database, example `root`

`DATABASE_PASSWORD` = Password to use to log in to the database, example `my-secret-pw`

### Port

The default port of the application will be `8080`, but you can remap it to what you like
using Docker.

---

## Operations

### Schedule a release
HTTP POST `/releases`

Body:

```json
{
  "releaseTitle": "Upgraded version",
  "releaseVersion": "v0.0.1",
  "releaseDate": "2022-08-03T16:00:06.00"
}
```
#### releaseTitle
Represents a unique title of the release.

#### releaseVersion
Represents a unique semantic version of the release. You can follow any convention, no
constraints applied. You can even say version:IWANNA_NAMME_IT_LIKE_THIS. Though,
wouldn't recommend it. Other engineers will look at you strangely.

#### releaseDate
Represents a release date in the future. It will be used as a deadline!

---
### Get all releases
HTTP GET `/releases`

Response body:
```json
[
  {
    "id": 5,
    "releaseTitle": "Upgraded version",
    "releaseVersion": "v0.0.1",
    "releaseDate": "2022-08-03T16:00:06.00"
  }
]
```

#### id
Represents a unique id of the release.

#### releaseTitle
Represents a unique title of the release.

#### releaseVersion
Represents a unique semantic version of the release. You can follow any convention, no
constraints applied. You can even say version:IWANNA_NAMME_IT_LIKE_THIS. Though,
wouldn't recommend it. Other engineers will look at you strangely.

#### releaseDate
Represents a release date in the future. It will be used as a deadline!
