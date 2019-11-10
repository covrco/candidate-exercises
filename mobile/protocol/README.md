## Dashboard Proto

## Building and running mock server

```shell
docker build -t mobile-mock-server .
docker run --rm -d mobile-mock-server
```

# Developing mocks

```shell
# Now you can interactively develop mocks, and stop and re-run for instant changes
docker run --rm -it -v $PWD:/src -p 4770:4770 -p 4771:4771 --entrypoint sh mobile-mock-server
```