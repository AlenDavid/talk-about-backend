load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml('deployments/plants.yaml')
k8s_yaml('deployments/database.yaml')

k8s_resource('database', port_forwards=5432)

compile_args = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

local_resource(
  'codegen',
  'go run ./cmd/gen/gen.go',
  deps=['cmd', 'internal', 'pkg'],
  ignore=['internal/persistence'],
  resource_deps=['database'],
  labels="compilation"
)

local_resource(
  'plants-compile',
  compile_args + 'go build -o ./build/plants ./cmd/plants',
  deps=['internal/persistence'],
  labels="compilation",
)

docker_build_with_restart(
  'plants-app',
  '.',
  entrypoint=['/app/build/plants'],
  dockerfile='deployments/Dockerfile',
  only=[
    './build',
  ],
  live_update=[
    sync('./build', '/app/build'),
  ],
)

k8s_resource('plants-app-deployment',
  port_forwards=8080,
  labels="plants",
  resource_deps=['plants-compile', 'database', 'broker'],
)
