# TODO: dynamically fetch this from cargo file
ARG SP1_VERSION="v4.1.0"
FROM ghcr.io/succinctlabs/sp1:${SP1_VERSION} as build

RUN apt update && apt install -y just protobuf-compiler

WORKDIR /solidity-ibc-eureka/
COPY . .

RUN cargo build --release --locked

FROM ubuntu:24.04
RUN apt-get update && apt-get install -y wget curl nano

WORKDIR /usr/local/bin
COPY --from=build /solidity-ibc-eureka/target/release/relayer /usr/local/bin/relayer
COPY --from=build /solidity-ibc-eureka/target/release/operator /usr/local/bin/operator

EXPOSE 3000

ENTRYPOINT ["/bin/bash", "-c"]
