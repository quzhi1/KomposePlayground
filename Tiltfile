# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/hello-go ./'
if os.name == 'nt':
  compile_cmd = 'build.bat'

local_resource(
  'hello-go-compile',
  compile_cmd,
  deps=['main.go'],
)

docker_build_with_restart(
  'hello-go-image',
  '.',
  entrypoint=['/opt/app/bin/hello-go'],
  dockerfile='Dockerfile',
  only=[
    './bin',
  ],
  live_update=[
    sync('./bin', '/opt/app/bin'),
  ],
)

k8s_yaml(helm('./kompose-playground'))
