# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

# Records the current time, then kicks off a server update.
# Normally, you would let Tilt do deploys automatically, but this
# shows you how to set up a custom workflow that measures it.
# local_resource(
#     'deploy',
#     'python record-start-time.py',
# )

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/hello-go ./'
if os.name == 'nt':
  compile_cmd = 'build.bat'

local_resource(
  'hello-go-compile',
  compile_cmd,
  deps=['main.go'],
#   resource_deps=['deploy']
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
