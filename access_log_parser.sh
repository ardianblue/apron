log_url=https://raw.githubusercontent.com/linuxacademy/content-elastic-log-samples/master/access.log
curl -s $log_url | awk '{print $1}' | sort -u  > unique_ip_addresses.log

