package hello

import (
  specs "github.com/opencontainers/runtime-spec/specs-go"
  "github.com/containerd/cgroups"
)

func Hello() {
  shares := uint64(100)
  control, err := cgroups.New(cgroups.V1, cgroups.StaticPath("/test"), &specs.LinuxResources{
    CPU: &specs.CPU{
      Shares: &shares,
    },
  })
  defer control.Delete()
}
