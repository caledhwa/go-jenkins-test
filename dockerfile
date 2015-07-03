FROM tianon/true
EXPOSE 1337

COPY certs/certs /etc/ssl/certs/
# other stuff
ADD main /

CMD ["/main"]
