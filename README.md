# torque-cli

<img alt="Logo" align="right" src="https://i.imgur.com/lfT6T9E.png" width="20%" />

The nuts and bolts of web services in Go.

## Quick Overview

```sh
torque new HospitalRegistration
cd HospitalRegistration
torque watch
```

The above commands will get you up and running with a brand new Torque app using the Torque CLI. The `watch` command ensures that while you're making changes to your project, your tests are ran every time a file changes.

Happy building!

## Developing the CLI

- Install [modd](https://github.com/cortesi/modd)
- Run `modd` in the project directory
  - This will watch the files in the project and:
    - run the tests for torque
    - build the torque binary
    - install the `torque` binary on your system in `~/bin`

## Notes

- https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/
