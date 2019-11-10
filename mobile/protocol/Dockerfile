FROM charandas/gripmock:covr-optimizations

WORKDIR /src

RUN apk add --no-cache curl

COPY . /src

CMD [ \
  "sh", \
  "-c", \
  "/src/start-mock-server.sh" \
]