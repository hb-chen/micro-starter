# 服务发现-k8s

使用Pod的labels、annotations实现微服务的注册与发现功能
- 添加两个label
    > micro.mu/type:service，标记Pod为Micro服务
    
    > micro.mu/selector-{服务名}:service，标记Pod为指定服务名的Micro服务
    
- 添加一个annotations
    > micro.mu/service-{服务名}:{registry.Service的JSON编码}

## 服务注册
为Pod添加labels、annotations
```go
func (c *kregistry) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	if len(s.Nodes) == 0 {
		return errors.New("you must register at least one node")
	}

	// TODO: grab podname from somewhere better than this.
	podName := os.Getenv("HOSTNAME")
	svcName := s.Name

	// encode micro service
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	svc := string(b)

	pod := &client.Pod{
		Metadata: &client.Meta{
			Labels: map[string]*string{
				labelTypeKey:                             &labelTypeValueService,
				svcSelectorPrefix + serviceName(svcName): &svcSelectorValue,
			},
			Annotations: map[string]*string{
				annotationServiceKeyPrefix + serviceName(svcName): &svc,
			},
		},
	}

	if _, err := c.client.UpdatePod(podName, pod); err != nil {
		return err
	}

	return nil

}
```

## 服务注销
删除key=micro.mu/selector-{服务名}的label和key=micro.mu/service-{服务名}的annotation
```go
func (c *kregistry) Deregister(s *registry.Service) error {
	if len(s.Nodes) == 0 {
		return errors.New("you must deregister at least one node")
	}

	// TODO: grab podname from somewhere better than this.
	podName := os.Getenv("HOSTNAME")
	svcName := s.Name

	pod := &client.Pod{
		Metadata: &client.Meta{
			Labels: map[string]*string{
				svcSelectorPrefix + serviceName(svcName): nil,
			},
			Annotations: map[string]*string{
				annotationServiceKeyPrefix + serviceName(svcName): nil,
			},
		},
	}

	if _, err := c.client.UpdatePod(podName, pod); err != nil {
		return err
	}

	return nil

}
```

## 获取服务
通过key=micro.mu/selector-{服务名}的label查询Pods，并根据svc.Version对服务节点进行分类合并
```go
func (c *kregistry) GetService(name string) ([]*registry.Service, error) {
	pods, err := c.client.ListPods(map[string]string{
		svcSelectorPrefix + serviceName(name): svcSelectorValue,
	})
	if err != nil {
		return nil, err
	}

	if len(pods.Items) == 0 {
		return nil, registry.ErrNotFound
	}

	// svcs mapped by version
	svcs := make(map[string]*registry.Service)

	// loop through items
	for _, pod := range pods.Items {
		if pod.Status.Phase != podRunning {
			continue
		}
		// get serialised service from annotation
		svcStr, ok := pod.Metadata.Annotations[annotationServiceKeyPrefix+serviceName(name)]
		if !ok {
			continue
		}

		// unmarshal service string
		var svc registry.Service
		err := json.Unmarshal([]byte(*svcStr), &svc)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal service '%s' from pod annotation", name)
		}

		// merge up pod service & ip with versioned service.
		vs, ok := svcs[svc.Version]
		if !ok {
			svcs[svc.Version] = &svc
			continue
		}

		vs.Nodes = append(vs.Nodes, svc.Nodes...)
	}

	var list []*registry.Service
	for _, val := range svcs {
		list = append(list, val)
	}
	return list, nil
}
```

## 全部微服务列表 
通过micro.mu/type:service标签查找全部Micro微服务，并对Pod的Annotations进行验证，确保微服务信息的完整而未被注销  
```go
func (c *kregistry) ListServices() ([]*registry.Service, error) {
	pods, err := c.client.ListPods(podSelector)
	if err != nil {
		return nil, err
	}

	// svcs mapped by name
	svcs := make(map[string]bool)

	for _, pod := range pods.Items {
		if pod.Status.Phase != podRunning {
			continue
		}
		for k, v := range pod.Metadata.Annotations {
			if !strings.HasPrefix(k, annotationServiceKeyPrefix) {
				continue
			}

			// we have to unmarshal the annotation itself since the
			// key is encoded to match the regex restriction.
			var svc registry.Service
			if err := json.Unmarshal([]byte(*v), &svc); err != nil {
				continue
			}
			svcs[svc.Name] = true
		}
	}

	var list []*registry.Service
	for val := range svcs {
		list = append(list, &registry.Service{Name: val})
	}
	return list, nil
}
```