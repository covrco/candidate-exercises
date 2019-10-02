# Internal Customer Records Portal

Build a management dashboard app for our support staff to be able to visualize
a user's lifecycle with our mobile fintech app. Without going into too many details,
the app currently allows user to do the following:

1. Verify their identity
2. Optionally link a financial institution to take advantage of our services.

Internally, we use a series of claims to authorize access to a variety of features. Some are auto-generated claims, and others are managed, as you will see below in the JSON schema.

## Requirements

The dashboard you would build would allow for the following:

1. A grid search that enables filtering by any combination of following values:

   - Name
   - Number of linked bank accounts
   - VIP status
   - Payment history

2. Each user entity should be clickable to navigate to a details view. This view should have the following features:

  - We should be able to edit a user's claims.
  - We should be able to review a user's verification flow history in a visually pleasing manner. For now, the random data has each record first having failed, followed by a success. The implementation should be able to handle other cases as well
  - We should be able to review any linked institutions for the user along with info about upcoming payments.
  - The custom claim for payment history is also displayed for all users.

## Sample Data

  For simplicity we've provide both a sample of the data below and a script to generate a large set of test records to use in the application.

### Sample

```json
{
  "name": {
    "first": "Bulah",
    "last": "Padberg"
  },
  "managedClaims": {
    "vip": true
  },
  "verficationFlowHistory": [
    {
      "state": "FAILED",
      "date": "2019-02-13T07:42:29.318Z"
    },
    {
      "state": "SUCCEEDED",
      "date": "2019-04-30T08:28:05.272Z"
    }
  ],
  "linkedFinancialInstitutions": [
    {
      "name": "Bank of America",
      "account": "Investment Account",
      "accountNumber": "88268710",
      "linkedDate": "2019-10-01T03:18:52.959Z"
    }
  ],
  "nextPaymentDate": "2020-02-27T23:39:22.046Z",
  "customClaims": {
    "paymentHistory": "NO_MISSED_PAYMENTS"
  }
}
```

### Generating Data

```shell
cd customers/data-gen && npm install
node index.js #outputs customers_100.json
```