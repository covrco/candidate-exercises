# frontend exercise

Build a management dashboard app for our support staff to be able to visualize
a user's lifecycle with our mobile fintech app. Without going into too many details,
the app currently allows user to do the following:

1. Prove who they are for our KYC needs
2. Optionally link a financial institution, with which they hope to interface with our app

Internally, we manage the metadata of the user with claims. Some are auto-generated claims, and others are
managed, as you will see below in the JSON schema.

The dashboard you would build would allow for the following:

1. A grid search which displays the name, and claims for the user.
2. The user entity is a link that directs one to a different details view per user. This view has following features:

  1. It allows one to modify managed claims.
  2. It shows the verification flow history in a visually pleasing manner. For now, the random data has each record first having failed, followed by a success.
  3. It shows any linked institutions for the user, with an upcoming payment date.
  4. The custom claim for payment history is also displayed for all users.
3. Back on the grid page, also implement filtering by following criteria. Note that criteria should allowed to be compounded using an AND operator.

  1. Number of linked bank accounts
  2. Vip customers
  3. Payment history claim

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

