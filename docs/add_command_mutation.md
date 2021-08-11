# Add command mutation:

```yaml
 - mutation: UpdateCounter
   in_command:
    name: CounterInfo
    payload: CounterInfoPayload
   out_event:
    name: CounterUpdated
    payload: CounterUpdatedPayload
```