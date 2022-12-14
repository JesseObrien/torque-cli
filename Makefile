build:
	go build -o dist/torque cmd/cli/main.go

install:
	make build
	rm -rf ~/bin/torque
	cp dist/torque ~/bin/torque
	chmod +x ~/bin/torque

test:
	go test ./... -count 1

scaffold:
	make install
	cd ~/tmp && rm -rf TorqueApp && torque new TorqueApp --mod-name github.com/jesseobrien/testapp && tree -a TestApp

scaffold-watch:
	reflex -R 'dist' -- make scaffold

