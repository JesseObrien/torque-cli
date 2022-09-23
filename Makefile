build:
	go build -o dist/torque cmd/cli/main.go

install:
	make build
	rm -rf ~/bin/torque
	cp dist/torque ~/bin/torque
	chmod +x ~/bin/torque

install-watch:
	reflex -R 'dist' -- make install

test:
	go test ./... -count 1

test-watch:
	reflex -R 'dist' -- make test

scaffold:
	make install
	cd ~/tmp && rm -rf TestApp && torque new TestApp --mod-name github.com/jesseobrien/testapp && tree -a TestApp

scaffold-watch:
	reflex -R 'dist' -- make scaffold

