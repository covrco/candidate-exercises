'use strict'

const faker = require('faker')
const _ = require('lodash')
const fs = require('fs')

const banks = [
  "Bank of America",
  "US Bank",
  "Wells Fargo",
  "City Bank",
  "Peoples Bank"
]

const failedVerificationStates = [1, 2]
const successVerificationStates = [3]
const verificationState = new Map()
verificationState.set(1, 'FAILED')
verificationState.set(2, 'WAITING_ON_RETRY')
verificationState.set(3, 'SUCCEEDED')

const paymentHistoryCodesWithLinkedAccount = [2, 3]
const paymentHistoryMap = new Map()
paymentHistoryMap.set(1, 'NO_MISSED_PAYMENTS')
paymentHistoryMap.set(2, 'MISSED_PAYMENT_LAST_MONTH')
paymentHistoryMap.set(3, 'ON_PARTIAL_PAYMENT_PLAN')

const random = faker.random

function getLinkedInstitution () {
  if (random.number() % 2 === 0) {
    return [
      { name: random.arrayElement(banks), account: faker.finance.accountName(), accountNumber: faker.finance.account(), linkedDate: faker.date.recent() }
    ]
  } else {
    return []
  }
}

function getManagedClaims () {
  if (random.number() % 2 === 0) {
    return {
      vip: true
    }
  } else {
    return {}
  }
}

function randomCustomer () {
  const dateOne = faker.date.past()
  const dateTwo = faker.date.past()

  let dateFirstAttempt
  let dateSecondAttempt
  if (dateOne < dateTwo) {
    dateFirstAttempt = dateOne
    dateSecondAttempt = dateTwo
  } else {
    dateFirstAttempt = dateTwo
    dateSecondAttempt = dateOne
  }

  const institutions = getLinkedInstitution()
  const customer = {
    id: faker.random.uuid(),
    name:  {
      first: faker.name.firstName(),
      last: faker.name.lastName()
    },
    managedClaims: getManagedClaims(),
    verficationFlowHistory: [
      { state: verificationState.get(random.arrayElement(failedVerificationStates)), date: dateFirstAttempt },
      { state: verificationState.get(random.arrayElement(successVerificationStates)), date: dateSecondAttempt }
    ],
    linkedFinancialInstitutions: institutions
  }

  if (institutions.length) {
    customer.nextPaymentDate = faker.date.future()
    customer.customClaims = {
      paymentHistory: paymentHistoryMap.get(random.arrayElement(paymentHistoryCodesWithLinkedAccount))
    }
  }
  else{
    customer.customClaims = {
      paymentHistory: paymentHistoryMap.get(1)
    }
  }

  return customer
}

const customers = _.times(1000, randomCustomer)
fs.writeFileSync('./customers.json', JSON.stringify(customers, null, 2))
