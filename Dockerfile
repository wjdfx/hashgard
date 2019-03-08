FROM golang:1.12 as build
ENV GOBIN /go/bin
ADD . /go/src/github.com/hashgard/hashgard/
RUN cd /go/src/github.com/hashgard/hashgard/ && make get_tools && make get_vendor_deps && make install

FROM ubuntu:16.04
EXPOSE 26656
EXPOSE 26657
RUN apt update && \
    apt install -y iputils-ping net-tools vim curl wget && \
    apt clean && apt autoclean
COPY --from=build /go/bin/hashgard /usr/local/bin/
COPY --from=build /go/bin/hashgardcli /usr/local/bin/
COPY --from=build /go/bin/hashgardkeyutil /usr/local/bin/
COPY --from=build /go/bin/hashgardlcd /usr/local/bin/
COPY --from=build /go/bin/hashgardreplay /usr/local/bin/
ADD docker/start.sh /
CMD /start.sh
