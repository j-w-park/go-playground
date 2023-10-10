FROM golang as builder

WORKDIR /app

# 빌드 시 VCS 관련 meta 정보를 바이너리에 포함하는데 빌드할 때 VCS 관련 meta 정보를 바이너리에 포함하는데 컨테이너 환경에서
# VCS에 safe directory 설정이 안되어 있는 경우 정보를 가져오지 못해 빌드가 실패하는 경우가 있음.
# 이를 해결하기 위해 git을 설치하고 safe directory 설정을 추가
# 다른 방법으로는 빌드 플래그에 `-buildvcs=false`를 추가하는 방법이 있음.
RUN apt update
RUN apt install -y git
RUN git config --global --add safe.directory $PWD

# ENTRYPOINT [ "/bin/bash", "-c", "sleep infinity" ]
COPY . .
RUN make build

FROM scratch as runner
WORKDIR /app
COPY --from=builder /app/bin/go-playground /app
EXPOSE 8080
ENTRYPOINT [ "/app/go-playground" ]
