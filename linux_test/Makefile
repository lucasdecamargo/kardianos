
all: sysv systemd upstart openrc clean

# compile `go test` binary statically
test:
	@CGO_ENABLED=0 go test -installsuffix netgo -a -c ..

clean:
	-rm kardianos.test

sysv: test
	@echo sysv
	@cp kardianos.test sysv/
	@docker build -q --tag="kardianos.test.sysv" sysv
	@-docker run kardianos.test.sysv
	@-docker rm $(shell docker ps -l -q)
	@-docker rmi -f kardianos.test.sysv
	@-rm sysv/kardianos.test

systemd: test
	@echo systemd
	@cp kardianos.test systemd/
	@docker build -q --tag="kardianos.test.systemd" systemd
	@-docker run --privileged -v /sys/fs/cgroup:/sys/fs/cgroup:ro kardianos.test.systemd
	@-docker rm $(shell docker ps -l -q)
	@-docker rmi -f kardianos.test.systemd
	@-rm systemd/kardianos.test

upstart: test
	@echo upstart
	@cp kardianos.test upstart/
	@docker build -q --tag="kardianos.test.upstart" upstart
	@-docker run kardianos.test.upstart
	@-docker rm $(shell docker ps -l -q)
	@-docker rmi -f kardianos.test.upstart
	@-rm upstart/kardianos.test

openrc: test
	@echo openrc
	@cp kardianos.test openrc/
	@docker build -q --tag="kardianos.test.openrc" openrc
	@-docker run kardianos.test.openrc
	@-docker rm $(shell docker ps -l -q)
	@-docker rmi -f kardianos.test.openrc
	@-rm openrc/kardianos.test