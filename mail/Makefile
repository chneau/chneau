.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run uninstall install logs adduser

run: install

uninstall:
	docker rm --force mail

extractcerts:
	cd /containers/traefik
	wget https://raw.githubusercontent.com/containous/traefik/master/contrib/scripts/dumpcerts.sh && chmod +x dumpcerts.sh
	./dumpcerts.sh acme.json ./certs/

install: uninstall
	# -p 25:25 mailservers
	# -p 143:143 TLS IMAP client
	# -p 587:587 TLS client email submission
	# -p 110:110 POP3 client

	# -p 465:465 SSL client email submission
	# -p 993:993 SSL IMAP client
	# -p 995:995 SSL POP3 client
	docker run -d --restart=always --hostname mail.chneau.xyz --name mail -p 25:25 -p 143:143 -p 587:587 -v /containers/mail/mail:/var/mail -v /containers/mail/state:/var/mail-state -v /containers/mail/config/:/tmp/docker-mailserver/ -v /etc/localtime:/etc/localtime:ro -e ONE_DIR=1 --cap-add NET_ADMIN --cap-add SYS_PTRACE tvial/docker-mailserver:latest

installwithcerts: uninstall
	docker run -d --restart=always --hostname mail.chneau.xyz --name mail -p 25:25 -p 993:993 -p 465:465 -v /containers/mail/mail:/var/mail -v /containers/mail/state:/var/mail-state -v /containers/mail/config/:/tmp/docker-mailserver/ -v /containers/traefik/certs:/certs -v /etc/localtime:/etc/localtime:ro -e SSL_TYPE=manual -e SSL_CERT_PATH=/certs/certs/\*.chno.xyz.crt -e SSL_KEY_PATH=/certs/private/\*.chno.xyz.key -e ONE_DIR=1 --cap-add NET_ADMIN --cap-add SYS_PTRACE tvial/docker-mailserver:latest

logs:
	docker logs -f mail

adduser:
	curl -o setup.sh https://raw.githubusercontent.com/tomav/docker-mailserver/master/setup.sh; chmod a+x ./setup.sh
	echo "USAGE: ./setup.sh email add user@domain pwd"


