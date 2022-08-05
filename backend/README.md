# Backend

<h6 align="right">Wow.. what an original name..</h6>

---

The backend is used as a proxy to be able to store and get releases. But in addition to that
it has the task of authorizing the user to do so. If you provide the right header
value, you can access the release database, if not.. well..

---

## Configuration

To successfully run the application it will expect to have the following
environment variables set.

### Environment variables

`AUTH_TOKEN` = Value used to authorize access. It will be compared to the `Authorization: <token>` value.

`RELEASES_URL` = URL of the release service.

### Port

The default port of the application will be `8081`, but you can remap it to what you like
using Docker.

---

## Operations

Please check [database_service](../database_service)