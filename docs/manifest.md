# gulfstream.yml

```yaml
name: My project
go_package_name: myproject
go_stream_name: Project
go_modules: github.com/myproject
go_version: "1.16"
go_get_packages:
- github.com/golang/mock/mockgen@v1.6.0
- github.com/go-kit/kit@v0.11.0
- github.com/gorilla/mux@v1.8.0
- github.com/go-gulfstream/tmpevents/pkg/tmpevents
go_events_pkg_name: myprojectevents
go_commands_pkg_name: myprojectcommands
go_stream_pkg_name: myprojectstream
description: Some description
mutations:
  from_commands:
  - mutation: UpdateCounter
    in_command:
      name: CounterInfo
      payload: CounterInfoPayload
    out_event:
      name: CounterUpdated
      payload: CounterUpdatedPayload
  - mutation: AddSession
    in_command:
      name: SessionInfo
      payload: SessionInfoPayload
    out_event:
      name: AddedSession
      payload:  AddedSessionPayload 
  from_events:
  - mutation: RegisterSession
    in_event:
      name: tmpevents.SessionRegistered
      payload: tmpevents.SessionRegisteredPayload
    out_event:
      name: Confirmed
      payload: ConfirmedPayload
import_events:
- github.com/go-gulfstream/tmpevents/pkg/tmpevents
storage_adapter:
  id: 0
  enable_journal: false
publisher_adapter:
  id: 0
contributors: []
created_at: 2021-08-09T21:40:48.342305Z
updated_at: 2021-08-09T21:52:45.896692Z

# available storage adapters:
# id:0, name: Memory
# id:1, name: Redis
# id:2, name: PostgreSQL

# available publisher adapters:
# id:0, name: Memory
# id:1, name: Kafka
# id:2, name: WAL Connector
```