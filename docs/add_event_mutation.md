# Add event mutation:

```yaml
 - mutation: RegisterSession
   in_event:
    name: tmpevents.SessionRegistered
    payload: tmpevents.SessionRegisteredPayload
   out_event:
    name: Confirmed
    payload: ConfirmedPayload
```

Be sure specified the import path for ```in_event``` from another package. 
For example:
```yaml
import_events:
- github.com/go-gulfstream/tmpevents/pkg/tmpevents
```

Don't forget to add the package:
```yaml
go_get_packages:
  - github.com/go-gulfstream/tmpevents/pkg/tmpevents
```
