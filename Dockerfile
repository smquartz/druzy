FROM alpine:3.2
ADD druzy-srv /druzy-srv
ENTRYPOINT [ "/druzy-srv" ]
