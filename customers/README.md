# Internal Customer Records Portal

Build a management dashboard app to provide our support staff the capability to visualize
a user's lifecycle within our fintech product. For this exercise we will focus on 2 aspects of the product:

1. A user can verify their identity
2. A user can optionally link a financial institution to take advantage of our services.

We use a series of claims to authorize access to a variety of features. Some are auto-generated claims, and others are managed, as you will see below in the JSON schema.

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

## Sample Wireframe

We've provided a sample wireframe that represents what the searchable datagrid _could_ look like.  This does not account for all the possible views, so you have the freedom to implement the design however you'd like.

![portal-mock](/Users/arri/_data/covr/candidate-exercises/customers/portal-mock.png)

## Sample Data

  For simplicity we've provide both a sample of the data below and a script to generate a large set of test records to use in the application.

### Sample

```json
{
  "id": "3da4867d-15cc-4482-b5a7-f880ab0cff6d",
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

Running the provided data generation script will produce 1,000 customer records by default.

```shell
cd customers/data-gen && npm install
node index.js #outputs customers.json
```

## Submission

To submit your work, please send an email to your recruiter with a link to a public GitHub repository  used to host your code.
