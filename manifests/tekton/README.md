# Tekton CICD

> `TODO` CD 使用 KubeVela Task

- Tekton 环境搭建参考 [hb-chen/tekton-practice](https://github.com/hb-chen/tekton-practice)
  - 需要完成 `Install` `Task` 和 `Config` 三个步骤
- microt-starter 提供了两种方式构建镜像 `buildpacks` 和 `kaniko`
  - `TriggerTemplate` 默认使用的 `buildpacks`
    - Trigger 的 `EventListener` 配置在 [hb-chen/tekton-practice/example/trigger.yml](https://github.com/hb-chen/tekton-practice/blob/main/example/trigger.yml)