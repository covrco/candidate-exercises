# Mobile Coding Challenge

Using the provided service based around gRPC, build out a mobile application that allows a user to see a collection of bank accounts and selecting a bank shows the transactions for that account. There are three aspects of this challenge we are interested in.

1. Use of the provided docker container to spin up the gRPC service locally.
2. Given the provided protoc file, generate and integrate the gRPC client.
3. Build a mobile application based on the below requirements.

### Note
You're required to use [Sketch](https://www.sketch.com/), which is available for free as a trial.

## Requirements

The mobile application you would build should allow for the following:

1. Use of the provided docker container to spin up the local gRPC service which uses the mocks also provided.
2. Given the protocol file, generate the appropriate client code for both [swift-protobuf](https://github.com/apple/swift-protobuf) and [grpc-swift](https://github.com/grpc/grpc-swift) and implement. There is a lot of documentation in the linked repositories on how to get started.
3. Build the following user experience based on the provided designs.
    * Implement a basic landing screen to be presented while loading the data.
    * Implement the dashboard screen that shows the user's name, today's date and three credit cards using the provided data.
    * When tapping on one of these credit cards, the user should be transitioned to a detail screen that presents the user with the transactions that have been fetched.
4. For bonus points, the following could be completed.
    * Using the provided Lottie animation, show this animation while loading the data on the landing page.
    * Implement a custom transition between the dashboard and transaction details screen.
    * Provide unit tests for any functionality that you deem worthy.


## Submission

To submit your work, please send an email to your recruiter with a link to a public GitHub repository used to host your code. Any technical documentation is greatly appreciated.