# Software Engineering School 4.0

## API Documentation

### Main end-points
**1. Subscribe to Email Notifications**

Method: **POST**

`URL: /api/subscribe`

Description: This endpoint allows users to subscribe to email notifications by providing their email addresses.

Request Body Example
```
{
  "email": "user@gmail.com"
}
```
Response Example
```
{
    "message": "Email added successfully"
}

```

If Email already exists
```
{
    "error": "Email already exists"
}
```

Other problems:
```
{
    "error": "Email request body is not correct."
}
```

**2. Get Exchange Rate**

Method: **GET**

`URL: /api/rate`

Description: This endpoint returns the current exchange rate USD - UAH.

Response Example:
```
39.55
```

### Util end-points
**1. Ping-Pong**

Method: **GET**

`URL: /api/ping`

Description: This endpoint is used to check if the server is running. It responds with a "pong" message.

Response Example:

```
{
  "message": "pong"
}

```

**2. Force Email notification to all users**

Method: **POST**

`URL: /api/notify`

Description: This endpoint sends email notifications to all users immediately, bypassing the scheduled interval.

Response Example:
```
{
  "message": "Emails sent successfully"
}
```

## Properties description

All variables are stored in .env file

**1. Default port is 8080, but you can specify another one.**

`PORT=3000`

**2. Database properties**
```
DB_URL="host=localhost user=root password=root dbname=case port=5432 sslmode=disable"
DB_FULL_URL="postgres://root:root@localhost:5432/case?sslmode=disable"
```
I am using Postgresql. Docker command to run database separately:

`docker run --name my-postgres-db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=case -p 5432:5432 -d postgres`

**3. Exchange rate API**

I am using API for the exchange rate USD - UAH from [PrivatBank API](https://api.privatbank.ua/#p24/exchange "PrivatBank API")

`https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5`

**4. SMTP settings**

I am using a Google account and Gmail SMTP service to send Emails to users.
Password is application password
```
GOOGLE_USERNAME="se.school.case.2024.notification@gmail.com"
GOOGLE_PASSWORD=VERY_SECRET_PASSWORD
```

**5. Email time of notification**

You can set the time when the app will send emails to all users.

`EMAIL_SEND_TIME="17:43"`


### Example of Email notification
[![](https://raw.githubusercontent.com/Gurmigou/se-school-case-2024/main/show/data/email_notification_example.png)](https://raw.githubusercontent.com/Gurmigou/se-school-case-2024/main/show/data/email_notification_example.png)


### Other points of implementation
1. I decided to store the current exchange rate in a database instead of in-memory storage. This approach allows for future scalability as the application may provide more exchange rates. The database now includes additional fields: currency_from and currency_to (defaulting to USD and UAH, respectively).

2. The current exchange rate is cached in the database and is not fetched from the PrivatBank API with every request. A lazy check mechanism is implemented: if the last fetch from the API was more than 1 hour ago, a new fetch is made and the updated rate is stored in the database.

