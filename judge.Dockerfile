ARG PROJECT_NAME

FROM debian:trixie-slim AS judge_deps
RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates git gcc libc6-dev make libcap-dev libsystemd-dev pkgconf \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
RUN git clone --depth 1 --branch v2.2.1 https://github.com/ioi/isolate.git
WORKDIR /app/isolate
RUN make isolate isolate-cg-keeper && make install DESTDIR=/out

FROM ${PROJECT_NAME}-base AS njudge-bin

FROM debian:trixie-slim
ENV TZ=Europe/Budapest

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates tzdata pandoc gcc g++ libc6-dev \
        openjdk-25-jdk-headless mono-mcs fpc cython3 gccgo python3-dev \
        libcap2 libsystemd0 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=judge_deps /out/ /

# PyPy 8.0.1 (Python 3.12) from the official image. PyPy, Julia and Nim live
# under /usr/local, as isolate boxes only expose /usr, /bin and /lib*.
COPY --from=pypy:3.12-8.0.1-trixie /opt/pypy /usr/local/pypy
COPY --from=julia:1.13.4 /usr/local/julia/ /usr/local/julia
COPY --from=nimlang/nim:1.6.18 /opt/nim /usr/local/nim
RUN ln -s /usr/local/pypy/bin/pypy3 /usr/bin/pypy3 && \
    ln -s /usr/local/julia/bin/julia /usr/local/bin/julia && \
    ln -s /usr/local/nim/bin/nim /usr/bin/nim

# Ensure that the standard library is pre-compiled to bytecode (.pyc) so that
# each run in isolate does not spend time building it, only to be discarded
# afterwards as isolate does not have write access to the cache.
# Also smoke test this by importing some common modules.
RUN pypy3 -m compileall -q -j0 /usr/local/pypy/lib && \
    pypy3 -c "import decimal, functools, itertools, math, random, sys"

WORKDIR /app
COPY --from=njudge-bin /app/njudge ./njudge
COPY configs/docker/judge.yaml ./judge.yaml
CMD ["./njudge", "judge"]
