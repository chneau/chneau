.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run install

run: install

uninstall:
	docker rm --force mail


install: uninstall
	# -p 25:25 mailservers
	# -p 143:143 TLS IMAP client
	# -p 587:587 TLS client email submission
	# -p 110:110 POP3 client

	# -p 465:465 SSL client email submission
	# -p 993:993 SSL IMAP client
	# -p 995:995 SSL POP3 client
	docker run -d --restart=always --hostname mail.chneau.xyz --name mail -p 25:25 -p 143:143 -p 587:587 -v /containers/mail/mail:/var/mail -v /containers/mail/state:/var/mail-state -v /containers/mail/config/:/tmp/docker-mailserver/ -v /etc/localtime:/etc/localtime:ro -e ENABLE_SPAMASSASSIN=0 -e ENABLE_CLAMAV=0 -e ENABLE_FAIL2BAN=0 -e ENABLE_POSTGREY=0 -e ONE_DIR=1 -e DMS_DEBUG=0 --cap-add NET_ADMIN --cap-add SYS_PTRACE tvial/docker-mailserver:latest

adduser:
	curl -o setup.sh https://raw.githubusercontent.com/tomav/docker-mailserver/master/setup.sh; chmod a+x ./setup.sh
	echo "USAGE: ./setup.sh email add user@domain pwd"


