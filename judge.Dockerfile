ARG PROJECT_NAME

FROM debian:bookworm as judge_deps
RUN apt-get update
RUN apt-get install -y wget gcc g++ git build-essential
RUN apt-get install -y libcap-dev systemd libsystemd-dev pkgconf
WORKDIR /app
RUN git clone --depth 1 --branch v2.0 https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate isolate-cg-keeper

FROM ${PROJECT_NAME}-base

RUN apt-get update && apt-get install -y openjdk-17-jdk mono-mcs fpc cython3 golang pandoc gccgo pypy3 python3-dev g++ gcc libsystemd-dev
COPY --from=judge_deps /app/isolate /app/isolate
WORKDIR /app/isolate
RUN make install

COPY --from=julia:1.10.1 /usr/local/julia/ /usr/local/julia
COPY --from=nimlang/nim:1.6.18 /usr/bin/nim /usr/bin/nim
RUN ln -s /usr/local/julia/bin/julia /usr/local/bin/julia

WORKDIR /app
COPY configs/docker/judge.yaml ./judge.yaml
CMD ["./njudge", "judge"]
