topics:
	curl -d 'init topic go-nsq-1' 'http://127.0.0.1:4151/pub?topic=go-nsq-1' & \
	curl -d 'init topic go-nsq-2' 'http://127.0.0.1:4151/pub?topic=go-nsq-2' & \
	curl -d 'init topic go-nsq-3' 'http://127.0.0.1:4151/pub?topic=go-nsq-3' & \
	curl -d 'init topic go-nsq-4' 'http://127.0.0.1:4151/pub?topic=go-nsq-4' & \
	curl -d 'init topic go-nsq-5' 'http://127.0.0.1:4151/pub?topic=go-nsq-5'
run-to-file:
	nsq_to_file --topic=go-nsq-1 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161 & \
	nsq_to_file --topic=go-nsq-2 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161 & \
	nsq_to_file --topic=go-nsq-3 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161 & \
	nsq_to_file --topic=go-nsq-4 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161 & \
	nsq_to_file --topic=go-nsq-5 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
run-consumer:
	go run consumer/main.go
run-producer:
	go run producer/main.go