# gRPC and Kubernetes Exercise

## Requirements

A go gRPC server is provided to test your knowledge on several topics dear to us at Covr.

Perform the following and document your thought-process as you make choices to use one thing over another.

1. Using your tool of choice, provision a small kubernetes cluster in AWS.
2. Deploy an instance of the provided gRPC server to the cluster.
3. Expose the service publicly and secure all traffic with TLS.
    * Note that once this is done, you can use the `grpcc` utility to test without needing the insecure (`-i`) flag.
4. Ensure the service is highly available.

All configuration and scripts, along with documentation of the approach and a step by step deployment guide, should be
published in a public git repository for us to review. We should be able to use your solution to recreate the infrastructure required to host the service.

## Additional Information

### Service Description

The gRPC service implements two RPC endpoints which both split a sentence into its component words. The first endpoint is
called `SplitUnary` and takes a sentence and returns the words in a single response. The second endpoint is called
`SplitStream` and takes a sentence and returns the words as individual responses to the client. Both are hosted on port
`9001` which is exposed by the container.

### Image Hosting

The container image for the GRPC service is located [here](https://hub.docker.com/r/covrco/sre-grpc-exercise) on Docker
Hub.

### Building and Running Locally

If for some reason you would like to build the container image locally, you can use the command

```shell script
docker build -t covrco/sre-grpc-exercise .
```

to build the image and then run the image in a container with the command

```shell script
docker run -it -p 9001:9001 covrco/sre-grpc-exercise
```

### Testing the GRPC Server

To test the server, it is best to use a tool like [GRPCC](https://github.com/njpatel/grpcc), which uses the protocol
definitions to create a client implementation from the spec. When running from the same directory as this README, the
command to connect to the server would be:

```shell script
grpcc -i --proto cmd/protocol/sentence.proto --address <ADDRESS_GOES_HERE>:9001
```

Once connected, the following command can be used to test the unary endpoint.

```shell script
client.SplitUnary({ "sentence": "This, the best of all sentences." }, printReply)
```

The following command can be used to test the streaming endpoint. It will print both data messages and status messages to the
console as they arrive.

```shell script
client.SplitStream({ "sentence": "This, the best of all sentences." }).on('data', streamReply).on('status', streamReply)
```
