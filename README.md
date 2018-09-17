# chneau
Le website

# :(){ :|:&};:

# todo

Test this thing:  
```bash
docker run -d --restart=always --hostname mail.chneau.xyz --name mail -p 25:25 -p 143:143 -p 587:587 -p 993:993 -v /containers/mail/mail:/var/mail -v /containers/mail/state:/var/mail-state -v /containers/mail/config/:/tmp/docker-mailserver/ -e ENABLE_SPAMASSASSIN=0 -e ENABLE_CLAMAV=1 -e ENABLE_FAIL2BAN=1 -e ENABLE_POSTGREY=1 -e ONE_DIR=1 -e DMS_DEBUG=0 --cap-add NET_ADMIN --cap-add SYS_PTRACE tvial/docker-mailserver:latest
```
